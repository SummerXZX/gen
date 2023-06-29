package field

import "fmt"

func (field Field) RawExprString() string {
	return fmt.Sprintf("`%s`.`%s`", field.col.Table, field.col.Name)
}
