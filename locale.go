package abstraction

import (
	"golang.org/x/text/currency"
	"time"
)

type Locale interface {
	InitLocale()
	// Get translation item by key
	Get(key string) string
	// Plural translation item by key contains multiple variables, the param map key should be
	// the name of the translation item variables.
	// example:
	//
	//	"messages_count": {
	//	   "msg": "You have {{.Count}} new message.",
	//	}"
	//
	// map[string]string{"Count": "12"}
	Plural(key string, params map[string]string) string
	FormatNumber(number int64) string
	FormatDate(date time.Time) string
	FormatCurrency(value float64, cur currency.Unit) string
}
