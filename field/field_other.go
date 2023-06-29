package field

import "fmt"

func (e expr) ColumnNameWithTable() string {
	return fmt.Sprintf("`%s`.`%s`", e.col.Table, e.col.Name)
}
