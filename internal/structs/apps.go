package structs

type Apps struct {
	Applications []Application `json:"applications"`
}

type Application struct {
	ID int `json:"id"`
	Name string `json:"name"`
	TemplateName string `json:"template_name"`
	Realm int `json:"realm"`
}