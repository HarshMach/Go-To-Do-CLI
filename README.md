# Go To-Do CLI App

A simple command-line To-Do list manager written in Go 
Add tasks, mark them as done, and list everything — all from the terminal. 

---

Features

- Add tasks from terminal
- Mark tasks as done
- View all tasks
- Tasks saved to JSON file
- Built with flags (`-add`, `-list`, `-done`)
- Unit tested 

---

## 🛠️ Usage

### Add a task

```bash
go run main.go todo.go  -add="Test application"
```
### List the tasks 

```bash
go run main.go todo.go -list
```

### Mark a task as done
```bash
go run main.go todo.go  -done=1
````

### Run Tests
```bash
go test -v
````
---
How to Start
```bash
git clone https://github.com/HarshMach/go-todo-cli.git
cd go-todo-cli
```
---
This project was built to practice:

- Go basics
- File I/O
- Structs & JSON
- Command-line flags
- Writing tests
