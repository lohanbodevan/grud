package api

import (
	. "gopkg.in/check.v1"
	"os"
)

func (s *TestSuite) TestGivenCorrectTokenShouldAuthorizeRequest(c *C) {
	os.Setenv("SECRET", "secret")
	authorization := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIn0.KCxI3hD2u8QlZ76J_T_OqOUwWmIetfKJfmwRDBiBFVQ"

	response := IsAuthorized(authorization)
	c.Assert(response, Equals, true)
}

func (s *TestSuite) TestGivenTokenWithWrongSecretShouldUnauthorizeRequest(c *C) {
	os.Setenv("SECRET", "OTHER")
	authorization := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluLWZhbHNlQGV4YW1wbGUuY29tIn0.eoIVDOXKodTWgDl_lZGjVo0JvRtuMvCJjgAsDKyYqF0"

	response := IsAuthorized(authorization)
	c.Assert(response, Equals, false)
}

func (s *TestSuite) TestGivenEmptyTokenShouldUnauthorizeRequest(c *C) {
	os.Setenv("SECRET", "secret")
	authorization := ""

	response := IsAuthorized(authorization)
	c.Assert(response, Equals, false)
}
