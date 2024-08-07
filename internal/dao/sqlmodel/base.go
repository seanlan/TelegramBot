package sqlmodel

import "gorm.io/gorm/clause"

type FieldBase struct {
	Name      string // 字段名
	FieldName string // 完整字段名
}

func (field FieldBase) Eq(value interface{}) clause.Expression {
	return clause.Eq{Column: field.FieldName, Value: value}
}

func (field FieldBase) Neq(value interface{}) clause.Expression {
	return clause.Neq{Column: field.FieldName, Value: value}
}

func (field FieldBase) In(value interface{}) clause.Expression {
	return clause.Expr{SQL: field.FieldName + " IN ? ", Vars: []interface{}{value}}
}

func (field FieldBase) Gt(value interface{}) clause.Expression {
	return clause.Gt{Column: field.FieldName, Value: value}
}

func (field FieldBase) Gte(value interface{}) clause.Expression {
	return clause.Gte{Column: field.FieldName, Value: value}
}

func (field FieldBase) Lt(value interface{}) clause.Expression {
	return clause.Lt{Column: field.FieldName, Value: value}
}

func (field FieldBase) Lte(value interface{}) clause.Expression {
	return clause.Lte{Column: field.FieldName, Value: value}
}

func (field FieldBase) Like(value interface{}) clause.Expression {
	return clause.Like{Column: field.FieldName, Value: value}
}

func (field FieldBase) NotLike(value interface{}) clause.Expression {
	return clause.Expr{SQL: field.FieldName + " NOT LIKE ", Vars: []interface{}{value}}
}

func (field FieldBase) Add(value interface{}) clause.Expression {
	return clause.Expr{
		SQL:  field.FieldName + " + ?",
		Vars: []interface{}{value},
	}
}

func (field FieldBase) Mul(value interface{}) clause.Expression {
	return clause.Expr{
		SQL:  field.FieldName + " * ?",
		Vars: []interface{}{value},
	}
}

func (field FieldBase) IF(symbol string, c, value interface{}) clause.Expression {
	return clause.Expr{
		SQL:  "IF(" + field.FieldName + symbol + " ? ,?," + field.FieldName + ")",
		Vars: []interface{}{c, value},
	}
}

func (field FieldBase) IFAdd(symbol string, target, delta, value interface{}) clause.Expression {
	return clause.Expr{
		SQL:  "IF(" + field.FieldName + symbol + " ? ," + field.FieldName + "+ ?, ?)",
		Vars: []interface{}{target, delta, value},
	}
}

func (field FieldBase) Any(value interface{}) clause.Expression {
	return clause.Expr{
		SQL:  "? = ANY(" + field.FieldName + ")",
		Vars: []interface{}{value},
	}
}

func (field FieldBase) All(value interface{}) clause.Expression {
	return clause.Expr{
		SQL:  "? = ALL(" + field.FieldName + ")",
		Vars: []interface{}{value},
	}
}

func (field FieldBase) Desc() string {
	return field.FieldName + " DESC"
}

func (field FieldBase) Asc() string {
	return field.FieldName + " ASC"
}
