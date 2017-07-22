package api

type ActorActress struct {
	Name string `json:"name"`
}

type TVSerie struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Casting     []ActorActress `json:"casting"`
	Stars       int            `json:"stars"`
}
