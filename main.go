package main

import (
	"encoding/json"
	"fmt"

	. "github.com/hackwave/color"

	//"github.com/scrambleshell/wormholes/domains"
	"github.com/scrambleshell/wormholes/domains/providers/namecheap"
)

type Config struct {
	NamecheapConfig namecheap.Config `json:"namecheap"`
}

func main() {
	fmt.Println(Magento("Wormholes"), Gray(":"), Blue("Endpoint Management"))
	fmt.Println(Gray("=========================="))

	// TODO: Load JSON configuration
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(Red("[Error]"), Gray(err))
	}

	var config Config
	json.Unmarshal(jsonFile, &config)

	fmt.Println("Loaded JSON Configuration: ", config)
	// TODO: Example usage of the domain provider interface
	//nc := NamecheapProvider{
	//	Config: {
	//		username:    "kosmosblack",
	//		apiToken:    "",
	//		apiUsername: "",
	//	},
	//}
}
