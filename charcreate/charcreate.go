package charcreate

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xackery/rebuildeq/util"
)

func Build(cmd *cobra.Command, args []string) error {
	start := time.Now()
	fmt.Printf("CharCreate...")
	var err error
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	}()
	err = generate()
	return nil
}

func generate() error {
	data, err := os.ReadFile("charcreate.yaml")
	if err != nil {
		return fmt.Errorf("read charCreate.yaml: %w", err)
	}

	charCreate := CharCreateYaml{}
	err = yaml.Unmarshal(data, &charCreate)
	if err != nil {
		return fmt.Errorf("charCreate unmarshal: %w", err)
	}

	err = charCreate.sanitize()
	if err != nil {
		return fmt.Errorf("charCreate sanitize: %w", err)
	}

	err = generateCharCreateSQL(&charCreate)
	if err != nil {
		return fmt.Errorf("generateCharCreateSQL: %w", err)
	}

	return nil
}

func generateCharCreateSQL(sp *CharCreateYaml) error {
	w, err := os.Create("charcreate.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("REPLACE INTO `char_create_combinations` (`allocation_id`, `race`, `class`, `deity`, `start_zone`, `expansions_req`) VALUES\n")

	for charCreateIndex, charCreate := range sp.CharCreates {
		for _, class := range charCreate.Classes {
			for _, deity := range class.Deities {
				w.WriteString(fmt.Sprintf("(%d, %d, %d, %d, %d, %d",
					deity.AllocationID,
					util.RaceNameToID(charCreate.Race),
					util.ClassNameToID(class.Class),
					util.DeityNameToID(deity.Deity),
					util.ZoneNameToID(deity.StartZone),
					deity.ExpansionsReq))

				//w.WriteString(fmt.Sprintf("%d, ", util.RaceNameToID(charCreate.Race)))
				if charCreateIndex == len(sp.CharCreates)-1 {
					w.WriteString(");\n")
				} else {
					w.WriteString("),\n")
				}
			}
		}
	}

	w.WriteString("\n\n")

	w.WriteString("REPLACE INTO `char_create_point_allocations` (`id`, `base_str`, `base_sta`, `base_dex`, `base_agi`, `base_int`, `base_wis`, `base_cha`, `alloc_str`, `alloc_sta`, `alloc_dex`, `alloc_agi`, `alloc_int`, `alloc_wis`, `alloc_cha`) VALUES\n")
	for allocIndex, allocation := range sp.Allocations {
		w.WriteString(fmt.Sprintf("(%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
			allocation.ID,
			allocation.BaseStr,
			allocation.BaseSta,
			allocation.BaseDex,
			allocation.BaseAgi,
			allocation.BaseInt,
			allocation.BaseWis,
			allocation.BaseCha,
			allocation.AllocStr,
			allocation.AllocSta,
			allocation.AllocDex,
			allocation.AllocAgi,
			allocation.AllocInt,
			allocation.AllocWis,
			allocation.AllocCha))

		if allocIndex == len(sp.Allocations)-1 {
			w.WriteString(");\n")
		} else {
			w.WriteString("),\n")
		}
	}

	return nil
}

// Import takes database info and dumps to yaml
func Import(cmd *cobra.Command, args []string) error {
	if !viper.IsSet("db_host") {
		return fmt.Errorf("db_host is not set, pass it as an argument --db_host=... or set it in .rebuildeq.yaml")
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8", viper.GetString("db_user"), viper.GetString("db_pass"), viper.GetString("db_host"), viper.GetString("db_name")))
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	charCreate := &CharCreateYaml{}
	err = charCreate.sanitize()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows1, err := db.QueryxContext(ctx, "SELECT * FROM char_create_point_allocations")
	if err != nil {
		return fmt.Errorf("db query allocations: %w", err)
	}
	defer rows1.Close()

	type charCreatePointAllocationDB struct {
		ID       int `db:"id"`
		BaseStr  int `db:"base_str"`
		BaseSta  int `db:"base_sta"`
		BaseDex  int `db:"base_dex"`
		BaseAgi  int `db:"base_agi"`
		BaseInt  int `db:"base_int"`
		BaseWis  int `db:"base_wis"`
		BaseCha  int `db:"base_cha"`
		AllocStr int `db:"alloc_str"`
		AllocSta int `db:"alloc_sta"`
		AllocDex int `db:"alloc_dex"`
		AllocAgi int `db:"alloc_agi"`
		AllocInt int `db:"alloc_int"`
		AllocWis int `db:"alloc_wis"`
		AllocCha int `db:"alloc_cha"`
	}

	for rows1.Next() {
		r := charCreatePointAllocationDB{}
		err = rows1.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db allocation struct scan: %w", err)
		}

		charCreate.Allocations = append(charCreate.Allocations, &Allocation{
			ID:       r.ID,
			BaseStr:  r.BaseStr,
			BaseSta:  r.BaseSta,
			BaseDex:  r.BaseDex,
			BaseAgi:  r.BaseAgi,
			BaseInt:  r.BaseInt,
			BaseWis:  r.BaseWis,
			BaseCha:  r.BaseCha,
			AllocStr: r.AllocStr,
			AllocSta: r.AllocSta,
			AllocDex: r.AllocDex,
			AllocAgi: r.AllocAgi,
			AllocInt: r.AllocInt,
			AllocWis: r.AllocWis,
			AllocCha: r.AllocCha,
		})
	}

	rows, err := db.QueryxContext(ctx, "SELECT * FROM char_create_combinations")
	if err != nil {
		return fmt.Errorf("db query combinations: %w", err)
	}
	defer rows.Close()

	type charCreateDB struct {
		AllocationID  int `db:"allocation_id"`
		Race          int `db:"race"`
		Class         int `db:"class"`
		Deity         int `db:"deity"`
		StartZone     int `db:"start_zone"`
		ExpansionsReq int `db:"expansions_req"`
	}

	for rows.Next() {
		r := charCreateDB{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db combinations struct scan: %w", err)
		}

		isFound := false
		var raceIndex int
		var c *CharCreate
		for raceIndex, c = range charCreate.CharCreates {
			if util.RaceNameToID(c.Race) == r.Race {
				isFound = true
				break
			}
		}
		if !isFound {
			charCreate.CharCreates = append(charCreate.CharCreates, &CharCreate{Race: util.RaceIDToName(r.Race)})
			raceIndex = len(charCreate.CharCreates) - 1
		}

		var classIndex int
		isFound = false
		var cl *CharClass
		for classIndex, cl = range charCreate.CharCreates[raceIndex].Classes {
			if util.ClassNameToID(cl.Class) == r.Class {
				isFound = true
				break
			}
		}

		if !isFound {
			charCreate.CharCreates[raceIndex].Classes = append(charCreate.CharCreates[raceIndex].Classes, &CharClass{Class: util.ClassIDToName(r.Class)})
			classIndex = len(charCreate.CharCreates[raceIndex].Classes) - 1
		}

		var deityIndex int
		isFound = false
		var dl *CharDeity
		for deityIndex, dl = range charCreate.CharCreates[raceIndex].Classes[classIndex].Deities {
			if util.ClassNameToID(dl.Deity) == r.Deity {
				isFound = true
				break
			}
		}

		if !isFound {
			charCreate.CharCreates[raceIndex].Classes[classIndex].Deities = append(charCreate.CharCreates[raceIndex].Classes[classIndex].Deities, &CharDeity{Deity: util.DeityIDToName(r.Deity)})
			deityIndex = len(charCreate.CharCreates[raceIndex].Classes[classIndex].Deities) - 1
		}

		focus := charCreate.CharCreates[raceIndex].Classes[classIndex].Deities[deityIndex]
		focus.StartZone = util.ZoneIDToName(r.StartZone)
		focus.ExpansionsReq = r.ExpansionsReq
		focus.AllocationID = r.AllocationID
	}

	w, err := os.Create("charcreate_dump.yaml")
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	err = enc.Encode(charCreate)
	if err != nil {
		return err
	}

	fmt.Println("Created charcreate_dump.yaml")
	return nil
}
