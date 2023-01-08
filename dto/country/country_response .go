package countrydto

type CountryResponse struct {
	// ID   int    `json:`
	Name string `json:"name"`
}
type CountryDeleteResponse struct {
	ID int `json:"id"`
	// Name string `json:"name"`
}
