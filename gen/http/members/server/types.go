// Code generated by goa v3.2.5, DO NOT EDIT.
//
// members HTTP server types
//
// Command:
// $ goa gen members/design

package server

import (
	members "members/gen/members"
)

// NewAddPayload builds a members service add endpoint payload.
func NewAddPayload(a int, b int) *members.AddPayload {
	v := &members.AddPayload{}
	v.A = a
	v.B = b

	return v
}
