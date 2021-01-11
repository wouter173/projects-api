package main

//Project struct
type Project struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Links       links  `json:"links,omitempty"`
	Info        info   `json:"info,omitempty"`
}

type links struct {
	Home   string `json:"home,omitempty"`
	Github string `json:"github,omitempty"`
}

type info struct {
	Problem  string `json:"problem,omitempty"`
	Solution string `json:"solution,omitempty"`
}
