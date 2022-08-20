package types

type Set interface {
	Add(interface{})
	Contains(interface{}) bool
}
