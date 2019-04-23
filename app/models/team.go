package models

import (
	"fmt"
	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type Team struct {
	Id 	      int64 `db:"team_id, primarykey, autoincrement"`
	Name      int64 `db:"team_name"`
	LeagueId  int64 `db:"league_id"`

	// Transient
	League    *League
}

func (t *Team) Validate(v *revel.Validation) {
	v.Check(t.Name,
		revel.ValidRequired())
}

func (t *Team) PostGet(exe gorp.SqlExecutor) error {
	var (
		obj interface{}
		err error
	)

	obj, err = exe.Get(League{}, t.LeagueId)
	if err != nil {
		return fmt.Errorf("Error loading a team's league (%d): %s", t.LeagueId, err)
	}

	t.League = obj.(*League)

	return nil
}