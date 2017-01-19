package show



type Show struct {
	name string
	age int
}
func (s *Show) Ping(pings chan<- string,msg string){
	pings <- msg
}
func (s *Show) Pong(pings <-chan string,pongs chan <- string){
	msg:= <- pings
	pongs <- msg
}
func (s *Show) Set(n string,a int) {
	s.name = n
	s.age = a
}