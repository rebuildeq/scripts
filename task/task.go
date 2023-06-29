package task

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/go-yaml/yaml"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Build(cmd *cobra.Command, args []string) error {
	start := time.Now()
	fmt.Printf("Task...")
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
	}()
	data, err := os.ReadFile("task.yaml")
	if err != nil {
		return err
	}
	task := TaskYaml{}
	err = yaml.Unmarshal(data, &task)
	if err != nil {
		return fmt.Errorf("task unmarshal: %w", err)
	}

	err = task.sanitize()
	if err != nil {
		return fmt.Errorf("task sanitize: %w", err)
	}

	err = generateTaskSQL(&task)
	if err != nil {
		return fmt.Errorf("generateTaskSQL: %w", err)
	}

	return nil
}

func generateTaskSQL(sp *TaskYaml) error {
	w, err := os.Create("task.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("REPLACE INTO `tasks` (`taskid`, `activityid`, `req_activity_id`, `step`, `activitytype`, `target_name`, `goalmethod`, `goalcount`, `description_override`, `npc_match_list`, `item_id_list`, `item_list`, `dz_switch_id`, `min_x`, `min_y`, `min_z`, `max_x`, `max_y`, `max_z`, `skill_list`, `spell_list`, `zones`, `zone_version`, `optional`) VALUES\n")

	for taskIndex, task := range sp.Tasks {
		fields := structs.Fields(task)
		w.WriteString("(")
		for fieldIndex, field := range fields {
			if !field.IsExported() {
				continue
			}
			switch field.Kind() {
			case reflect.Slice:
				// assert type
				switch v := field.Value().(type) {
				case []*Activity:
					if len(v) == 0 {
						continue
					}
					//ignore for now
				}
			case reflect.Struct:
				// assert type
				switch v := field.Value().(type) {
				case sql.NullString:
					if v.Valid {
						w.WriteString(fmt.Sprintf("\"%s\"", v.String))
					} else {
						w.WriteString("NULL")
					}
				}
			case reflect.String:
				w.WriteString(fmt.Sprintf("\"%s\"", field.Value()))
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
		if taskIndex == len(sp.Tasks)-1 {
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

	task := &TaskYaml{}
	err = task.sanitize()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT * FROM tasks")
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Task{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}

		activityRows, err := db.QueryxContext(ctx, "SELECT * FROM task_activities WHERE taskid = ?", r.ID)
		if err != nil {
			return fmt.Errorf("db query activities for taskid %d: %w", r.ID, err)
		}

		for activityRows.Next() {
			a := Activity{}
			err = activityRows.StructScan(&a)
			if err != nil {
				return fmt.Errorf("db struct scan activity: %w", err)
			}
			a.TaskID = 0
			r.Activities = append(r.Activities, &a)
		}

		err = r.omitEmpty()
		if err != nil {
			return fmt.Errorf("omit empty: %w", err)
		}
		task.Tasks = append(task.Tasks, &r)
	}

	w, err := os.Create("task_dump.yaml")
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	err = enc.Encode(task)
	if err != nil {
		return err
	}
	fmt.Println("Created task_dump.yaml")
	return nil
}
