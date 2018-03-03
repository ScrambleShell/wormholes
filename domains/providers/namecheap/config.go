package namecheap

type Config struct {
	Debug    bool   `json:"debug"`
	Username string `json:"username"`
	Password string `json:"password"`
	API      `json:"api"`
}

type API struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
