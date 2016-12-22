package database

type Query interface {
	GetStatement() string
	SetStatement(string)

	GetParams() interface{}
	SetParams(interface{})

	SetMeta(string, interface{})
	GetMeta(string) interface{}
}


type QueryResult interface {
	Query

	One(interface{}) error
	OneBytes() []byte

	Take(int) QueryResult
	Skip(int) QueryResult

	ForEach(func(int, []byte ))
	Map(func(int, []byte) interface{}) []interface{}

	Range() <- chan []byte

	Close() error
}