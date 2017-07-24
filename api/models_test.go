package api

import (
	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestTVSerieWithoutTitleShouldReturnError(c *C) {
	actorActress := ActorActress{
		Name: "Actor Test",
	}

	var casting []ActorActress
	casting = append(casting, actorActress)

	serie := TVSerie{
		Description: "Description Test",
		Casting:     casting,
		Stars:       5,
	}

	err := serie.Validate()
	c.Assert(err.Error(), Equals, VALIDATION_ERROR)
}

func (s *TestSuite) TestTVSerieWithoutDescriptionShouldReturnError(c *C) {
	actorActress := ActorActress{
		Name: "Actor Test",
	}

	var casting []ActorActress
	casting = append(casting, actorActress)

	serie := TVSerie{
		Title:   "Title Test",
		Casting: casting,
		Stars:   5,
	}

	err := serie.Validate()
	c.Assert(err.Error(), Equals, VALIDATION_ERROR)
}

func (s *TestSuite) TestTVSerieWithoutCastingShouldReturnError(c *C) {
	var casting []ActorActress

	serie := TVSerie{
		Title:       "Title Test",
		Description: "Description Test",
		Casting:     casting,
		Stars:       5,
	}

	err := serie.Validate()
	c.Assert(err.Error(), Equals, VALIDATION_ERROR)
}

func (s *TestSuite) TestTVSerieWithoutStarsShouldReturnError(c *C) {
	actorActress := ActorActress{
		Name: "Actor Test",
	}

	var casting []ActorActress
	casting = append(casting, actorActress)

	serie := TVSerie{
		Title:       "Title Test",
		Description: "Description Test",
		Casting:     casting,
	}

	err := serie.Validate()
	c.Assert(err.Error(), Equals, VALIDATION_ERROR)
}

func (s *TestSuite) TestTVSerieShouldntReturnError(c *C) {
	actorActress := ActorActress{
		Name: "Actor Test",
	}

	var casting []ActorActress
	casting = append(casting, actorActress)

	serie := TVSerie{
		Title:       "Title Test",
		Description: "Description Test",
		Casting:     casting,
		Stars:       5,
	}

	err := serie.Validate()
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestUserWithoutEmailShouldReturnError(c *C) {
	user := User{
		Password: "123456",
	}

	err := user.Validate()
	c.Assert(err.Error(), Equals, VALIDATION_ERROR)
}

func (s *TestSuite) TestUserWithoutPasswordShouldReturnError(c *C) {
	user := User{
		Email: "example@example.com",
	}

	err := user.Validate()
	c.Assert(err.Error(), Equals, VALIDATION_ERROR)
}

func (s *TestSuite) TestUserValidationShouldntReturnError(c *C) {
	user := User{
		Email:    "example@example.com",
		Password: "123456",
	}

	err := user.Validate()
	c.Assert(err, IsNil)
}
