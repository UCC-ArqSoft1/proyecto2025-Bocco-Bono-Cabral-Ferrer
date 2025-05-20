package domain

type Activity struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"`
	Category    string `json:"category"`
	Profesor    string `json:"profesor"`
	Day         string `json:"day"`
	Hour        string `json:"hour"`
}

type Activities []Activity
