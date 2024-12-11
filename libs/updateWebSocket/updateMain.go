package updatewebsocket

type StatusUpdateLoopCount struct {
	godUpdateCount         int
	managerUpdateCount     int
	disUpDateCount         int
	schoolOwnerUpdateCount int
	schoolOUpdateCount     int
	teacherUpdatecount     int
	studentUpdateCount     int
	busUpdateCount         int
}

func MainUpdateLoop() {
	// update loop control the responce reption of the websocket reponce
	// each app have thair own update loop based on api number or key
}



func ResetUpdateLoop() {}
