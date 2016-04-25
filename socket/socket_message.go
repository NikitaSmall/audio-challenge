/*Package socket holds websocket preferences, initialisation and hub initialisation.
 * This file holds constants that can help to communicate with webclients standart way.
 */
package socket

// constants contains special instuctions
// for the clients in browser
const (
	TaskAdd      = "taskadded"
	TaskComplete = "taskcompleted"
	TaskFail     = "taskfailed"
)

// Message is basic message that will be sent to the browser
type Message struct {
	Action  string      `json:"action"`
	Message interface{} `json:"message"`
}
