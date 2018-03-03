package main

type NameProvider struct {
	Config
}

func (self NameProvider) SetDebug(debug bool) {
	self.Config.debug = debug
}

func (self NameProvider) InitializeAPI() {
	// TODO: Initialize a API object
}

func (self NameProvider) RegisteredDomains() []domains.Domain {
	// TODO: List all domains registered with Name.com
}

func (self NameProvider) RegisterDomain(name string) {
	// TODO: Register a domain with Name.com
}
