package interfaces

import "github.com/go-pg/pg/v10"

type PgxQuerySeter interface {
	Column(columns ...string) PgxQuerySeter
	ColumnExpr(expr string, params ...interface{}) PgxQuerySeter
	Relation(name string) PgxQuerySeter
	Limit(count int) PgxQuerySeter
	Offset(offset int) PgxQuerySeter
	Where(query string, params ...interface{}) PgxQuerySeter
	Order(orders ...string) PgxQuerySeter
	Group(columns ...string) PgxQuerySeter
	Join(join string, params ...interface{}) PgxQuerySeter
	Select(dest ...interface{}) error
	SelectAndCount(dest ...interface{}) (int, error)
	Insert(dest ...interface{}) (pg.Result, error)
	Update(dest ...interface{}) (pg.Result, error)
	Delete(dest ...interface{}) (pg.Result, error)
}

type PgxDB interface {
	Model(model ...interface{}) PgxQuerySeter
}
