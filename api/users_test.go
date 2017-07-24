package api

import (
	. "gopkg.in/check.v1"
	"os"
)

func (s *TestSuite) TestShouldHashNumberCorrectly(c *C) {
	expected := "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"

	hashed := createHash("123456")
	c.Assert(expected, Equals, hashed)
}

func (s *TestSuite) TestShouldHashNumberAndLettersCorrectly(c *C) {
	expected := "0a05d7b27cc7a2b1ca704adcbd1d6e3ab2c19ece000586f03bceeabf24547e43"

	hashed := createHash("1a2b3c4d")
	c.Assert(expected, Equals, hashed)
}

func (s *TestSuite) TestShouldCreateTokenCorrectly(c *C) {
	os.Setenv("SECRET", "secret")
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIn0.KCxI3hD2u8QlZ76J_T_OqOUwWmIetfKJfmwRDBiBFVQ"

	token, _ := createToken("admin@example.com")
	c.Assert(expected, Equals, token)
}
