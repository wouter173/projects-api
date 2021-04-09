package structs

//Meta struct
type Meta struct {
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumb       string `json:"thumb"`
	Tags        []Tag  `json:"tags"`
	Links       Links  `json:"links"`
}

//Tag struct
type Tag struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Url   string `json:"url"`
}

// Link struct
type Links struct {
	Github string `json:"github"`
	Live   string `json:"live"`
}
