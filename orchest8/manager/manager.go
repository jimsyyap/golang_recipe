package manager

import (
    "cube/task"
    "fmt"
    "github.com/golang-collections/queue"
    "github.com/google/uuid"
)

type Manager struct {
    Pending         queue.Queue
    TaskDb          map[string]task.Task
    EventDb         map[string]task.EventDb
    Workers         []string
    WorkerTaskMap   map[string][]uuid.UUID
    TaskWorkerMap   map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
    fmt.Println("I will select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
    fmt.Println("I will update the tasks")
}

func (m *Manager) SendWork() {
    fmt.Println("I will send work to the worker")
}
