package field

import "gorm.io/gorm/clause"

func NewRaw(ex clause.Expr) Field {
	return Field{expr: expr{
		e: ex,
	}}
}
