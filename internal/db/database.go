// TODO: Implement this file
// internal/db/db.go
package db

import (
    "fmt"

    "github.com/example/db/internal/btree"
    "github.com/example/db/internal/storage"
    "github.com/example/db/internal/parser"
    "github.com/example/db/internal/bytecode"
    "github.com/example/db/internal/vm"
)

type DB struct {
    pager *storage.Pager
    tree  *btree.BTree
}

// NewDB opens/creates the database file.
func NewDB(path string) (*DB, error) {
    pager, err := storage.NewPager(path)
    if err != nil {
        return nil, err
    }
    return &DB{pager: pager, tree: btree.New()}, nil
}

// Insert adds or updates a key-value pair.
func (db *DB) Insert(key, value string) error {
    page := db.pager.AllocatePage()
    if err := db.pager.WritePage(page, []byte(value)); err != nil {
        return err
    }
    db.tree.Insert(key, page)
    return db.pager.Sync()
}

// Select retrieves the value for a key.
func (db *DB) Select(key string) (string, error) {
    page, ok := db.tree.Search(key)
    if !ok {
        return "", fmt.Errorf("key not found: %s", key)
    }
    data, err := db.pager.ReadPage(page)
    if err != nil {
        return "", err
    }
    // Convert page bytes to string (trimming any trailing zeros).
    return string(data), nil
}

// Delete removes a key and frees its page.
func (db *DB) Delete(key string) error {
    page, ok := db.tree.Search(key)
    if !ok {
        return fmt.Errorf("key not found: %s", key)
    }
    db.tree.Delete(key)
    db.pager.FreePage(page)
    return nil
}

// Close closes the database file.
func (db *DB) Close() error {
    return db.pager.Close()
}

// Exec parses and executes a command string, returning any result.
func (db *DB) Exec(input string) (string, error) {
    stmt, err := parser.Parse(input)
    if err != nil {
        return "", err
    }
    instructions := bytecode.Compile(stmt)
    return vm.NewVM(db).Run(instructions)
}
