package netrpcdemo2

const (
	HelloServiceName        = "HelloService"
	Provider         string = "provider"
	Consumer         string = "consumer"
)

type User struct {
	Id   int
	Name string
	Role string
}
