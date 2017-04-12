package database

// RESERVED Metadata Keys.
// Using such values may dictate different behaviour based on implementation
//noinspection ALL
const (
	GUID = "_uId"
	TYPE = "_type"

	// Meta
	LOCK           = "_lock"
	CAS            = "_cas"
	ADHOC          = "_adhoc"
	CONSISTENCY    = "_consistency"
	CONSISTENTWITH = "_consistent_with"
	TIMEOUT        = "_timeout"
	TTL            = "_ttl"
	EXPIRY         = TTL

	CREATEDON = "created_on"
	UPDATEDON = "updated_on"
)

type Interface interface {
	GetName() string
	NewRow(string) Row
	NewQuery(string) Query

	Create(...interface{}) ([]Row, bool)
	CreateOne(interface{}) Row

	Read(...interface{}) ([]Row, bool)
	ReadOne(interface{}) Row
	ReadOneWithType(interface{}, interface{}) Row

	Touch(...interface{}) ([]Row, bool)
	TouchOne(interface{}) Row

	Replace(...interface{}) ([]Row, bool)
	ReplaceOne(interface{}) Row

	Update(...interface{}) ([]Row, bool)
	UpdateOne(interface{}) Row

	Destroy(...interface{}) ([]Row, bool)
	DestroyOne(interface{}) Row

	Exec(Query) (QueryResult, error)

	Close()
}
