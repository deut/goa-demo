// Code generated by goa v3.20.1, DO NOT EDIT.
//
// HTTP request path constructors for the hamster service.
//
// Command:
// $ goa gen goa-demo/design

package client

import (
	"fmt"
)

// ListHamsterPath returns the URL path to the hamster service list HTTP endpoint.
func ListHamsterPath() string {
	return "/hamsters"
}

// CreateHamsterPath returns the URL path to the hamster service create HTTP endpoint.
func CreateHamsterPath() string {
	return "/haster"
}

// ShowHamsterPath returns the URL path to the hamster service show HTTP endpoint.
func ShowHamsterPath(hamsterID string) string {
	return fmt.Sprintf("/hamster/%v", hamsterID)
}
