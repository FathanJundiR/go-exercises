package main

import (
	"fmt"
)



func main() {
	packages := 100
	delivered := 20
	customers := 4
	fmt.Printf("I have %v packages to be delivered to %v customers. I've been delivered %v packages so far\n", packages, customers, delivered)
	packages -= delivered
	fmt.Printf("So, there are %v packages left or %v packages per customers\n",  packages, packages / customers)
}
