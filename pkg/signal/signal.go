package signal

var signals chan bool

func init() {
	signals = make(chan bool)
}

// Wait waits for signal using channel
func Wait() {
	<-signals
}

// Signal releases waited goroutine using channel
func Signal() {
	select {
	case signals <- true:
	default:
	}
}
