package lib

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          string
	Name        string
	Details     string `json:",omitempty"`
	Project     string `json:",omitempty"`
	Status      string `json:",omitempty"`
	CreatedAt   time.Time
	DueAt       *time.Time `json:",omitempty"`
	CompletedAt *time.Time `json:",omitempty"`
	Tags        []string
}

func NewTask(name string) *Task {
	// Task ID is a UUID prefixed with "task-" to avoid collisions
	// and allow for easy sorting.
	id := "task-" + uuid.New().String()

	return &Task{
		ID:          id,
		Name:        name,
		Details:     "",
		CreatedAt:   time.Now(),
		DueAt:       nil,
		CompletedAt: nil,
		Tags:        []string{},
	}
}

// TaskFromBytes unmarshals the given byte slices into a Task.
func TaskFromBytes(b []byte) (*Task, error) {
	var t *Task
	err := json.Unmarshal(b, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Task) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

// MarshalBytes returns the the task's ID and JSON-encoded data as byte slices.
func (t *Task) MarshalBytes() (id []byte, task []byte, err error) {
	id = []byte(t.ID)
	task, err = json.Marshal(t)
	if err != nil {
		return nil, nil, err
	}
	return id, task, nil
}

func (t *Task) IsComplete() bool {
	return t.CompletedAt != nil
}
