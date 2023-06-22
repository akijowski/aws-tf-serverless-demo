package types

// KeyValueEntry represents an internal model of an entry in our "key-value" store
type KeyValueEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
