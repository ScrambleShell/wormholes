package namecheap

type Config struct {
	debug    bool   `json:"debug"`
	username string `json:"username"`
	password string `json:"password"`
	api      `json:"api"`
}

type api struct {
	username string `json:"username"`
	token    string `json:"token"`
}
