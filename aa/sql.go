package aa

import (
	"fmt"
	"os"
)

func generateAASQL(aa *AAYaml) error {
	w, err := os.Create("bin/aa.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("INSERT INTO aa_ability (id, `name`, category, classes, races, drakkin_heritage, deities, `status`, `type`, charges, grant_only, first_rank_id, `enabled`, reset_on_death) VALUES\n")
	for skillIndex, skill := range aa.Skills {
		if len(skill.Ranks) == 0 {
			continue
		}
		w.WriteString(fmt.Sprintf("	(%d, '%s', -1, %d, 65535, 127, 131071, 0, %d, 0, %d, %d, 1, 0)", skill.ID, skill.Name, skill.Classes, skill.Type, skill.GrantOnly, skill.Ranks[0].ID))
		if skillIndex == len(aa.Skills)-1 {
			w.WriteString(";\n\n")
		} else {
			w.WriteString(",\n")
		}
	}

	w.WriteString("INSERT INTO aa_rank_effects (rank_id, slot, effect_id, base1, base2) VALUES\n")
	for skillIndex, skill := range aa.Skills {
		for rankIndex, rank := range skill.Ranks {
			w.WriteString(fmt.Sprintf("	(%d, 1, %d, %d, %d)", rank.ID, rank.Slot1.EffectID, rank.Slot1.Base1, rank.Slot1.Base2))
			if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
				w.WriteString(";\n\n")
			} else {
				w.WriteString(",\n")
			}
		}
	}
	w.WriteString("INSERT INTO aa_ranks (id, upper_hotkey_sid, lower_hotkey_sid, title_sid, desc_sid, cost, level_req, spell, spell_type, recast_time, expansion, prev_id, next_id) VALUES\n")
	for skillIndex, skill := range aa.Skills {
		for rankIndex, rank := range skill.Ranks {
			prevID := -1
			nextID := -1
			if rankIndex > 0 {
				prevID = skill.Ranks[rankIndex-1].ID
			}
			if rankIndex < len(skill.Ranks)-1 {
				nextID = skill.Ranks[rankIndex+1].ID
			}
			w.WriteString(fmt.Sprintf("	(%d, -1, -1, %d, %d, %d, %d, 0, 0, 0, 0, %d, %d)", rank.ID, rank.TitleSID, rank.DescriptionSID, rank.Cost, rank.LevelReq, prevID, nextID))
			if rankIndex == len(skill.Ranks)-1 && skillIndex == len(aa.Skills)-1 {
				w.WriteString(";\n\n")
			} else {
				w.WriteString(",\n")
			}
		}
	}

	return nil
}
