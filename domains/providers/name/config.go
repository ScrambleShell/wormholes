package name

type Config struct {
	debug    bool
	username string
	password string
	API
}

type API struct {
	token string
}
