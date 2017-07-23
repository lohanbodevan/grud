package api

import (
	. "gopkg.in/check.v1"
	"reflect"
)

func (s *TestSuite) TestGenerateUUIDShouldReturnString(c *C) {
	uuid := generateUUID()
	c.Assert(reflect.TypeOf(uuid).Kind(), Equals, reflect.String)
}
