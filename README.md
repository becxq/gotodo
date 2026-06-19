# gotodo - CLI Todo Application

gotodo is a command-line todo application built in Go that allows you to manage tasks with priorities, due dates, and completion status. Tasks are persisted in a JSON file (`tasks.json`). [1](#1-0) [2](#1-1) 

## Building and Running

### Prerequisites
- Go 1.26.4 or later [3](#1-2) 

### Build
```bash
go build cmd/api/main.go
```

### Run
```bash
go run cmd/api/main.go
```

After building, you can run the binary directly with the `gotodo` command. [1](#1-0) 

## Usage

The application uses Cobra for CLI command handling and supports the following subcommands: [4](#1-3) 

### Add a Task
```bash
gotodo add --name "Task name" --priority 1 --due "30m"
```
- `--name, -n`: Task name (required) [5](#1-4) 
- `--priority, -p`: Priority level 1-3 (default: 1) [6](#1-5) 
- `--due, -d`: Due duration (e.g., "30m", "1h", default: "30m") [7](#1-6) 

### Remove a Task
```bash
gotodo rm --id "task-id"
```
- `--id, -i`: Task ID (required) [8](#1-7) 

### Mark Task as Done
```bash
gotodo done --id "task-id"
```
- `--id, -i`: Task ID (required) [9](#1-8) 

### List All Tasks
```bash
gotodo list
```
Displays all tasks with their ID, name, due date, priority, and status. [10](#1-9) 

### Clear All Tasks
```bash
gotodo clear
```
Removes all tasks from storage. [11](#1-10) 

## Architecture

The application follows a layered architecture: [12](#1-11) 

- **CLI Layer** (`internal/commands`): Handles command-line interface using Cobra
- **Service Layer** (`internal/service`): Contains business logic via `TaskService` [13](#1-12) 
- **Repository Layer** (`internal/repository`): Manages data persistence to JSON file [14](#1-13) 
- **Models** (`internal/models`): Defines the `Task` struct and repository interface [15](#1-14) 

## Notes
The application stores all tasks in a single `tasks.json` file in the same directory where the application is run. The Task model includes fields for ID, name, priority (1-3), due date (time.Time), and completion status (boolean). [2](#1-1) 

Wiki pages you might want to explore:
- [Glossary (becxq/gotodo)](/wiki/becxq/gotodo#5)

