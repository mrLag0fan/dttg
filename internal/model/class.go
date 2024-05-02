package model

type Class struct {
	id            string      `bson:"_id"`
	Name          string      `bson:"name"`
	EngName       string      `bson:"eng_name"`
	Source        string      `bson:"source"`
	HitDice       string      `bson:"hit_dice"`
	FirstLevelHp  int         `bson:"1th_level_hp"`
	LevelUpAvgeHp int         `bson:"level_up_avge_hp"`
	Proficiency   proficiency `bson:"proficiency"`
	GoldEquipEcvl string      `bson:"gold_equip_ecvl"`
	Features      []feature   `bson:"features"`

	KnownCantrips     *[]int        `bson:"known_cantrips"`
	KnownSpells       *[]int        `bson:"known_spells"`
	KnownInvocation   *[]int        `bson:"known_invocation"`
	Invocations       *[]invocation `bson:"invocations"`
	Rage              *[]string     `bson:"rage"`
	MartialArts       *[]string     `bson:"martial_arts"`
	Ci                *[]int        `bson:"ci"`
	SpeedWithoutArmor *[]string     `bson:"speed_without_armor"`
	SneakAttack       *[]string     `bson:"sneak_attack"`
	SorceryPoints     *[]int        `bson:"sorcery_points"`
	Metamagics        *[]metamagic  `bson:"metamagic"`
	SpellSlots        *[]int        `bson:"spell_slots"`
	SpellSlotsLevel   *[]int        `bson:"spell_slots_level"`

	IsSpellcaster   bool `bson:"is_spellcaster"`
	IsFullCaster    bool `bson:"is_full_caster"`
	IsHalfCaster    bool `bson:"is_half_caster"`
	IsQuarterCaster bool `bson:"is_quarter_caster"`
}

type proficiency struct {
	Armor        []string `bson:"armor"`
	Weapon       []string `bson:"weapon"`
	Tools        []choice `bson:"tools"`
	SavingThrows []string `json:"saving_throws"`
	Skills       skills   `json:"skills"`
	Equipment    []choice `bson:"equipment"`
}

type choice struct {
	Choice  bool     `bson:"choice"`
	Options []string `bson:"options"`
}

type skills struct {
	NumAvailable int      `bson:"num_available"`
	Available    []string `bson:"available"`
}

type feature struct {
	Level       []int  `bson:"level"`
	Name        string `bson:"name"`
	Source      string `bson:"source"`
	Subclass    string `bson:"subclass"`
	Description string `bson:"description"`
}

type invocation struct {
	Name         string `bson:"name"`
	Requirements string `json:"requirements"`
	Description  string `json:"description"`
}

type metamagic struct {
	Name        string `bson:"name"`
	Cost        int    `bson:"cost"`
	Description string `bson:"description"`
}
