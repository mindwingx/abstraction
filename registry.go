package abstraction

type Registry interface {
	InitRegistry(configType string, configFilePath string) error
	ValueOf(key string) Registry
	Parse(interface{}) error
}
