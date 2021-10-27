package util

type BenchmarkMessage struct {
	Message string
	Closed  bool
	Title   bool
}

func NewMessage(message string, closed bool, title bool) *BenchmarkMessage {
	msg := &BenchmarkMessage{Message: message, Closed: closed, Title: title}
	return msg
}
