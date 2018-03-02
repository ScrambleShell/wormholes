package domainprovider

import (
//"fmt"
)

type DomainProvider interface {
	RegisteredDomains()
	RegisterDomain(string)
	//TODO: Add DNS configuration for domains
}
