package database

import (
	"time"
)

//noinspection ALL
const (

	GUID		= "_uId"
	TYPE		= "_type"

	// Meta Keys
	CREATEDON 	= "created_on"
	UPDATEDON 	= "updated_on"
	TTL 		= "ttl"
	EXPIRY = TTL
)

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

// A rows/records should be segmented by table/type
type Row interface {

	// row/doc key, a unique key used to direct
	// accessing the row/doc (may be its id or other unique value)
	GetKey()string
	SetKey(string)

	// row/doc unique id
	GetId()string
	SetId(string)

	// table/type
	GetType()string
	SetType(string)

	// actual columns/properties
	GetData() interface{}
	SetData(interface{})

	// metadata about the row/doc
	GetMeta(string) interface{}
	SetMeta(string, interface{})

	// util for accessing timestamps, usually the
	// available within the metadata value map
	CreatedOn() *time.Time
	UpdatedOn() *time.Time

	// after an operation (i.e bulk inserts) keep
	// errors within the result instead of failing
	// the whole operation
	// if object represents a failed op
	IsFaulted() bool

	// fault
	Fault() error
}


type Database interface {
	GetName() string
	NewRow(string) Row
	NewQuery(string) Query


	Create(...interface{}) ([]Row, bool)
	CreateOne(interface{}) Row

	Read(...interface{}) ([]Row, bool)
	ReadOne(interface{}) Row
	ReadOneWithType(interface{}, interface{}) Row

	Replace(...interface{}) ([]Row, bool)
	ReplaceOne(interface{}) Row

	Update(...interface{}) ([]Row, bool)
	UpdateOne(interface{}) Row

	Destroy(...interface{}) ([]Row, bool)
	DestroyOne(interface{}) Row

	Exec(Query) (QueryResult, error)

	Close()
}