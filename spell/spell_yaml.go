package spell

import "fmt"

type SpellYaml struct {
	Spells []*struct {
		ID                   int    `yaml:"id"`                   // int(11) NOT NULL DEFAULT 0,
		Name                 string `yaml:"name"`                 // varchar(64) DEFAULT NULL,
		Player_1             string `yaml:"player_1"`             // varchar(64) DEFAULT 'BLUE_TRAIL',
		Teleport_zone        int    `yaml:"teleport_zone"`        // varchar(64) DEFAULT NULL,
		You_cast             int    `yaml:"you_cast"`             // varchar(120) DEFAULT NULL,
		Other_casts          int    `yaml:"other_casts"`          // varchar(120) DEFAULT NULL,
		Cast_on_you          int    `yaml:"cast_on_you"`          // varchar(120) DEFAULT NULL,
		Cast_on_other        int    `yaml:"cast_on_other"`        // varchar(120) DEFAULT NULL,
		Spell_fades          int    `yaml:"spell_fades"`          // varchar(120) DEFAULT NULL,
		Range                int    `yaml:"range"`                // int(11) NOT NULL DEFAULT 100,
		Aoerange             int    `yaml:"aoerange"`             // int(11) NOT NULL DEFAULT 0,
		Pushback             int    `yaml:"pushback"`             // int(11) NOT NULL DEFAULT 0,
		Pushup               int    `yaml:"pushup"`               // int(11) NOT NULL DEFAULT 0,
		Cast_time            int    `yaml:"cast_time"`            // int(11) NOT NULL DEFAULT 0,
		Recovery_time        int    `yaml:"recovery_time"`        // int(11) NOT NULL DEFAULT 0,
		Recast_time          int    `yaml:"recast_time"`          // int(11) NOT NULL DEFAULT 0,
		Buffdurationformula  int    `yaml:"buffdurationformula"`  // int(11) NOT NULL DEFAULT 7,
		Buffduration         int    `yaml:"buffduration"`         // int(11) NOT NULL DEFAULT 65,
		AEDuration           int    `yaml:"AEDuration"`           // int(11) NOT NULL DEFAULT 0,
		Mana                 int    `yaml:"mana"`                 // int(11) NOT NULL DEFAULT 0,
		Effect_base_value1   int    `yaml:"effect_base_value1"`   // int(11) NOT NULL DEFAULT 100,
		Effect_base_value2   int    `yaml:"effect_base_value2"`   // int(11) NOT NULL DEFAULT 0,
		Effect_base_value3   int    `yaml:"effect_base_value3"`   // int(11) NOT NULL DEFAULT 0,
		Effect_base_value4   int    `yaml:"effect_base_value4"`   // int(11) NOT NULL DEFAULT 0,
		Effect_base_value5   int    `yaml:"effect_base_value5"`   // int(11) NOT NULL DEFAULT 0,
		Effect_base_value6   int    `yaml:"effect_base_value6"`   // int(11) NOT NULL DEFAULT 0,
		Effect_base_value7   int    `yaml:"effect_base_value7"`   // int(11) NOT NULL DEFAULT 0,
		Effect_base_value8   int    `yaml:"effect_base_value8"`   // int(11) NOT NULL DEFAULT 0,
		Effect_base_value9   int    `yaml:"effect_base_value9"`   // int(11) NOT NULL DEFAULT 0,
		Effect_base_value10  int    `yaml:"effect_base_value10"`  // int(11) NOT NULL DEFAULT 0,
		Effect_base_value11  int    `yaml:"effect_base_value11"`  // int(11) NOT NULL DEFAULT 0,
		Effect_base_value12  int    `yaml:"effect_base_value12"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value1  int    `yaml:"effect_limit_value1"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value2  int    `yaml:"effect_limit_value2"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value3  int    `yaml:"effect_limit_value3"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value4  int    `yaml:"effect_limit_value4"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value5  int    `yaml:"effect_limit_value5"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value6  int    `yaml:"effect_limit_value6"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value7  int    `yaml:"effect_limit_value7"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value8  int    `yaml:"effect_limit_value8"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value9  int    `yaml:"effect_limit_value9"`  // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value10 int    `yaml:"effect_limit_value10"` // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value11 int    `yaml:"effect_limit_value11"` // int(11) NOT NULL DEFAULT 0,
		Effect_limit_value12 int    `yaml:"effect_limit_value12"` // int(11) NOT NULL DEFAULT 0,
		Max1                 int    `yaml:"max1"`                 // int(11) NOT NULL DEFAULT 0,
		Max2                 int    `yaml:"max2"`                 // int(11) NOT NULL DEFAULT 0,
		Max3                 int    `yaml:"max3"`                 // int(11) NOT NULL DEFAULT 0,
		Max4                 int    `yaml:"max4"`                 // int(11) NOT NULL DEFAULT 0,
		Max5                 int    `yaml:"max5"`                 // int(11) NOT NULL DEFAULT 0,
		Max6                 int    `yaml:"max6"`                 // int(11) NOT NULL DEFAULT 0,
		Max7                 int    `yaml:"max7"`                 // int(11) NOT NULL DEFAULT 0,
		Max8                 int    `yaml:"max8"`                 // int(11) NOT NULL DEFAULT 0,
		Max9                 int    `yaml:"max9"`                 // int(11) NOT NULL DEFAULT 0,
		Max10                int    `yaml:"max10"`                // int(11) NOT NULL DEFAULT 0,
		Max11                int    `yaml:"max11"`                // int(11) NOT NULL DEFAULT 0,
		Max12                int    `yaml:"max12"`                // int(11) NOT NULL DEFAULT 0,
		Icon                 int    `yaml:"icon"`                 // int(11) NOT NULL DEFAULT 0,
		Memicon              int    `yaml:"memicon"`              // int(11) NOT NULL DEFAULT 0,
		Components1          int    `yaml:"components1"`          // int(11) NOT NULL DEFAULT -1,
		Components2          int    `yaml:"components2"`          // int(11) NOT NULL DEFAULT -1,
		Components3          int    `yaml:"components3"`          // int(11) NOT NULL DEFAULT -1,
		Components4          int    `yaml:"components4"`          // int(11) NOT NULL DEFAULT -1,
		Component_counts1    int    `yaml:"component_counts1"`    // int(11) NOT NULL DEFAULT 1,
		Component_counts2    int    `yaml:"component_counts2"`    // int(11) NOT NULL DEFAULT 1,
		Component_counts3    int    `yaml:"component_counts3"`    // int(11) NOT NULL DEFAULT 1,
		Component_counts4    int    `yaml:"component_counts4"`    // int(11) NOT NULL DEFAULT 1,
		NoexpendReagent1     int    `yaml:"NoexpendReagent1"`     // int(11) NOT NULL DEFAULT -1,
		NoexpendReagent2     int    `yaml:"NoexpendReagent2"`     // int(11) NOT NULL DEFAULT -1,
		NoexpendReagent3     int    `yaml:"NoexpendReagent3"`     // int(11) NOT NULL DEFAULT -1,
		NoexpendReagent4     int    `yaml:"NoexpendReagent4"`     // int(11) NOT NULL DEFAULT -1,
		Formula1             int    `yaml:"formula1"`             // int(11) NOT NULL DEFAULT 100,
		Formula2             int    `yaml:"formula2"`             // int(11) NOT NULL DEFAULT 100,
		Formula3             int    `yaml:"formula3"`             // int(11) NOT NULL DEFAULT 100,
		Formula4             int    `yaml:"formula4"`             // int(11) NOT NULL DEFAULT 100,
		Formula5             int    `yaml:"formula5"`             // int(11) NOT NULL DEFAULT 100,
		Formula6             int    `yaml:"formula6"`             // int(11) NOT NULL DEFAULT 100,
		Formula7             int    `yaml:"formula7"`             // int(11) NOT NULL DEFAULT 100,
		Formula8             int    `yaml:"formula8"`             // int(11) NOT NULL DEFAULT 100,
		Formula9             int    `yaml:"formula9"`             // int(11) NOT NULL DEFAULT 100,
		Formula10            int    `yaml:"formula10"`            // int(11) NOT NULL DEFAULT 100,
		Formula11            int    `yaml:"formula11"`            // int(11) NOT NULL DEFAULT 100,
		Formula12            int    `yaml:"formula12"`            // int(11) NOT NULL DEFAULT 100,
		LightType            int    `yaml:"LightType"`            // int(11) NOT NULL DEFAULT 0,
		GoodEffect           int    `yaml:"goodEffect"`           // int(11) NOT NULL DEFAULT 0,
		Activated            int    `yaml:"Activated"`            // int(11) NOT NULL DEFAULT 0,
		Resisttype           int    `yaml:"resisttype"`           // int(11) NOT NULL DEFAULT 0,
		Effectid1            int    `yaml:"effectid1"`            // int(11) NOT NULL DEFAULT 254,
		Effectid2            int    `yaml:"effectid2"`            // int(11) NOT NULL DEFAULT 254,
		Effectid3            int    `yaml:"effectid3"`            // int(11) NOT NULL DEFAULT 254,
		Effectid4            int    `yaml:"effectid4"`            // int(11) NOT NULL DEFAULT 254,
		Effectid5            int    `yaml:"effectid5"`            // int(11) NOT NULL DEFAULT 254,
		Effectid6            int    `yaml:"effectid6"`            // int(11) NOT NULL DEFAULT 254,
		Effectid7            int    `yaml:"effectid7"`            // int(11) NOT NULL DEFAULT 254,
		Effectid8            int    `yaml:"effectid8"`            // int(11) NOT NULL DEFAULT 254,
		Effectid9            int    `yaml:"effectid9"`            // int(11) NOT NULL DEFAULT 254,
		Effectid10           int    `yaml:"effectid10"`           // int(11) NOT NULL DEFAULT 254,
		Effectid11           int    `yaml:"effectid11"`           // int(11) NOT NULL DEFAULT 254,
		Effectid12           int    `yaml:"effectid12"`           // int(11) NOT NULL DEFAULT 254,
		Targettype           int    `yaml:"targettype"`           // int(11) NOT NULL DEFAULT 2,
		Basediff             int    `yaml:"basediff"`             // int(11) NOT NULL DEFAULT 0,
		Skill                int    `yaml:"skill"`                // int(11) NOT NULL DEFAULT 98,
		Zonetype             int    `yaml:"zonetype"`             // int(11) NOT NULL DEFAULT -1,
		EnvironmentType      int    `yaml:"EnvironmentType"`      // int(11) NOT NULL DEFAULT 0,
		TimeOfDay            int    `yaml:"TimeOfDay"`            // int(11) NOT NULL DEFAULT 0,
		Classes1             int    `yaml:"classes1"`             // int(11) NOT NULL DEFAULT 255,
		Classes2             int    `yaml:"classes2"`             // int(11) NOT NULL DEFAULT 255,
		Classes3             int    `yaml:"classes3"`             // int(11) NOT NULL DEFAULT 255,
		Classes4             int    `yaml:"classes4"`             // int(11) NOT NULL DEFAULT 255,
		Classes5             int    `yaml:"classes5"`             // int(11) NOT NULL DEFAULT 255,
		Classes6             int    `yaml:"classes6"`             // int(11) NOT NULL DEFAULT 255,
		Classes7             int    `yaml:"classes7"`             // int(11) NOT NULL DEFAULT 255,
		Classes8             int    `yaml:"classes8"`             // int(11) NOT NULL DEFAULT 255,
		Classes9             int    `yaml:"classes9"`             // int(11) NOT NULL DEFAULT 255,
		Classes10            int    `yaml:"classes10"`            // int(11) NOT NULL DEFAULT 255,
		Classes11            int    `yaml:"classes11"`            // int(11) NOT NULL DEFAULT 255,
		Classes12            int    `yaml:"classes12"`            // int(11) NOT NULL DEFAULT 255,
		Classes13            int    `yaml:"classes13"`            // int(11) NOT NULL DEFAULT 255,
		Classes14            int    `yaml:"classes14"`            // int(11) NOT NULL DEFAULT 255,
		Classes15            int    `yaml:"classes15"`            // int(11) NOT NULL DEFAULT 255,
		Classes16            int    `yaml:"classes16"`            // int(11) NOT NULL DEFAULT 255,
		CastingAnim          int    `yaml:"CastingAnim"`          // int(11) NOT NULL DEFAULT 44,
		TargetAnim           int    `yaml:"TargetAnim"`           // int(11) NOT NULL DEFAULT 13,
		TravelType           int    `yaml:"TravelType"`           // int(11) NOT NULL DEFAULT 0,
		SpellAffectIndex     int    `yaml:"SpellAffectIndex"`     // int(11) NOT NULL DEFAULT -1,
		Disallow_sit         int    `yaml:"disallow_sit"`         // int(11) NOT NULL DEFAULT 0,
		Deities0             int    `yaml:"deities0"`             // int(11) NOT NULL DEFAULT 0,
		Deities1             int    `yaml:"deities1"`             // int(11) NOT NULL DEFAULT 0,
		Deities2             int    `yaml:"deities2"`             // int(11) NOT NULL DEFAULT 0,
		Deities3             int    `yaml:"deities3"`             // int(11) NOT NULL DEFAULT 0,
		Deities4             int    `yaml:"deities4"`             // int(11) NOT NULL DEFAULT 0,
		Deities5             int    `yaml:"deities5"`             // int(11) NOT NULL DEFAULT 0,
		Deities6             int    `yaml:"deities6"`             // int(11) NOT NULL DEFAULT 0,
		Deities7             int    `yaml:"deities7"`             // int(11) NOT NULL DEFAULT 0,
		Deities8             int    `yaml:"deities8"`             // int(11) NOT NULL DEFAULT 0,
		Deities9             int    `yaml:"deities9"`             // int(11) NOT NULL DEFAULT 0,
		Deities10            int    `yaml:"deities10"`            // int(11) NOT NULL DEFAULT 0,
		Deities11            int    `yaml:"deities11"`            // int(11) NOT NULL DEFAULT 0,
		Deities12            int    `yaml:"deities12"`            // int(12) NOT NULL DEFAULT 0,
		Deities13            int    `yaml:"deities13"`            // int(11) NOT NULL DEFAULT 0,
		Deities14            int    `yaml:"deities14"`            // int(11) NOT NULL DEFAULT 0,
		Deities15            int    `yaml:"deities15"`            // int(11) NOT NULL DEFAULT 0,
		Deities16            int    `yaml:"deities16"`            // int(11) NOT NULL DEFAULT 0,
		Field142             int    `yaml:"field142"`             // int(11) NOT NULL DEFAULT 100,
		Field143             int    `yaml:"field143"`             // int(11) NOT NULL DEFAULT 0,
		New_icon             int    `yaml:"new_icon"`             // int(11) NOT NULL DEFAULT 161,
		Spellanim            int    `yaml:"spellanim"`            // int(11) NOT NULL DEFAULT 0,
		Uninterruptable      int    `yaml:"uninterruptable"`      // int(11) NOT NULL DEFAULT 0,
		ResistDiff           int    `yaml:"ResistDiff"`           // int(11) NOT NULL DEFAULT -150,
		Dot_stacking_exempt  int    `yaml:"dot_stacking_exempt"`  // int(11) NOT NULL DEFAULT 0,
		Deleteable           int    `yaml:"deleteable"`           // int(11) NOT NULL DEFAULT 0,
		RecourseLink         int    `yaml:"RecourseLink"`         // int(11) NOT NULL DEFAULT 0,
		No_partial_resist    int    `yaml:"no_partial_resist"`    // int(11) NOT NULL DEFAULT 0,
		Field152             int    `yaml:"field152"`             // int(11) NOT NULL DEFAULT 0,
		Field153             int    `yaml:"field153"`             // int(11) NOT NULL DEFAULT 0,
		Short_buff_box       int    `yaml:"short_buff_box"`       // int(11) NOT NULL DEFAULT -1,
		Descnum              int    `yaml:"descnum"`              // int(11) NOT NULL DEFAULT 0,
		Typedescnum          int    `yaml:"typedescnum"`          // int(11) DEFAULT NULL,
		Effectdescnum        int    `yaml:"effectdescnum"`        // int(11) DEFAULT NULL,
		Effectdescnum2       int    `yaml:"effectdescnum2"`       // int(11) NOT NULL DEFAULT 0,
		Npc_no_los           int    `yaml:"npc_no_los"`           // int(11) NOT NULL DEFAULT 0,
		Field160             int    `yaml:"field160"`             // int(11) NOT NULL DEFAULT 0,
		Reflectable          int    `yaml:"reflectable"`          // int(11) NOT NULL DEFAULT 0,
		Bonushate            int    `yaml:"bonushate"`            // int(11) NOT NULL DEFAULT 0,
		Field163             int    `yaml:"field163"`             // int(11) NOT NULL DEFAULT 100,
		Field164             int    `yaml:"field164"`             // int(11) NOT NULL DEFAULT -150,
		Ldon_trap            int    `yaml:"ldon_trap"`            // int(11) NOT NULL DEFAULT 0,
		EndurCost            int    `yaml:"EndurCost"`            // int(11) NOT NULL DEFAULT 0,
		EndurTimerIndex      int    `yaml:"EndurTimerIndex"`      // int(11) NOT NULL DEFAULT 0,
		IsDiscipline         int    `yaml:"IsDiscipline"`         // int(11) NOT NULL DEFAULT 0,
		Field169             int    `yaml:"field169"`             // int(11) NOT NULL DEFAULT 0,
		Field170             int    `yaml:"field170"`             // int(11) NOT NULL DEFAULT 0,
		Field171             int    `yaml:"field171"`             // int(11) NOT NULL DEFAULT 0,
		Field172             int    `yaml:"field172"`             // int(11) NOT NULL DEFAULT 0,
		HateAdded            int    `yaml:"HateAdded"`            // int(11) NOT NULL DEFAULT 0,
		EndurUpkeep          int    `yaml:"EndurUpkeep"`          // int(11) NOT NULL DEFAULT 0,
		Numhitstype          int    `yaml:"numhitstype"`          // int(11) NOT NULL DEFAULT 0,
		Numhits              int    `yaml:"numhits"`              // int(11) NOT NULL DEFAULT 0,
		Pvpresistbase        int    `yaml:"pvpresistbase"`        // int(11) NOT NULL DEFAULT -150,
		Pvpresistcalc        int    `yaml:"pvpresistcalc"`        // int(11) NOT NULL DEFAULT 100,
		Pvpresistcap         int    `yaml:"pvpresistcap"`         // int(11) NOT NULL DEFAULT -150,
		Spell_category       int    `yaml:"spell_category"`       // int(11) NOT NULL DEFAULT -99,
		Pvp_duration         int    `yaml:"pvp_duration"`         // int(11) NOT NULL DEFAULT 0,
		Pvp_duration_cap     int    `yaml:"pvp_duration_cap"`     // int(11) NOT NULL DEFAULT 0,
		Pcnpc_only_flag      int    `yaml:"pcnpc_only_flag"`      // int(11) DEFAULT 0,
		Cast_not_standing    int    `yaml:"cast_not_standing"`    // int(11) DEFAULT 0,
		Can_mgb              int    `yaml:"can_mgb"`              // int(11) NOT NULL DEFAULT 0,
		Nodispell            int    `yaml:"nodispell"`            // int(11) NOT NULL DEFAULT -1,
		Npc_category         int    `yaml:"npc_category"`         // int(11) NOT NULL DEFAULT 0,
		Npc_usefulness       int    `yaml:"npc_usefulness"`       // int(11) NOT NULL DEFAULT 0,
		MinResist            int    `yaml:"MinResist"`            // int(11) NOT NULL DEFAULT 0,
		MaxResist            int    `yaml:"MaxResist"`            // int(11) NOT NULL DEFAULT 0,
		Viral_targets        int    `yaml:"viral_targets"`        // int(11) NOT NULL DEFAULT 0,
		Viral_timer          int    `yaml:"viral_timer"`          // int(11) NOT NULL DEFAULT 0,
		Nimbuseffect         int    `yaml:"nimbuseffect"`         // int(11) DEFAULT 0,
		ConeStartAngle       int    `yaml:"ConeStartAngle"`       // int(11) NOT NULL DEFAULT 0,
		ConeStopAngle        int    `yaml:"ConeStopAngle"`        // int(11) NOT NULL DEFAULT 0,
		Sneaking             int    `yaml:"sneaking"`             // int(11) NOT NULL DEFAULT 0,
		Not_extendable       int    `yaml:"not_extendable"`       // int(11) NOT NULL DEFAULT 0,
		Field198             int    `yaml:"field198"`             // int(11) NOT NULL DEFAULT 0,
		Field199             int    `yaml:"field199"`             // int(11) NOT NULL DEFAULT 1,
		Suspendable          int    `yaml:"suspendable"`          // int(11) DEFAULT 0,
		Viral_range          int    `yaml:"viral_range"`          // int(11) NOT NULL DEFAULT 0,
		Songcap              int    `yaml:"songcap"`              // int(11) DEFAULT 0,
		Field203             int    `yaml:"field203"`             // int(11) DEFAULT 0,
		Field204             int    `yaml:"field204"`             // int(11) DEFAULT 0,
		No_block             int    `yaml:"no_block"`             // int(11) NOT NULL DEFAULT 0,
		Field206             int    `yaml:"field206"`             // int(11) DEFAULT -1,
		Spellgroup           int    `yaml:"spellgroup"`           // int(11) DEFAULT 0,
		Rank                 int    `yaml:"rank"`                 // int(11) NOT NULL DEFAULT 0,
		Field209             int    `yaml:"field209"`             // int(11) DEFAULT 0,
		Field210             int    `yaml:"field210"`             // int(11) DEFAULT 1,
		CastRestriction      int    `yaml:"CastRestriction"`      // int(11) NOT NULL DEFAULT 0,
		Allowrest            int    `yaml:"allowrest"`            // int(11) DEFAULT 0,
		InCombat             int    `yaml:"InCombat"`             // int(11) NOT NULL DEFAULT 0,
		OutofCombat          int    `yaml:"OutofCombat"`          // int(11) NOT NULL DEFAULT 0,
		Field215             int    `yaml:"field215"`             // int(11) DEFAULT 0,
		Field216             int    `yaml:"field216"`             // int(11) DEFAULT 0,
		Field217             int    `yaml:"field217"`             // int(11) DEFAULT 0,
		Aemaxtargets         int    `yaml:"aemaxtargets"`         // int(11) NOT NULL DEFAULT 0,
		Maxtargets           int    `yaml:"maxtargets"`           // int(11) DEFAULT 0,
		Field220             int    `yaml:"field220"`             // int(11) DEFAULT 0,
		Field221             int    `yaml:"field221"`             // int(11) DEFAULT 0,
		Field222             int    `yaml:"field222"`             // int(11) DEFAULT 0,
		Field223             int    `yaml:"field223"`             // int(11) DEFAULT 0,
		Persistdeath         int    `yaml:"persistdeath"`         // int(11) DEFAULT 0,
		Field225             int    `yaml:"field225"`             // int(11) NOT NULL DEFAULT 0,
		Field226             int    `yaml:"field226"`             // int(11) NOT NULL DEFAULT 0,
		Min_dist             int    `yaml:"min_dist"`             // float NOT NULL DEFAULT 0,
		Min_dist_mod         int    `yaml:"min_dist_mod"`         // float NOT NULL DEFAULT 0,
		Max_dist             int    `yaml:"max_dist"`             // float NOT NULL DEFAULT 0,
		Max_dist_mod         int    `yaml:"max_dist_mod"`         // float NOT NULL DEFAULT 0,
		Min_range            int    `yaml:"min_range"`            // int(11) NOT NULL DEFAULT 0,
		Field232             int    `yaml:"field232"`             // int(11) NOT NULL DEFAULT 0,
		Field233             int    `yaml:"field233"`             // int(11) NOT NULL DEFAULT 0,
		Field234             int    `yaml:"field234"`             // int(11) NOT NULL DEFAULT 0,
		Field235             int    `yaml:"field235"`             // int(11) NOT NULL DEFAULT 0,
		Field236             int    `yaml:"field236"`             // int(11) NOT NULL DEFAULT 0,
	} `yaml:"spells"`
}

func (e *SpellYaml) sanitize() error {
	for _, spell := range e.Spells {

		if spell.ID < 1 {
			return fmt.Errorf("spell id must be greater than 0")
		}
		if spell.Name == "" {
			return fmt.Errorf("spell name must not be empty")
		}
		if spell.Player_1 == "" {
			spell.Player_1 = "BLUE_TRAIL"
		}

		//`teleport_zone` varchar(64) DEFAULT NULL,
		//`you_cast` varchar(120) DEFAULT NULL,
		//`other_casts` varchar(120) DEFAULT NULL,
		//`cast_on_you` varchar(120) DEFAULT NULL,
		//`cast_on_other` varchar(120) DEFAULT NULL,
		//`spell_fades` varchar(120) DEFAULT NULL,
		if spell.Range == 0 {
			spell.Range = 100
		}

		//`aoerange` int(11) NOT NULL DEFAULT 0,
		//`pushback` int(11) NOT NULL DEFAULT 0,
		//`pushup` int(11) NOT NULL DEFAULT 0,
		//`cast_time` int(11) NOT NULL DEFAULT 0,
		//`recovery_time` int(11) NOT NULL DEFAULT 0,
		//`recast_time` int(11) NOT NULL DEFAULT 0,

		if spell.Components1 == 0 {
			spell.Components1 = -1
		}
		if spell.Components2 == 0 {
			spell.Components2 = -1
		}
		if spell.Components3 == 0 {
			spell.Components3 = -1
		}
		if spell.Components4 == 0 {
			spell.Components4 = -1
		}
		if spell.NoexpendReagent1 == 0 {
			spell.NoexpendReagent1 = -1
		}
		if spell.NoexpendReagent2 == 0 {
			spell.NoexpendReagent2 = -1
		}
		if spell.NoexpendReagent3 == 0 {
			spell.NoexpendReagent3 = -1
		}
		if spell.NoexpendReagent4 == 0 {
			spell.NoexpendReagent4 = -1
		}
		if spell.Formula1 == 0 {
			spell.Formula1 = 100
		}
		if spell.Formula2 == 0 {
			spell.Formula2 = 100
		}
		if spell.Formula3 == 0 {
			spell.Formula3 = 100
		}
		if spell.Formula4 == 0 {
			spell.Formula4 = 100
		}
		if spell.Formula5 == 0 {
			spell.Formula5 = 100
		}
		if spell.Formula6 == 0 {
			spell.Formula6 = 100
		}
		if spell.Formula7 == 0 {
			spell.Formula7 = 100
		}
		if spell.Formula8 == 0 {
			spell.Formula8 = 100
		}
		if spell.Formula9 == 0 {
			spell.Formula9 = 100
		}
		if spell.Formula10 == 0 {
			spell.Formula10 = 100
		}
		if spell.Formula11 == 0 {
			spell.Formula11 = 100
		}
		if spell.Formula12 == 0 {
			spell.Formula12 = 100
		}

		if spell.Effectid1 == 0 {
			spell.Effectid1 = 254
		}
		if spell.Effectid2 == 0 {
			spell.Effectid2 = 254
		}
		if spell.Effectid3 == 0 {
			spell.Effectid3 = 254
		}
		if spell.Effectid4 == 0 {
			spell.Effectid4 = 254
		}
		if spell.Effectid5 == 0 {
			spell.Effectid5 = 254
		}
		if spell.Effectid6 == 0 {
			spell.Effectid6 = 254
		}
		if spell.Effectid7 == 0 {
			spell.Effectid7 = 254
		}
		if spell.Effectid8 == 0 {
			spell.Effectid8 = 254
		}
		if spell.Effectid9 == 0 {
			spell.Effectid9 = 254
		}
		if spell.Effectid10 == 0 {
			spell.Effectid10 = 254
		}
		if spell.Effectid11 == 0 {
			spell.Effectid11 = 254
		}
		if spell.Effectid12 == 0 {
			spell.Effectid12 = 254
		}
		if spell.Targettype == 0 {
			spell.Targettype = 2
		}
		if spell.Zonetype == 0 {
			spell.Zonetype = -1
		}

		if spell.Classes1 == 0 {
			spell.Classes1 = 255
		}
		if spell.Classes2 == 0 {
			spell.Classes2 = 255
		}
		if spell.Classes3 == 0 {
			spell.Classes3 = 255
		}
		if spell.Classes4 == 0 {
			spell.Classes4 = 255
		}
		if spell.Classes5 == 0 {
			spell.Classes5 = 255
		}
		if spell.Classes6 == 0 {
			spell.Classes6 = 255
		}
		if spell.Classes7 == 0 {
			spell.Classes7 = 255
		}
		if spell.Classes8 == 0 {
			spell.Classes8 = 255
		}
		if spell.Classes9 == 0 {
			spell.Classes9 = 255
		}
		if spell.Classes10 == 0 {
			spell.Classes10 = 255
		}
		if spell.Classes11 == 0 {
			spell.Classes11 = 255
		}
		if spell.Classes12 == 0 {
			spell.Classes12 = 255
		}
		if spell.Classes13 == 0 {
			spell.Classes13 = 255
		}
		if spell.Classes14 == 0 {
			spell.Classes14 = 255
		}
		if spell.Classes15 == 0 {
			spell.Classes15 = 255
		}
		if spell.Classes16 == 0 {
			spell.Classes16 = 255
		}
		if spell.CastingAnim == 0 {
			spell.CastingAnim = 44
		}

		//if `SpellAffectIndex` int(11) NOT NULL DEFAULT -1,

		if spell.Field142 == 0 {
			spell.Field142 = 100
		}
		if spell.New_icon == 0 {
			spell.New_icon = 161
		}

		if spell.Descnum == 0 {
			return fmt.Errorf("descnum is 0 for spell %d", spell.ID)
		}

		if spell.Typedescnum == 0 {
			return fmt.Errorf("typedescnum is 0 for spell %d", spell.ID)
		}

		if spell.Effectdescnum == 0 {
			return fmt.Errorf("effectdescnum is 0 for spell %d", spell.ID)
		}

		if spell.Spell_category == 0 {
			spell.Spell_category = -99
		}

		if spell.Field206 == 0 {
			spell.Field206 = -1
		}

		if spell.Field210 == 0 {
			spell.Field210 = 1
		}
	}
	return nil
}
