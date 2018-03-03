package namecheap

import (
	"fmt"

	. "github.com/hackwave/color"
	namecheap "github.com/scrambleshell/namecheap-go"
)

type NamecheapProvider struct {
	Config
	API *namecheap.Client
}

func (self NamecheapProvider) SetDebug(debug bool) {
	self.Config.Debug = debug
}

func (self NamecheapProvider) InitializeAPI() NamecheapProvider {
	self.API = namecheap.NewClient(self.Config.Username, self.Config.API.Token, self.Config.API.Username)
	return self
}

// TODO: Domain should probably be moved to a more general models package to avoid circular imports
func (self NamecheapProvider) RegisteredDomains() []string {
	if self.Config.Debug {
		fmt.Println(Gray("Looking up domains for the API account: "), Green(self.Config.API.Username))
	}
	// TODO: Get maximum then use that to get all the domains
	var domainNames []string
	fmt.Println("API as string is: ", self.API)
	domains, err := self.API.DomainsGetList(1, 100)
	if err != nil {
		fmt.Println(Red("[Error]"), err)
	}
	for _, domain := range domains {
		fmt.Println("_______________________________")
		fmt.Println("|                             |")
		fmt.Println("|___Namecheap_Domain_Object___|")
		fmt.Println("| ID         |", domain.ID)
		fmt.Println("| Name       |", domain.Name)
		fmt.Println("| User       |", domain.User)
		fmt.Println("| Created    |", domain.Name)
		fmt.Println("| Expires    |", domain.Expires)
		fmt.Println("| IsExpired  |", domain.IsExpired)
		fmt.Println("| IsLocked   |", domain.IsLocked)
		fmt.Println("| AutoRenew  |", domain.AutoRenew)
		fmt.Println("| WhoisGuard |", domain.WhoisGuard)
		fmt.Println("|____________|________________|")

		domainNames = append(domainNames, domain.Name)
	}
	return domainNames
}

func (self NamecheapProvider) RegisterDomain(name string) {
	// TODO: Register a domain with Namecheap.com
}
