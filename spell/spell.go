package spell

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	start := time.Now()
	data, err := os.ReadFile("data/spell.yaml")
	if err != nil {
		return err
	}
	spell := SpellYaml{}
	err = yaml.Unmarshal(data, &spell)
	if err != nil {
		return fmt.Errorf("spell unmarshal: %w", err)
	}

	err = spell.sanitize()
	if err != nil {
		return fmt.Errorf("spell sanitize: %w", err)
	}

	err = generateSpellSQL(&spell)
	if err != nil {
		return fmt.Errorf("generateSpellSQL: %w", err)
	}

	err = modifySpellsUS(&spell)
	if err != nil {
		return fmt.Errorf("modifySpellsUS: %w", err)
	}

	fmt.Println("finished in", time.Since(start).String())
	return nil
}

func generateSpellSQL(sp *SpellYaml) error {
	w, err := os.Create("bin/spell_out.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("REPLACE INTO `spells_new` (`id`, `name`, `player_1`, `teleport_zone`, `you_cast`, `other_casts`, `cast_on_you`, `cast_on_other`, `spell_fades`, `range`, `aoerange`, `pushback`, `pushup`, `cast_time`, `recovery_time`, `recast_time`, `buffdurationformula`, `buffduration`, `AEDuration`, `mana`, `effect_base_value1`, `effect_base_value2`, `effect_base_value3`, `effect_base_value4`, `effect_base_value5`, `effect_base_value6`, `effect_base_value7`, `effect_base_value8`, `effect_base_value9`, `effect_base_value10`, `effect_base_value11`, `effect_base_value12`, `effect_limit_value1`, `effect_limit_value2`, `effect_limit_value3`, `effect_limit_value4`, `effect_limit_value5`, `effect_limit_value6`, `effect_limit_value7`, `effect_limit_value8`, `effect_limit_value9`, `effect_limit_value10`, `effect_limit_value11`, `effect_limit_value12`, `max1`, `max2`, `max3`, `max4`, `max5`, `max6`, `max7`, `max8`, `max9`, `max10`, `max11`, `max12`, `icon`, `memicon`, `components1`, `components2`, `components3`, `components4`, `component_counts1`, `component_counts2`, `component_counts3`, `component_counts4`, `NoexpendReagent1`, `NoexpendReagent2`, `NoexpendReagent3`, `NoexpendReagent4`, `formula1`, `formula2`, `formula3`, `formula4`, `formula5`, `formula6`, `formula7`, `formula8`, `formula9`, `formula10`, `formula11`, `formula12`, `LightType`, `goodEffect`, `Activated`, `resisttype`, `effectid1`, `effectid2`, `effectid3`, `effectid4`, `effectid5`, `effectid6`, `effectid7`, `effectid8`, `effectid9`, `effectid10`, `effectid11`, `effectid12`, `targettype`, `basediff`, `skill`, `zonetype`, `EnvironmentType`, `TimeOfDay`, `classes1`, `classes2`, `classes3`, `classes4`, `classes5`, `classes6`, `classes7`, `classes8`, `classes9`, `classes10`, `classes11`, `classes12`, `classes13`, `classes14`, `classes15`, `classes16`, `CastingAnim`, `TargetAnim`, `TravelType`, `SpellAffectIndex`, `disallow_sit`, `deities0`, `deities1`, `deities2`, `deities3`, `deities4`, `deities5`, `deities6`, `deities7`, `deities8`, `deities9`, `deities10`, `deities11`, `deities12`, `deities13`, `deities14`, `deities15`, `deities16`, `field142`, `field143`, `new_icon`, `spellanim`, `uninterruptable`, `ResistDiff`, `dot_stacking_exempt`, `deleteable`, `RecourseLink`, `no_partial_resist`, `field152`, `field153`, `short_buff_box`, `descnum`, `typedescnum`, `effectdescnum`, `effectdescnum2`, `npc_no_los`, `field160`, `reflectable`, `bonushate`, `field163`, `field164`, `ldon_trap`, `EndurCost`, `EndurTimerIndex`, `IsDiscipline`, `field169`, `field170`, `field171`, `field172`, `HateAdded`, `EndurUpkeep`, `numhitstype`, `numhits`, `pvpresistbase`, `pvpresistcalc`, `pvpresistcap`, `spell_category`, `pvp_duration`, `pvp_duration_cap`, `pcnpc_only_flag`, `cast_not_standing`, `can_mgb`, `nodispell`, `npc_category`, `npc_usefulness`, `MinResist`, `MaxResist`, `viral_targets`, `viral_timer`, `nimbuseffect`, `ConeStartAngle`, `ConeStopAngle`, `sneaking`, `not_extendable`, `field198`, `field199`, `suspendable`, `viral_range`, `songcap`, `field203`, `field204`, `no_block`, `field206`, `spellgroup`, `rank`, `field209`, `field210`, `CastRestriction`, `allowrest`, `InCombat`, `OutofCombat`, `field215`, `field216`, `field217`, `aemaxtargets`, `maxtargets`, `field220`, `field221`, `field222`, `field223`, `persistdeath`, `field225`, `field226`, `min_dist`, `min_dist_mod`, `max_dist`, `max_dist_mod`, `min_range`, `field232`, `field233`, `field234`, `field235`, `field236`) VALUES\n")

	//(3,'Summon Corpse','PLAYER_1','','','','','','',10000,0,0,0,5000,1500,12000,0,0,0,700,70,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2512,2106,17355,-1,-1,-1,1,1,1,1,-1,-1,-1,-1,100,100,100,100,100,100,100,100,100,100,100,100,0,-1,0,0,91,254,254,254,254,254,254,254,254,254,254,254,6,20,14,-1,0,0,255,255,255,255,35,255,255,255,255,255,35,255,255,255,255,255,43,0,0,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,100,0,0,98,0,0,0,0,0,0,0,0,0,3,125,64,0,-1,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,5,100,49,52,0,0,0,-1,-1,0,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,0,0,0,1,0,0,0,1,1,0,-1,0,0,0,1,39,-1,0,1,0,0,1,0,1,0,0,0,0,0,0),
	for spellIndex, spell := range sp.Spells {
		fields := structs.Fields(spell)
		w.WriteString("(")
		for fieldIndex, field := range fields {
			if !field.IsExported() {
				continue
			}
			switch field.Kind() {
			case reflect.String:
				w.WriteString(fmt.Sprintf("'%s'", field.Value()))
			case reflect.Int:
				w.WriteString(fmt.Sprintf("%d", field.Value()))
			case reflect.Float64:
				w.WriteString(fmt.Sprintf("%f", field.Value()))
			case reflect.Bool:
				w.WriteString(fmt.Sprintf("%t", field.Value()))
			default:
				return fmt.Errorf("unknown type %s", field.Kind())
			}
			if fieldIndex == len(fields)-1 {
				w.WriteString("")
			} else {
				w.WriteString(", ")
			}
		}
		if spellIndex == len(sp.Spells)-1 {
			w.WriteString(");\n")
		} else {
			w.WriteString("),\n")
		}
	}

	return nil
}
