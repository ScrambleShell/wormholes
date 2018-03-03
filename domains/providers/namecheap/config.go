package namecheap

type Config struct {
	debug    bool   `json:"debug"`
	username string `json:"username"`
	password string `json:"password"`
	API      `json:"api"`
}

type API struct {
	username string `json:"username"`
	token    string `json:"token"`
}
