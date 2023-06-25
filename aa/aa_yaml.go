package aa

type AAYaml struct {
	Skills []struct {
		ShortName string `yaml:"short_name"`
		Name      string `yaml:"aa_name"`
		ID        int    `yaml:"id"`
		Type      int    `yaml:"type"`
		Classes   int    `yaml:"classes"`
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
		} `yaml:"ranks"`
	} `yaml:"skills"`
}
