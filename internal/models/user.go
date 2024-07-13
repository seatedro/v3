package models

type User struct {
	Name     string      `json:"name"`
	Location string      `json:"location"`
	Role     string      `json:"role"`
	Bio      string      `json:"bio"`
	Github   string      `json:"github"`
	Twitter  string      `json:"twitter"`
	Linkedin string      `json:"linkedin"`
	Work     []WorkEx    `json:"work"`
	Projects []ProjectEx `json:"projects"`
}

type WorkEx struct {
	Company     string `json:"company"`
	Href        string `json:"href"`
	Position    string `json:"position"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}

type ProjectEx struct {
	Title       string `json:"title"`
	Href        string `json:"href"`
	Description string `json:"description"`
}
