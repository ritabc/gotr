import { Todo } from "./todo";

export class App {
    constructor() {
        this.heading = 'Todos';
        this.todos = [];
        this.todoDescription = '';
        this.username = '';
        this.loggedIn = false;
    }

    doLogin() {
        var app = this
        // upon logging in:
        if (!this.username) {
            return
        }
        this.loggedIn = true
        // Create WebSocket connection
        var socket = new WebSocket('ws://localhost:8081/ws');
        this.socket = socket;

        // on websocket error
        this.socket.addEventListener('error', function (event) {
            console.log(event)
        });

        // Connection opened
        this.socket.addEventListener('open', function (event) {
            var msg = { "type": "hello", "username": app.username }
            app.socket.send(JSON.stringify(msg));
            console.log(msg)
        });

        // Listen for messages
        this.socket.addEventListener('message', function (event) {
            var msg = JSON.parse(event.data)
            app.todos = msg.todos;
        });
    }

    addTodo() {
        if (this.todoDescription) {
            var todo = new Todo(this.todoDescription);
            var msg = { "type": "add", "todo": todo, "username": this.username }
            this.socket.send(JSON.stringify(msg));
            this.todoDescription = '';
        }
    }

    removeTodo(id) {
        var msg = { "type": "delete", "id": id, "username": this.username }
        this.socket.send(JSON.stringify(msg));
    }

    toggleTodo(id) {
        console.log("updateTodo():", id)
        var msg = { "type": "toggle.done", "id": id, "username": this.username }
        this.socket.send(JSON.stringify(msg));
    }
}
