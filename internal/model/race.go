package model

type Race struct {
	Name            string           `bson:"name" json:"name"`
	EngName         string           `bson:"eng_name" json:"eng_name"`
	Source          string           `bson:"source" json:"source"`
	Type            string           `bson:"type" json:"type"`
	Attributes      Attributes       `bson:"attributes" json:"attributes"`
	Size            string           `bson:"size" json:"size"`
	Speed           string           `bson:"speed" json:"speed"`
	DarkVision      string           `bson:"dark_vision" json:"dark_vision"`
	Features        []Features       `bson:"features" json:"features"`
	DragonsLegacies *[]DragonsLegacy `bson:"dragons_legacies,omitempty" json:"dragons_legacies,omitempty"`
}

type Attributes struct {
	Strength     int    `bson:"strength" json:"strength"`
	Constitution int    `bson:"constitution" json:"constitution"`
	Dexterity    int    `bson:"dexterity" json:"dexterity"`
	Intelligence int    `bson:"intelligence" json:"intelligence"`
	Wisdom       int    `bson:"wisdom" json:"wisdom"`
	Charisma     int    `bson:"charisma" json:"charisma"`
	Any          *[]Any `bson:"any,omitempty" json:"any,omitempty"`
}
type Any struct {
	Number   int `bson:"number" json:"number"`
	Modifier int `bson:"modifier" json:"modifier"`
}

type Features struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}
type DragonsLegacy struct {
	Color       string `bson:"color"`
	DamageType  string `bson:"damage_type"`
	BreathShape string `bson:"breath_shape"`
	SavingThrow string `bson:"saving_throw"`
}
