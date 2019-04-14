package models

import "github.com/revel/revel"

type League struct {
	Id 	  int64 `db:"league_id, primarykey, autoincrement"`
	Name  int64 `db:"league_name"`
}

func (l *League) Validate(v *revel.Validation) {
	v.Check(l.Name,
		revel.ValidRequired())
}