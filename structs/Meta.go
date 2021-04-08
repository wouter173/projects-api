package structs

//Meta struct
type Meta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumb       string `json:"thumb"`
	Tags        []Tag  `json:"tags"`
	Links       []Link `json:"links"`
}

//Tag struct
type Tag struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Url   string `json:"url"`
}

// Link struct
type Link struct {
	Url string `json:"url"`
}
