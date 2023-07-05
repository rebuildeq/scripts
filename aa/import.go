package aa

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Import(cmd *cobra.Command, args []string) error {
	if !viper.IsSet("db_host") {
		return fmt.Errorf("db_host is not set, pass it as an argument --db_host=... or set it in .rebuildeq.yaml")
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8", viper.GetString("db_user"), viper.GetString("db_pass"), viper.GetString("db_host"), viper.GetString("db_name")))
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	aa := &AAYaml{}
	err = aa.sanitize()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM aa_ability")
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	type AADB struct {
		ID              int    `db:"id"`               // int(10) unsigned NOT NULL,
		Name            string `db:"name"`             // text NOT NULL,
		Category        int    `db:"category"`         // int(10) NOT NULL DEFAULT -1,
		Classes         int    `db:"classes"`          // int(10) NOT NULL DEFAULT 131070,
		Races           int    `db:"races"`            // int(10) NOT NULL DEFAULT 65535,
		DrakkinHeritage int    `db:"drakkin_heritage"` // int(10) NOT NULL DEFAULT 127,
		Deities         int    `db:"deities"`          // int(10) NOT NULL DEFAULT 131071,
		Status          int    `db:"status"`           // int(10) NOT NULL DEFAULT 0,
		Type            int    `db:"type"`             // int(10) NOT NULL DEFAULT 0,
		Charges         int    `db:"charges"`          // int(11) NOT NULL DEFAULT 0,
		GrantOnly       int    `db:"grant_only"`       // tinyint(4) NOT NULL DEFAULT 0,
		FirstRankId     int    `db:"first_rank_id"`    // int(10) NOT NULL DEFAULT -1,
		Enabled         int    `db:"enabled"`          // tinyint(3) unsigned NOT NULL DEFAULT 1,
		ResetOnDeath    int    `db:"reset_on_death"`   // tinyint(4) NOT NULL DEFAULT 0,
	}

	for rows.Next() {
		r := AADB{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		skill := &AASkill{
			ID:              r.ID,
			Name:            r.Name,
			Type:            r.Type,
			GrantOnly:       r.GrantOnly,
			DrakkinHeritage: r.DrakkinHeritage,
			Classes:         r.Classes,
			Races:           r.Races,
		}

		aa.Skills = append(aa.Skills, skill)
	}

	err = importRanks(db, aa)
	if err != nil {
		return fmt.Errorf("importRanks: %w", err)
	}

	w, err := os.Create("aa_dump.yaml")
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	err = enc.Encode(aa)
	if err != nil {
		return err
	}
	fmt.Println("Created aa_dump.yaml")

	return nil
}

func importRanks(db *sqlx.DB, aa *AAYaml) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM aa_ranks")
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	type AARankDB struct {
		ID             int `db:"id"`               // int(10) unsigned NOT NULL,
		UpperHotkeySID int `db:"upper_hotkey_sid"` // int(10) NOT NULL DEFAULT -1,
		LowerHotkeySID int `db:"lower_hotkey_sid"` // int(10) NOT NULL DEFAULT -1,
		TitleSID       int `db:"title_sid"`        // int(10) NOT NULL DEFAULT -1,
		DescSID        int `db:"desc_sid"`         // int(10) NOT NULL DEFAULT -1,
		Cost           int `db:"cost"`             // int(10) NOT NULL DEFAULT 1,
		LevelReq       int `db:"level_req"`        // int(10) NOT NULL DEFAULT 51,
		Spell          int `db:"spell"`            // int(10) NOT NULL DEFAULT -1,
		SpellType      int `db:"spell_type"`       // int(10) NOT NULL DEFAULT 0,
		RecastTime     int `db:"recast_time"`      // int(10) NOT NULL DEFAULT 0,
		Expansion      int `db:"expansion"`        // int(10) NOT NULL DEFAULT 0,
		PrevID         int `db:"prev_id"`          // int(10) NOT NULL DEFAULT -1,
		NextID         int `db:"next_id"`          // int(10) NOT NULL DEFAULT -1,
	}

	for rows.Next() {
		r := AARankDB{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		rank := &AARank{
			ID:             r.ID,
			UpperHotkeySID: r.UpperHotkeySID,
			LowerHotkeySID: r.LowerHotkeySID,
			TitleSID:       r.TitleSID,
			DescriptionSID: r.DescSID,
			Cost:           r.Cost,
			LevelReq:       r.LevelReq,
			SpellID:        r.Spell,
			SpellType:      r.SpellType,
			RecastTime:     r.RecastTime,
			Expansion:      r.Expansion,
			PrevID:         r.PrevID,
			NextID:         r.NextID,
		}

		isFound := false
		for _, skill := range aa.Skills {
			if skill.FirstRankID == rank.ID {
				skill.Ranks = append(skill.Ranks, rank)
				isFound = true
				break
			}
			for _, rk := range skill.Ranks {
				if rk.nextID == rank.ID {
					skill.Ranks = append(skill.Ranks, rank)
					isFound = true
					break
				}
			}
			if isFound {
				break
			}
		}
		if !isFound {
			return fmt.Errorf("rank %d not found", rank.ID)
		}
	}

	return nil
}
