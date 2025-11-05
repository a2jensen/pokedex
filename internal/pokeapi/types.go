package pokeapi

type ListResp struct {
    Count    int                `json:"count"`
    Next     *string            `json:"next"`     // can be null
    Previous *string            `json:"previous"` // can be null
    Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}