package models

type Auditor struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (a Auditor) GetHandshake() string {
	return "Se consultó información del usuario " + a.Name
}
