package pass

type Pass interface{
	Ping() string
	Pong() string
}
