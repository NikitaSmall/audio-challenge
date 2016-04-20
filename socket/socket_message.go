package socket

// constants contains special instuctions
// for the clients in browser
const (
	TaskAdd      = "taskAdded"
	TaskComplete = "taskCompleted"
	TaskFail     = "taskFailed"
)

// basic message that will be sent to the browser
type SocketMessage struct {
	Action  string      `json:"action"`
	Message interface{} `json:"message"`
}
