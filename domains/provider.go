package domainprovider

import (
//"fmt"
)

type DomainProvider interface {
	// Pass in configuration, using an interface to support different
	// configuration object

	// TODO: Kinda hate this function name, and it should probably
	// return a usable connection object
	InitializeAPI()

	RegisteredDomains() []Domain
	RegisterDomain(string)
	//TODO: Add DNS configuration for domains
}
