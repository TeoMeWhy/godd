package rpg

type Manager interface {
	Listen(userName, action string, ch chan<- string) int
}
