package namecheap

import (
	"fmt"

	. "github.com/hackwave/color"
	namecheap "github.com/scrambleshell/namecheap-go"
)

type NamecheapProvider struct {
	config
	API *namecheap.Client
}

func (self NamecheapProvider) SetDebug(debug bool) {
	self.config.debug = debug
}

func (self NamecheapProvider) InitializeAPI() {
	self.API = namecheap.NewClient(self.config.username, self.config.api.token, self.config.api.username)
}

func (self NamecheapProvider) RegisteredDomains() {
	if self.config.debug {
		fmt.Println(Gray("Looking up domains for the API account: "), Green(self.config.api.username))
	}
	// TODO: Get maximum then use that to get all the domains
	domains, err := self.API.DomainsGetList(1, 100)
	if err != nil {
		fmt.Println(Red("[Error]"), err)
	}
}

func (self NamecheapProvider) RegisterDomain(name string) {
	// TODO: Register a domain with Namecheap.com
}
