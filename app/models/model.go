package models

var (
	quotaArr   []Quota
	quotaIdSeq = 0
)

/*
!!!!!!!!!
ATTENTION: This part of code should be in db package connection.go file
but it is here because we use just arrays instead of real Database.
If we put this code there we get "import cycle not allowed" ERROR

In the real application with normal DB this will not happen!!! and structure will be beautiful
!!!!!!!!!
*/

// Connect - creates Quota map, so-called "database"
func ConnectDB() {
	quotaArr = make([]Quota, 0)
}

func DisconnectDB() {
	quotaArr = nil
}
