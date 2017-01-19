package show



type Show struct {
	name string
	age int
}
//msg传进来string类型数据 放进pings通道
func (s *Show) Ping(pings chan<- string,msg string){
	pings <- msg
}
//pings 接收的数据放进msg，msg将数据放进pongs通道
func (s *Show) Pong(pings <-chan string,pongs chan <- string){
	msg:= <- pings
	pongs <- msg
}
func (s *Show) Set(n string,a int) {
	s.name = n
	s.age = a
}