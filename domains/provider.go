package domains

//"fmt"

type Provider interface {
	// Pass in configuration, using an interface to support different
	// configuration object

	// TODO: Kinda hate this function name, and it should probably
	// return a usable connection object
	InitializeAPI()

	// TODO: Need some way to populate a morecomplex domain object that will hold more detailed
	// information about the domain that includes DNS records and registered name servers
	RegisteredDomains() []string
	RegisterDomain(string)
	//TODO: Add DNS configuration for domains
}
