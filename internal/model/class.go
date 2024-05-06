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

	KnownCantrips     *[]int        `bson:"known_cantrips,omitempty" json:"known_cantrips,omitempty"`
	KnownSpells       *[]int        `bson:"known_spells,omitempty" json:"known_spells,omitempty"`
	KnownInvocation   *[]int        `bson:"known_invocation,omitempty" json:"known_invocation,omitempty"`
	Invocations       *[]invocation `bson:"invocations,omitempty" json:"invocations,omitempty"`
	Rage              *[]string     `bson:"rage,omitempty" json:"rage,omitempty"`
	MartialArts       *[]string     `bson:"martial_arts,omitempty" json:"martial_arts,omitempty"`
	Ci                *[]int        `bson:"ci,omitempty" json:"ci,omitempty"`
	SpeedWithoutArmor *[]string     `bson:"speed_without_armor,omitempty" json:"speed_without_armor,omitempty"`
	SneakAttack       *[]string     `bson:"sneak_attack,omitempty" json:"sneak_attack,omitempty"`
	SorceryPoints     *[]int        `bson:"sorcery_points,omitempty" json:"sorcery_points,omitempty"`
	Metamagics        *[]metamagic  `bson:"metamagic,omitempty" json:"metamagic,omitempty"`
	SpellSlots        *[]int        `bson:"spell_slots,omitempty" json:"spell_slots,omitempty"`
	SpellSlotsLevel   *[]int        `bson:"spell_slots_level,omitempty" json:"spell_slots_level,omitempty"`

	IsSpellcaster   bool `bson:"is_spellcaster"`
	IsFullCaster    bool `bson:"is_full_caster"`
	IsHalfCaster    bool `bson:"is_half_caster"`
	IsQuarterCaster bool `bson:"is_quarter_caster"`
}

type proficiency struct {
	Armor        []string `bson:"armor"`
	Weapon       []string `bson:"weapon"`
	Tools        []choice `bson:"tools"`
	SavingThrows []string `bson:"saving_throws"`
	Skills       skills   `bson:"skills"`
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
	Requirements string `bson:"requirements"`
	Description  string `bson:"description"`
}

type metamagic struct {
	Name        string `bson:"name"`
	Cost        int    `bson:"cost"`
	Description string `bson:"description"`
}
