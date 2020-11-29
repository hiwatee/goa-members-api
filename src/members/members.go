package membersapi

import (
	"context"
	"log"
	members "members/gen/members"
)

// members service example implementation.
// The example methods log the requests and return zero values.
type memberssrvc struct {
	logger *log.Logger
}

// NewMembers returns the members service implementation.
func NewMembers(logger *log.Logger) members.Service {
	return &memberssrvc{logger}
}

// Add implements add.
func (s *memberssrvc) Add(ctx context.Context, p *members.AddPayload) (res int, err error) {
	s.logger.Print("members.add")
	s.logger.Print(p)
	return p.A + p.B, nil
}
