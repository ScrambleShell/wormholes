package domainprovider

import (
//"fmt"
)

type DomainProvider interface {
	// Pass in configuration, using an interface to support different
	// configuration object
	InitializeAPI(interface{}) // TODO: Kinda hate this function name

	RegisteredDomains()
	RegisterDomain(string)
	//TODO: Add DNS configuration for domains
}
