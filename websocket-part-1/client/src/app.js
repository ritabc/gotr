import { Todo } from "./todo";

export class App {
    constructor() {
        this.heading = 'Todos';
        this.todos = [];
        this.todoDescription = '';
        // Create WebSockete connection
        this.socket = new WebSocket('ws://localhost:8081/ws');
        // Connection opened
        this.socket.addEventListener('open', function (event) {
            this.socket.send('Hello Server!');
        });
        // Listen for messages
        this.socket.addEventListener('message', function (event) {
            console.log('Message from server ', event.data);
        });
    }

    addTodo() {
        if (this.todoDescription) {
            var todo = new Todo(this.todoDescription);
            this.todos.push(todo);
            this.todoDescription = '';
        }
    }

    removeTodo(todo) {
        var index = this.todos.indexOf(todo)
        if (index != -1) {
            this.todos.splice(index, 1)
        }
    }
}
