package main

import (
	//"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	. "github.com/hackwave/color"

	//"github.com/scrambleshell/wormholes/domains"
	"github.com/scrambleshell/wormholes/domains/providers/namecheap"
)

type Config struct {
	Debug     bool             `json:"debug" yaml:"debug"`
	Namecheap namecheap.Config `json:"namecheap" yaml:"namecheap"`
}

type Application struct {
	Name           string
	Version        string
	Description    string
	ConfigFilepath string
	Config         `json:"config" yaml:"config"`
}

func main() {
	// TODO: Get i18n translations and strings loading from JSON or YAML early, so its not
	// a difficult transition later.
	app := Application{
		Name:           "Wormholes",
		Version:        "0.1.0",
		Description:    "Diverse end point management",
		ConfigFilepath: "./config.yaml",
	}

	fmt.Println(Magenta(app.Name), Gray(":"), Blue(app.Description))
	fmt.Println(Gray("=========================="))
	fmt.Println(Gray("[Config]"), "loading YAML configuration file: ", app.ConfigFilepath)
	yamlFile, err := ioutil.ReadFile(app.ConfigFilepath)
	if err != nil {
		fmt.Println(Red("[Fatal Error]"), err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(yamlFile, &app.Config)
	if err != nil {
		fmt.Println(Red("[Fatal Error]"), err)
		os.Exit(1)
	}
	fmt.Println(Green("[Config]"), "Configuration successfully loaded.")
	if app.Config.Namecheap.Username != "" &&
		app.Config.Namecheap.API.Username != "" &&
		app.Config.Namecheap.API.Token != "" {
		fmt.Println(Gray("[Config]"), "Namecheap Domain Provider API configuration found, initializing an API connection to the domain provider...")
		nc := namecheap.NamecheapProvider{
			Config: namecheap.Config{
				Debug:    app.Config.Debug,
				Username: app.Config.Namecheap.Username,
				API: namecheap.API{
					Username: app.Config.Namecheap.API.Username,
					Token:    app.Config.Namecheap.API.Token,
				},
			},
		}

		nc = nc.InitializeAPI()
		fmt.Println(Green("[Domains:Namecheap:API]"), "API connection successfully Established.")

		fmt.Println(Gray("[Domains:Namecheap:API"), "Query API for all registered domains with domain provider...")
		domainNames := nc.RegisteredDomains()
		fmt.Println("Domain name count returned from domains names is: ", len(domainNames))

	} else {
		fmt.Println(Red("[Fatal Error]"), Gray("No domain provider configuration found, modify the loaded configuration so it contains the API information needed to connect to the domain provider API."))
	}
}
