package main

import (
	"github.com/gorilla/websocket"
)

type (
	Todo struct {
		ID          int    `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
		Done        bool   `json:"done"`
	}
	Todos         []Todo
	ClientRequest struct { // What is the client sending
		Username string `json:"username,omitempty"`
		Type     string `json:"type,omitempty"` // hello, or add, or remove
		Todo     Todo   `json:"todo,omitempty"` // useful for adding todos
		ID       int    `json:"id,omitempty"`   // usefule for deleting todo
	}
	ClientResponse struct { // what is the client sending
		Todos `json:"todos,omitempty"`
	}
	Connections []*websocket.Conn
	Client      struct {
		Todos
		Connections
	}
)
