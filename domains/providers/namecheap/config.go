package namecheap

type Config struct {
	Debug    bool   `json:"debug" yaml:"debug"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	API      `json:"api" yaml:"api"`
}

type API struct {
	Username  string `json:"username" yaml:"username"`
	Token     string `json:"token" yaml:"token"`
	IPAddress string `json:"ip_address" yaml:"ip_address"`
}
