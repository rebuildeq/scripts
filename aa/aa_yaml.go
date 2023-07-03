package aa

import "fmt"

type AAYaml struct {
	Skills []struct {
		Name            string `yaml:"aa_name"`
		NameSID         int    `yaml:"aa_name_sid"`
		ID              int    `yaml:"id"`
		Type            int    `yaml:"type"`
		Classes         int    `yaml:"classes"`
		Races           int    `yaml:"races"` // if 0, default to 65535
		GrantOnly       int    `yaml:"grant_only"`
		DrakkinHeritage int    `yaml:"drakkin_heritage"`
		Ranks           []struct {
			Index          int    `yaml:"index"`
			ID             int    `yaml:"id"`
			Title          string `yaml:"title"`
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
