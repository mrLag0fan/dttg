package model

type Spell struct {
	Name          string
	EngName       string
	Level         int
	School        string
	Source        string
	CastTime      string
	Distance      string
	Duration      string
	Comments      *Components
	IsRitual      bool
	Concentration bool
	DamageType    string
	Description   string
	Classes       []string
}

type Components struct {
	Verbal   bool
	Somatic  bool
	Material *MaterialComponents
}

type MaterialComponents struct {
	IsConsumable bool
	Price        string
	Description  string
}
