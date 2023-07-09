package signalmultipleclose

type Signal interface {
	Close()                // Closes internal chan
	Done() <-chan struct{} // Similar to context/Context.Done
	CloseCount() int       // Only for testing
}
