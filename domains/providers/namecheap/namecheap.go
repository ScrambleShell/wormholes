package namecheap

import (
	"fmt"

	. "github.com/hackwave/color"
	namecheap "github.com/scrambleshell/namecheap-go"
	"github.com/scrambleshell/wormholes/domains"
)

type NamecheapProvider struct {
	Config
	API               namecheap.Client
	RegisteredDomains []domains.Domain
}

func (self NamecheapProvider) SetDebug(debug bool) {
	self.Config.debug = debug
}

func (self NamecheapProvider) InitializeAPI() {
	self.API = namecheap.NewClient(self.Config.username, self.Config.apiToken, self.Config.apiUsername)
}

func (self NamecheapProvider) RegisteredDomains() []domains.Domain {
	if self.Config.debug {
		fmt.Println(Gray("Looking up domains for the API account: "), Green(account.username))
	}
	// TODO: Get maximum then use that to get all the domains
	domains, err := client.DomainsGetList(1, 100)
	if err != nil {
		fmt.Println(Red("[Error]"), err)
	}
}

func (self NamecheapProvider) RegisterDomain(name string) {

}
