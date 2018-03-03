package name

type Config struct {
	debug    bool   `json:"debug" yaml:"debug"`
	username string `json:"username" yaml:"username"`
	password string `json:"password" yaml:"password"`
	API      `json:"api" yaml:"api"`
}

type API struct {
	token string `json:"token" yaml:"yaml"`
}
