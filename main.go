package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	. "github.com/hackwave/color"

	//"github.com/scrambleshell/wormholes/domains"
	"github.com/scrambleshell/wormholes/domains/providers/namecheap"
)

type Config struct {
	Debug           bool             `json:"debug"`
	NamecheapConfig namecheap.Config `json:"namecheap"`
}

func main() {
	fmt.Println(Magenta("Wormholes"), Gray(":"), Blue("Endpoint Management"))
	fmt.Println(Gray("=========================="))

	// TODO: Load JSON configuration
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(Red("[Error]"), Gray(err))
	}

	var config Config
	err := json.Unmarshal(jsonFile, &config)
	if err != nil {
		fmt.Printf(Red("[ Error ]"), Gray(err))
	}

	fmt.Println("Loaded JSON Configuration: ", config)
	fmt.Println("Where is the false?")
	fmt.Println("Config.NamecheapConfig: ", config.NamecheapConfig)
	fmt.Println("Config.NamecheapConfig.Username: ", config.NamecheapConfig.Username)
	fmt.Println("Config.NamecheapConfig.Password: ", config.NamecheapConfig.Password)

	// TODO: Example usage of the domain provider interface
	//nc := NamecheapProvider{
	//	Config: {
	//		username:    "kosmosblack",
	//		apiToken:    "",
	//		apiUsername: "",
	//	},
	//}
}
