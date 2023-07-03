package charcreate

type CharCreateYaml struct {
	CharCreates []*CharCreate `yaml:"charCreates,omitempty" db:"charCreates"`
	Allocations []*Allocation `yaml:"allocations,omitempty" db:"allocations"`
}

type CharCreate struct {
	Race    string       `yaml:"race"`
	Classes []*CharClass `yaml:"classes,omitempty" db:"classes"`
}

type CharClass struct {
	Class   string       `yaml:"class"`
	Deities []*CharDeity `yaml:"deities,omitempty" db:"deities"`
}

type CharDeity struct {
	Deity         string `yaml:"deity"`
	StartZone     string `yaml:"start_zone"`
	ExpansionsReq int    `yaml:"expansions_req"`
	AllocationID  int    `yaml:"allocation_id"`
}

type Allocation struct {
	ID       int `yaml:"id"`
	BaseStr  int `yaml:"base_str"`
	BaseSta  int `yaml:"base_sta"`
	BaseDex  int `yaml:"base_dex"`
	BaseAgi  int `yaml:"base_agi"`
	BaseInt  int `yaml:"base_int"`
	BaseWis  int `yaml:"base_wis"`
	BaseCha  int `yaml:"base_cha"`
	AllocStr int `yaml:"alloc_str"`
	AllocSta int `yaml:"alloc_sta"`
	AllocDex int `yaml:"alloc_dex"`
	AllocAgi int `yaml:"alloc_agi"`
	AllocInt int `yaml:"alloc_int"`
	AllocWis int `yaml:"alloc_wis"`
	AllocCha int `yaml:"alloc_cha"`
}

func (e *CharCreateYaml) sanitize() error {
	for _, charCreate := range e.CharCreates {
		err := charCreate.sanitize()
		if err != nil {
			return err
		}
	}
	return nil
}

func (charCreate *CharCreate) sanitize() error {

	return nil
}

func (charCreate *CharCreate) omitEmpty() error {
	return nil
}
