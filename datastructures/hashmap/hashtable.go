package hashmap

// HashTable represents a hash table interface
type HashTable interface {
	// Put will insert value with key
	Put(key interface{}, value interface{})

	// Get : returns value of a given key
	// return nil if key doesn't exists
	Get(key interface{}) interface{}

	// Delete : remove the key from the table
	Delete(key interface{})

	// Contains : return true if given exists in table and false otherwise
	Contains(key interface{}) bool

	// Len : returns length of the hash table
	Len() int

	// Cap : returns current capacity of the hash table
	Cap() int
}

// NewHashTable : creates new HashTable
func NewHashTable() HashTable {
	return newTable()
}
