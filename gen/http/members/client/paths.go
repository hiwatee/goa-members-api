// Code generated by goa v3.2.5, DO NOT EDIT.
//
// HTTP request path constructors for the members service.
//
// Command:
// $ goa gen members/design

package client

import (
	"fmt"
)

// AddMembersPath returns the URL path to the members service add HTTP endpoint.
func AddMembersPath(a int, b int) string {
	return fmt.Sprintf("/add/%v/%v", a, b)
}
