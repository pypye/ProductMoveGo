package infrastructure

type Injector struct {
	db     *Database
	server *Server
	any    map[string]interface{}
}

var instance *Injector

func NewInjector() *Injector {
	if instance == nil {
		instance = &Injector{}
	}
	return instance
}

func InitDB(db *Database) {
	instance.db = db
}

func InitServer(server *Server) {
	instance.server = server
}

func Init(key string, value interface{}) {
	instance.any[key] = value
}

func GetServer() *Server {
	return instance.server
}

func GetDB() *Database {
	return instance.db
}

func Get(key string) interface{} {
	return instance.any[key]
}
