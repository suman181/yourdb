// TODO: Implement this file
// internal/ast/ast.go
package ast

// Statement is a generic SQL-like statement.
type Statement interface {
    statementNode()
}

// InsertStmt represents: INSERT key value
type InsertStmt struct {
    Key   string
    Value string
}
func (*InsertStmt) statementNode() {}

// SelectStmt represents: SELECT key
type SelectStmt struct {
    Key string
}
func (*SelectStmt) statementNode() {}

// DeleteStmt represents: DELETE key
type DeleteStmt struct {
    Key string
}
func (*DeleteStmt) statementNode() {}
