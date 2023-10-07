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
		Where(query interface{}, args ...interface{}) SqlQuery
		Or(query interface{}, args ...interface{}) SqlQuery
		Not(query interface{}, args ...interface{}) SqlQuery
		Limit(value int) SqlQuery
		Offset(value int) SqlQuery
		Order(value string) SqlQuery
		Select(query interface{}, args ...interface{}) SqlQuery
		Omit(columns ...string) SqlQuery
		Group(query string) SqlQuery
		Having(query string, values ...interface{}) SqlQuery
		Joins(query string, args ...interface{}) SqlQuery
		Scopes(funcs ...func(SqlQuery) SqlQuery) SqlQuery
		Unscoped() SqlQuery
		Attrs(attrs ...interface{}) SqlQuery
		Assign(attrs ...interface{}) SqlQuery
		First(out interface{}, where ...interface{}) SqlQuery
		Last(out interface{}, where ...interface{}) SqlQuery
		Find(out interface{}, where ...interface{}) SqlQuery
		Scan(dest interface{}) SqlQuery
		Row() *sql.Row
		Rows() (*sql.Rows, error)
		ScanRows(rows *sql.Rows, result interface{}) error
		Pluck(column string, value interface{}) SqlQuery
		Count(value *int64) SqlQuery
		FirstOrInit(out interface{}, where ...interface{}) SqlQuery
		FirstOrCreate(out interface{}, where ...interface{}) SqlQuery
		Update(column string, attrs ...interface{}) SqlQuery
		Updates(values interface{}) SqlQuery
		UpdateColumn(column string, attrs ...interface{}) SqlQuery
		UpdateColumns(values interface{}) SqlQuery
		Save(value interface{}) SqlQuery
		Create(value interface{}) SqlQuery
		Delete(value interface{}, where ...interface{}) SqlQuery
		Raw(sql string, values ...interface{}) SqlQuery
		Exec(sql string, values ...interface{}) SqlQuery
		Model(value interface{}) SqlQuery
		Table(name string) SqlQuery
		Debug() SqlQuery
		Begin() SqlQuery
		Commit() SqlQuery
		Rollback() SqlQuery
		AutoMigrate(values ...interface{}) error
		Association(column string) *gorm.Association
		Preload(column string, conditions ...interface{}) SqlQuery
		Set(name string, value interface{}) SqlQuery
		InstanceSet(name string, value interface{}) SqlQuery
		Get(name string) (value interface{}, ok bool)
		SetJoinTable(model interface{}, column string, handler interface{}) error
		AddError(err error) error
		Error() error        // extra
		RowsAffected() int64 // extra
	}
)
