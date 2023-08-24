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
		Where(query interface{}, args ...interface{}) Query
		Or(query interface{}, args ...interface{}) Query
		Not(query interface{}, args ...interface{}) Query
		Limit(value int) Query
		Offset(value int) Query
		Order(value string) Query
		Select(query interface{}, args ...interface{}) Query
		Omit(columns ...string) Query
		Group(query string) Query
		Having(query string, values ...interface{}) Query
		Joins(query string, args ...interface{}) Query
		Scopes(funcs ...func(Query) Query) Query
		Unscoped() Query
		Attrs(attrs ...interface{}) Query
		Assign(attrs ...interface{}) Query
		First(out interface{}, where ...interface{}) Query
		Last(out interface{}, where ...interface{}) Query
		Find(out interface{}, where ...interface{}) Query
		Scan(dest interface{}) Query
		Row() *sql.Row
		Rows() (*sql.Rows, error)
		ScanRows(rows *sql.Rows, result interface{}) error
		Pluck(column string, value interface{}) Query
		Count(value *int64) Query
		FirstOrInit(out interface{}, where ...interface{}) Query
		FirstOrCreate(out interface{}, where ...interface{}) Query
		Update(column string, attrs ...interface{}) Query
		Updates(values interface{}) Query
		UpdateColumn(column string, attrs ...interface{}) Query
		UpdateColumns(values interface{}) Query
		Save(value interface{}) Query
		Create(value interface{}) Query
		Delete(value interface{}, where ...interface{}) Query
		Raw(sql string, values ...interface{}) Query
		Exec(sql string, values ...interface{}) Query
		Model(value interface{}) Query
		Table(name string) Query
		Debug() Query
		Begin() Query
		Commit() Query
		Rollback() Query
		AutoMigrate(values ...interface{}) error
		Association(column string) *gorm.Association
		Preload(column string, conditions ...interface{}) Query
		Set(name string, value interface{}) Query
		InstanceSet(name string, value interface{}) Query
		Get(name string) (value interface{}, ok bool)
		SetJoinTable(model interface{}, column string, handler interface{}) error
		AddError(err error) error
		Error() error        // extra
		RowsAffected() int64 // extra
	}
)
