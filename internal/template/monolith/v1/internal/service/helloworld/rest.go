package helloworld

import (
	"context"
)

// @ee.http.Controller
// @ee.http.Route("/api/v1/message")
// @ee.stereotype.Controller("msg")
type RestController struct {
	messages MessageRepository
}

func NewRestController(ctr MessageRepository) *RestController {
	return &RestController{ctr}
}

// All returns the messages.
//
// @ee.http.Route("/")
// @ee.http.Method("GET")
// @ee.http.QueryParam("limit")
func (s *RestController) All(ctx context.Context, limit int) ([]Message, error) {
	return s.messages.FindAll(ctx,limit)
}


// One returns a message by id.
//
// @ee.http.Route("/")
// @ee.http.Method("GET")
// @ee.http.QueryParam("limit")
func (s *RestController) One(ctx context.Context, limit int) ([]Message, error) {
	return s.messages.FindAll(ctx,limit)
}
