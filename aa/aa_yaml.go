package aa

import "fmt"

type AAYaml struct {
	Skills []*AASkill `yaml:"skills,omitempty"`
}

type AASkill struct {
	Name            string    `yaml:"aa_name,omitempty"`
	NameSID         int       `yaml:"aa_name_sid,omitempty"`
	ID              int       `yaml:"id,omitempty"`
	Type            int       `yaml:"type,omitempty"`
	Classes         int       `yaml:"classes,omitempty"`
	Races           int       `yaml:"races,omitempty"` // if 0, default to 65535
	GrantOnly       int       `yaml:"grant_only,omitempty"`
	DrakkinHeritage int       `yaml:"drakkin_heritage,omitempty"`
	Ranks           []*AARank `yaml:"ranks,omitempty"`
	FirstRankID     int       `yaml:"first_rank_id,omitempty"`
}

type AARank struct {
	Index          int    `yaml:"index,omitempty"`
	ID             int    `yaml:"id,omitempty"`
	Title          string `yaml:"title,omitempty"`
	TitleSID       int    `yaml:"title_sid,omitempty"`
	Description    string `yaml:"description,omitempty"`
	DescriptionSID int    `yaml:"desc_sid,omitempty"`
	Cost           int    `yaml:"cost,omitempty"`
	LevelReq       int    `yaml:"level_req,omitempty"`
	UpperHotkeySID int    `yaml:"upper_hotkey_sid,omitempty"`
	LowerHotkeySID int    `yaml:"lower_hotkey_sid,omitempty"`
	SpellID        int    `yaml:"spell_id,omitempty"`
	SpellType      int    `yaml:"spell_type,omitempty"`
	RecastTime     int    `yaml:"recast_time,omitempty"`
	Expansion      int    `yaml:"expansion,omitempty"`
	PrevID         int    `yaml:"prev_id,omitempty"`
	NextID         int    `yaml:"next_id,omitempty"`
	Slot1          struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2,omitempty"`
	} `yaml:"slot1,omitempty"`
	Slot2 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2,omitempty"`
	} `yaml:"slot2,omitempty"`
	Slot3 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2,omitempty"`
	} `yaml:"slot3,omitempty"`
	Slot4 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2,omitempty"`
	} `yaml:"slot4,omitempty"`
	Slot5 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2,omitempty"`
	} `yaml:"slot5,omitempty"`
	Slot6 struct {
		EffectID int `yaml:"effect_id,omitempty"`
		Base1    int `yaml:"base1,omitempty"`
		Base2    int `yaml:"base2,omitempty"`
	} `yaml:"slot6,omitempty"`
	nextID int
}

func (e *AAYaml) sanitize() error {

	abilityNames := make(map[int]string)
	titleNames := make(map[int]string)
	descNames := make(map[int]string)

	uniqueRankIDs := make(map[int]bool)

	for skillIndex, skill := range e.Skills {
		abilityName, ok := abilityNames[skill.ID]
		if !ok {
			abilityNames[skill.ID] = skill.Name
			abilityName = skill.Name
		}
		if abilityName != skill.Name {
			return fmt.Errorf("ability name mismatch for skill id %d between '%s' and '%s'", skill.ID, abilityName, skill.Name)
		}
		titleName, ok := titleNames[skill.NameSID]
		if !ok {
			titleNames[skill.NameSID] = skill.Name
			titleName = skill.Name
		}
		if titleName != skill.Name {
			return fmt.Errorf("title name mismatch for nameSID %d for skill id %d between '%s' and '%s'", skill.NameSID, skillIndex, titleName, skill.Name)
		}
		for rankIndex, rank := range skill.Ranks {
			if rank.ID == 0 {
				return fmt.Errorf("rank id is 0 for skill id %d rank %d", skillIndex, rankIndex)
			}
			_, ok := uniqueRankIDs[rank.ID]
			if ok {
				return fmt.Errorf("duplicate rank id %d for skill id %d rank %d", rank.ID, skillIndex, rankIndex)
			}
			uniqueRankIDs[rank.ID] = true
			if rank.TitleSID != 0 {
				titleName, ok := titleNames[rank.TitleSID]
				if !ok {
					titleNames[rank.TitleSID] = rank.Title
					titleName = rank.Title
				}
				if titleName != rank.Title {
					return fmt.Errorf("title name mismatch for titleSID %d skill id %d rank %d between '%s' and '%s'", rank.TitleSID, skillIndex, rankIndex, titleName, rank.Title)
				}
			}

			descName, ok := descNames[rank.DescriptionSID]
			if !ok {
				descNames[rank.DescriptionSID] = rank.Description
				descName = rank.Description
			}
			if descName != rank.Description {
				return fmt.Errorf("description name mismatch for descriptionSID %d skill id %d rank %d between '%s' and '%s'", rank.DescriptionSID, skillIndex, rankIndex, descName, rank.Description)
			}
		}
	}
	return nil
}
