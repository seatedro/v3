package models

type User struct {
	Name      string `json:"name"`
	Workplace string `json:"workplace"`
	Role      string `json:"role"`
	Bio       string `json:"bio"`
	Github    string `json:"github"`
	Twitter   string `json:"twitter"`
	Linkedin  string `json:"linkedin"`
}
