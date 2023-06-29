package aa

import "fmt"

type AAYaml struct {
	Skills []struct {
		ShortName string `yaml:"short_name"`
		Name      string `yaml:"aa_name"`
		ID        int    `yaml:"id"`
		Type      int    `yaml:"type"`
		Classes   int    `yaml:"classes"`
		Races     int    `yaml:"races"` // if 0, default to 65535
		GrantOnly int    `yaml:"grant_only"`
		Ranks     []struct {
			Index          int    `yaml:"index"`
			Name           string `yaml:"name"`
			ID             int    `yaml:"id"`
			TitleSID       int    `yaml:"title_sid"`
			Description    string `yaml:"description"`
			DescriptionSID int    `yaml:"desc_sid"`
			Cost           int    `yaml:"cost"`
			LevelReq       int    `yaml:"level_req"`
			Slot1          struct {
				EffectID int `yaml:"effect_id"`
				Base1    int `yaml:"base1"`
				Base2    int `yaml:"base2"`
			} `yaml:"slot1"`
			Slot2 struct {
				EffectID int `yaml:"effect_id"`
				Base1    int `yaml:"base1"`
				Base2    int `yaml:"base2"`
			} `yaml:"slot2"`
			Slot3 struct {
				EffectID int `yaml:"effect_id"`
				Base1    int `yaml:"base1"`
				Base2    int `yaml:"base2"`
			} `yaml:"slot3"`
			Slot4 struct {
				EffectID int `yaml:"effect_id"`
				Base1    int `yaml:"base1"`
				Base2    int `yaml:"base2"`
			} `yaml:"slot4"`
			Slot5 struct {
				EffectID int `yaml:"effect_id"`
				Base1    int `yaml:"base1"`
				Base2    int `yaml:"base2"`
			} `yaml:"slot5"`
			Slot6 struct {
				EffectID int `yaml:"effect_id"`
				Base1    int `yaml:"base1"`
				Base2    int `yaml:"base2"`
			} `yaml:"slot6"`
		} `yaml:"ranks"`
	} `yaml:"skills"`
}

func (e *AAYaml) sanitize() error {

	abilityNames := make(map[int]string)
	titleNames := make(map[int]string)
	descNames := make(map[int]string)

	for skillIndex, skill := range e.Skills {
		abilityName, ok := abilityNames[skill.ID]
		if !ok {
			abilityNames[skill.ID] = skill.Name
			abilityName = skill.Name
		}
		if abilityName != skill.Name {
			return fmt.Errorf("ability name mismatch for skill id %d between '%s' and '%s'", skill.ID, abilityName, skill.Name)
		}
		for rankIndex, rank := range skill.Ranks {
			titleName, ok := titleNames[rank.TitleSID]
			if !ok {
				titleNames[rank.TitleSID] = rank.Name
				titleName = rank.Name
			}
			if titleName != rank.Name {
				return fmt.Errorf("title name mismatch for titleSID %d for skill id %d rank %d between '%s' and '%s'", rank.TitleSID, skillIndex, rankIndex, titleName, rank.Name)
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
