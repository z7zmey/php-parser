package node

type Attributer interface {
	Attributes() map[string]interface{}
	Attribute(key string) interface{}
	SetAttribute(key string, value interface{})
}

type Node interface {
	Attributer
	Name() string
	Walk(v Visitor)
}
