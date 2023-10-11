package abstraction

import (
	"database/sql"
	"gorm.io/gorm"
)

type (
	// Query this custom type is defined to be adaptable
	//to handle any other sql driver
	Query *gorm.DB

	SeederItem struct {
		Dependency interface{}   // the module entity instance, ex. : domain.User{}
		Data       []interface{} // the domain mocked data, base on the related struct
	}

	Sql interface {
		InitSql()
		Migrate(path string)
		Seed(items []SeederItem)
		SqlQuery
	}

	SqlQuery interface {
		Close()
		Where(query interface{}, args ...interface{}) Sql
		Or(query interface{}, args ...interface{}) Sql
		Not(query interface{}, args ...interface{}) Sql
		Limit(value int) Sql
		Offset(value int) Sql
		Order(value string) Sql
		Select(query interface{}, args ...interface{}) Sql
		Omit(columns ...string) Sql
		Group(query string) Sql
		Having(query string, values ...interface{}) Sql
		Joins(query string, args ...interface{}) Sql
		Scopes(funcs ...func(Query) Sql) Sql
		Unscoped() Sql
		Attrs(attrs ...interface{}) Sql
		Assign(attrs ...interface{}) Sql
		First(out interface{}, where ...interface{}) Sql
		Last(out interface{}, where ...interface{}) Sql
		Find(out interface{}, where ...interface{}) Sql
		Scan(dest interface{}) Sql
		Row() *sql.Row
		Rows() (*sql.Rows, error)
		ScanRows(rows *sql.Rows, result interface{}) error
		Pluck(column string, value interface{}) Sql
		Count(value *int64) Sql
		FirstOrInit(out interface{}, where ...interface{}) Sql
		FirstOrCreate(out interface{}, where ...interface{}) Sql
		Update(column string, attrs ...interface{}) Sql
		Updates(values interface{}) Sql
		UpdateColumn(column string, attrs ...interface{}) Sql
		UpdateColumns(values interface{}) Sql
		Save(value interface{}) Sql
		Create(value interface{}) Sql
		Delete(value interface{}, where ...interface{}) Sql
		Raw(sql string, values ...interface{}) Sql
		Exec(sql string, values ...interface{}) Sql
		Model(value interface{}) Sql
		Table(name string) Sql
		Debug() Sql
		Begin() Sql
		Commit() Sql
		Rollback() Sql
		AutoMigrate(values ...interface{}) error
		Association(column string) *gorm.Association
		Preload(column string, conditions ...interface{}) Sql
		Set(name string, value interface{}) Sql
		InstanceSet(name string, value interface{}) Sql
		Get(name string) (value interface{}, ok bool)
		SetJoinTable(model interface{}, column string, handler interface{}) error
		AddError(err error) error
		Error() error        // extra
		RowsAffected() int64 // extra
	}
)
