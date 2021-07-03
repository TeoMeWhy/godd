package rpg

interface Mannager {

	Listen(userName, action string, ch chan<- string) int
	

}