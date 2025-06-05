// TODO: Implement this file
// internal/parser/parser.go
package parser

import (
    "fmt"
    "strings"

    "github.com/suman181/yourdb/internal/ast"
    "github.com/suman181/yourdb/internal/tokenizer"
)

// Parse builds an AST from an input string.
func Parse(input string) (ast.Statement, error) {
    tokens := tokenizer.Tokenize(strings.TrimSpace(input))
    if len(tokens) == 0 {
        return nil, fmt.Errorf("empty command")
    }
    cmd := strings.ToUpper(tokens[0])
    switch cmd {
    case "INSERT":
        if len(tokens) != 3 {
            return nil, fmt.Errorf("usage: INSERT key value")
        }
        return &ast.InsertStmt{Key: tokens[1], Value: tokens[2]}, nil
    case "SELECT":
        if len(tokens) != 2 {
            return nil, fmt.Errorf("usage: SELECT key")
        }
        return &ast.SelectStmt{Key: tokens[1]}, nil
    case "DELETE":
        if len(tokens) != 2 {
            return nil, fmt.Errorf("usage: DELETE key")
        }
        return &ast.DeleteStmt{Key: tokens[1]}, nil
    default:
        return nil, fmt.Errorf("unknown command: %s", tokens[0])
    }
}
