package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader

// Map users to Todos
var db map[string]*Client

// Map user to user's connections
var todoID int

//The application that listen to our websocket connection
func main() {
	db = make(map[string]*Client)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	http.HandleFunc("/ws", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	defer conn.Close()
	for {
		clientReq := &ClientRequest{}
		err := conn.ReadJSON(clientReq)
		if err != nil {
			log.Error(err)
			return
		}
		log.Infof("Message from client: %#v", clientReq)
		if len(clientReq.Username) == 0 {
			return
		}
		clientResp := &ClientResponse{}
		var todos Todos
		var username = clientReq.Username
		switch clientReq.Type {
		case "hello":
			doLogin(username, conn)
			todos = getTodos(username)
		case "add":
			todos = addTodo(username, clientReq.Todo)
		case "delete":
			todos = removeTodo(username, clientReq.ID)
		case "toggle.done":
			todos = toggleDone(username, clientReq.ID)
		}
		clientResp.Todos = todos
		connections := getConnections(username)
		log.Infof("Updatting %v clients for : %v", len(connections), username)
		for _, c := range connections {
			if err := c.WriteJSON(clientResp); err != nil {
				doLogout(username, c)
			}
		}
	}
}
func getConnections(username string) Connections {
	return db[username].Connections
}

func removeTodo(username string, id int) Todos {
	todos := db[username].Todos
	var tmp Todos
	for _, v := range todos {
		if id != v.ID {
			tmp = append(tmp, v)
		}
	}
	db[username].Todos = tmp
	return tmp
}

func toggleDone(username string, id int) Todos {
	for i, v := range db[username].Todos {
		if id == v.ID {
			db[username].Todos[i].Done = !v.Done
		}
	}
	return db[username].Todos
}

func addTodo(username string, todo Todo) Todos {
	todoID++
	todo.ID = todoID
	db[username].Todos = append(db[username].Todos, todo)
	return db[username].Todos
}

func getTodos(username string) Todos {
	return db[username].Todos
}

func doLogin(username string, c *websocket.Conn) {
	if db[username] == nil {
		db[username] = &Client{}
	}
	db[username].Connections = append(db[username].Connections, c)
}

func doLogout(username string, c *websocket.Conn) {
	conns := db[username].Connections
	var tmp Connections
	for _, v := range conns {
		if v != c {
			tmp = append(tmp, v)
		}
	}
	db[username].Connections = tmp
}
