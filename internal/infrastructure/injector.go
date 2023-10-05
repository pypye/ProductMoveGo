package infrastructure

type Injector struct {
	Db     *Database
	Server *Server
}

var instance *Injector

func NewInjector() *Injector {
	if instance == nil {
		instance = &Injector{}
	}
	return instance
}

func InitDB(db *Database) {
	instance.Db = db
}

func InitServer(server *Server) {
	instance.Server = server
}

func GetServer() *Server {
	return instance.Server
}

func GetDB() *Database {
	return instance.Db
}
