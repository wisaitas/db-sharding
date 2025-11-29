package create

type Request struct {
	Name string `json:"name" validate:"require"`
}
