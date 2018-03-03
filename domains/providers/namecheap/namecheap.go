package main

import (
	"fmt"

	namecheap "github.com/galaxyblack/namecheap-go"
	. "github.com/hackwave/color"
)

type

type Domain struct {
	name string
}

type Account struct {
	username string
	apiToken string
	domains  []Domain
}

func main() {
	// TODO: Use a configuration file
	client := namecheap.NewClient(a.username, a.apiToken, a.username)

	fmt.Println(Magenta("Namecheap Domains"))
	fmt.Println(Gray("================="))

	fmt.Println(Gray("Looking up domains for the API account: "), Green(a.username))
	// Get a list of your domains
	domains, err := client.DomainsGetList(1, 100)
	if err != nil {
		fmt.Println(Red("[Error]"), err)
	}

	fmt.Println(Gray("Number of found domains: "), Green(len(domains)))
	for _, domain := range domains {
		fmt.Printf("Domain: %+v\n\n", domain.Name)
	}
}
