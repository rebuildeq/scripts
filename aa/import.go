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

	for rows.Next() {
		r := &AASkill{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}
		aa.Skills = append(aa.Skills, r)
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

	ranks := []*AARank{}

	for rows.Next() {
		r := &AARank{}
		err = rows.StructScan(r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		ranks = append(ranks, r)
	}

	for _, r := range ranks {
		for _, skill := range aa.Skills {
			if skill.FirstRankID == r.ID {
				//fmt.Println("Skill", skill.ID, "has first rank", r.ID)
				err = searchRanks(skill, ranks, r.ID)
				if err != nil {
					return fmt.Errorf("searchRanks: %w", err)
				}
				break
			}
		}
	}

	return nil
}

func searchRanks(skill *AASkill, ranks []*AARank, id int) error {
	isFirst := true
	index := 1
	for {
		if id < 1 {
			return nil
		}
		next := findRank(ranks, id)
		if next == nil {
			return fmt.Errorf("rank %d not found", id)
		}
		if isFirst {
			skill.NameSID = next.TitleSID
			isFirst = false
		}
		next.Index = index
		index++
		skill.Ranks = append(skill.Ranks, next)
		id = next.nextID
	}
}

func findRank(ranks []*AARank, id int) *AARank {
	for _, rankEntry := range ranks {
		if rankEntry.ID == id {
			return rankEntry
		}
	}
	return nil
}
