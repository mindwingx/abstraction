package abstraction

type Registry interface {
	InitRegistry()
	ValueOf(string) Registry
	Parse(interface{})
}
