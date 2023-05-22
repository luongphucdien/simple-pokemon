package entity

type Pokemon struct {
	Entity
	Id               int       `json:"id"`
	Type             string    `json:"type"`
	Img              string    `json:"img_link"`
	Name             string    `json:"name"`
	BaseExperience   int       `json:"base_experience"`
	EffortValueYield []float32 `json:"effort_value_yield"`
	Form             []string  `json:"form"`
	Attack           float32   `json:"attack"`
	Defense          float32   `json:"defense"`
	SpecialAttack    float32   `json:"special_attack"`
	SpecialDefense   float32   `json:"special_defense"`
	Speed            int       `json:"speed"`
	MaxHP            float32   `json:"max_hp"`
}
