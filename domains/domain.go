package domains

import (
	"fmt"
	"time"
)

type Whois struct {
	Firstname string
	Lastname  string
	// TODO: Create a model for storing all whois data
}

type DNS struct {
	Records []string
}

type Domain struct {
	name         string
	nameservers  []string
	registeredOn time.Time
	expiresAt    time.Time
	locked       bool
	autoRenew    bool
	WhoisGuard   bool
	Whois
	DNS
}

// Example Methods
func (self Domain) PrintDomain() {
	fmt.Println(self.name)
}
