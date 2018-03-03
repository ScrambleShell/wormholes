package main

import (
	"fmt"
)

type DNS struct {
	Records []string
}

type Domain struct {
	name string
	DNS
}

// Example Methods
func (self Domain) PrintDomain() {
	fmt.Println(self.name)
}
