package aa

import (
	"fmt"
	"os"
	"strings"
)

func generateWeb(aa *AAYaml) error {
	page := `---
# NOTE: this page is auto generated by github.com/rebuildeq/scripts !!
title: "Custom AA List"
layout: "aa"
# url: "/archives"
summary: "List of Custom AAs available on RebuildEQ"
---

Rebuild EQ hand crafts all the AAs for a custom experience. AAs are not gained in the classic sense, instead granted when you level up or by completing achievements.`

	pages := map[string][]string{}

	for _, skill := range aa.Skills {
		for _, rank := range skill.Ranks {
			key := fmt.Sprintf("%d^%s^%s", skill.Type, categoryConvert(skill.Type), classConvert(skill.Classes))
			icon := ""
			if skill.Icon != "" {
				icon = fmt.Sprintf("{{< xa id=\"%s\" >}} ", skill.Icon)
			}
			pages[key] = append(pages[key], fmt.Sprintf("%s%s|%s|%s|%d|%d|%s\n", icon, strings.TrimPrefix(skill.Name, "TODO "), categoryConvert(skill.Type), classConvert(skill.Classes), len(skill.Ranks), rank.Cost, rank.Description))
			//page += fmt.Sprintf("%s|%s|%s|%s\n", skill.Name, categoryConvert(skill.Type), classConvert(skill.Classes), rank.Description)
			break
		}
	}

	for key, skills := range pages {
		info := strings.Split(key, "^")
		category := info[1]
		class := info[2]
		page += fmt.Sprintf("\n\n### %s AAs - %s\n\n", category, class)
		page += "Name|Category|Class|Ranks|Cost|Description\n"
		page += "----|--------|-----|-----|----|-----------\n"
		page += strings.Join(skills, "")
	}

	err := os.WriteFile("aa.md", []byte(page), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func categoryConvert(in int) string {
	switch in {
	case 1:
		return "General"
	case 2:
		return "Archetype"
	case 3:
		return "Class"
	case 4:
		return "Achievement"
	}
	return "Unknown"
}

func classConvert(in int) string {

	if in&65535 == 65535 {
		return "All"
	}

	if in&1 == 1 {
		return "Warrior"
	}
	if in&2 == 2 {
		return "Cleric"
	}
	if in&4 == 4 {
		return "Paladin"
	}
	if in&8 == 8 {
		return "Ranger"
	}
	if in&16 == 16 {
		return "Shadow Knight"
	}
	if in&32 == 32 {
		return "Druid"
	}
	if in&64 == 64 {
		return "Monk"
	}
	if in&128 == 128 {
		return "Bard"
	}
	if in&256 == 256 {
		return "Rogue"
	}
	if in&512 == 512 {
		return "Shaman"
	}
	if in&1024 == 1024 {
		return "Necromancer"
	}
	if in&2048 == 2048 {
		return "Wizard"
	}
	if in&4096 == 4096 {
		return "Magician"
	}
	if in&8192 == 8192 {
		return "Enchanter"
	}
	if in&16384 == 16384 {
		return "Beastlord"
	}
	if in&32768 == 32768 {
		return "Berserker"
	}

	return "Unknown"
}
