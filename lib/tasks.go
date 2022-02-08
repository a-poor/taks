package lib

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// TaskPriority is an enum for the priority of a task.
type TaskPriority uint32

const (
	TaskPriorityNotSet   TaskPriority = iota // An unset priority.
	TaskPriorityVeryHigh                     // Very high priority.
	TaskPriorityHigh                         // High priority.
	TaskPriorityMedium                       // Medium priority.
	TaskPriorityLow                          // Low priority.
	TaskPriorityVeryLow                      // Very low priority.
)

// Task stores data connected to a single task in the user's task list.
type Task struct {
	ID          string       `json:"id"`                    // Unique ID for the task, used as the DB key
	Name        string       `json:"name"`                  // Name of the task & main text
	Details     string       `json:"details,omitempty"`     // Optional task details
	Project     string       `json:"project,omitempty"`     // Optional project name associated with the task
	Status      string       `json:"status,omitempty"`      // Optional task status
	Priority    TaskPriority `json:"priority,omitempty"`    // Optional task priority (Priority scale: 1-5; 1 is the highest; 0 means not set)
	CreatedAt   time.Time    `json:"createdAt"`             // Timestamp of the task creation
	DueAt       *time.Time   `json:"dueAt,omitempty"`       // Optional due date for the task
	CompletedAt *time.Time   `json:"completedAt,omitempty"` // Optional completion timestamp
	Tags        []string     `json:"tags"`                  // Optional list of tags associated with the task
}

// NewTask creates a new Task with the supplied name, an auto-generated ID and
// creation timestamp, and nil values for the remaining fields.
func NewTask(name string) *Task {
	// Task ID is a UUID prefixed with "task-" to avoid collisions
	// and allow for easy sorting.
	id := "task-" + uuid.New().String()

	return &Task{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		Tags:      []string{},
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
