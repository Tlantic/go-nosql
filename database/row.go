package database

import "time"

// A rows/records should be segmented by table/type
type Row interface {

	// row/doc key, a unique key used for direct access
	// (may be its id or other unique value)
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
	Metadata() map[string]interface{}
	GetMeta(string) interface{}
	SetMeta(string, interface{})

	// util for accessing timestamps, usually
	// available within the metadata value map
	CreatedOn() *time.Time
	UpdatedOn() *time.Time

	// after an operation (i.e bulk inserts) keep
	// errors within the result instead of failing
	// the whole operation if object represents a failed op
	IsFaulted() bool

	// fault
	Fault() error
}