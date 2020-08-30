### Pre-step

Add `.env` file with format below
```
USER_DB={your username db}
PASSWORD_DB={your password db}
```

### How to Run
1. Install the package with `go get`
2. Compile the program with command `go build main.go db.go`
3. Run the file executable with subcommand:
  - `task-todo` to filter task that not completed
  - `task-done` to filter task that completed
  - `change-task` to change status task from not completed to completed
