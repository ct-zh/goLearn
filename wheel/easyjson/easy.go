package easyjson

const (
	None = iota + 1
	Man
	Woman
)

//easyjson:json
type User struct {
	Name   string
	Age    int
	Gender byte
	secret string
}
