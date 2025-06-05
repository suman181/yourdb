// TODO: Implement this file
// internal/btree/btree.go
package btree

// BTree is a simple key→page map (placeholder for a real B+Tree).
type BTree struct {
    data map[string]uint64
}

// New creates a new BTree.
func New() *BTree {
    return &BTree{data: make(map[string]uint64)}
}

// Insert adds or updates a key with a page number.
func (bt *BTree) Insert(key string, page uint64) {
    bt.data[key] = page
}

// Search looks up the key and returns the page number and if it exists.
func (bt *BTree) Search(key string) (uint64, bool) {
    page, ok := bt.data[key]
    return page, ok
}

// Delete removes a key (and its page) from the map.
func (bt *BTree) Delete(key string) {
    delete(bt.data, key)
}
