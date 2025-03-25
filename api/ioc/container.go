package ioc

type IContainer interface {
	RegisterDependencies() error
}
