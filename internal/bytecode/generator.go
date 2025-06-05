// TODO: Implement this file
// internal/bytecode/bytecode.go
package bytecode

import "github.com/suman181/yourdb/internal/ast"

type OpCode int

const (
    OpInsert OpCode = iota
    OpSelect
    OpDelete
    OpHalt
)

// Instruction holds an opcode and operands.
type Instruction struct {
    Op    OpCode
    Key   string
    Value string
}

// Compile converts an AST statement to instructions.
func Compile(stmt ast.Statement) []Instruction {
    switch s := stmt.(type) {
    case *ast.InsertStmt:
        return []Instruction{{Op: OpInsert, Key: s.Key, Value: s.Value}}
    case *ast.SelectStmt:
        return []Instruction{{Op: OpSelect, Key: s.Key}}
    case *ast.DeleteStmt:
        return []Instruction{{Op: OpDelete, Key: s.Key}}
    default:
        return []Instruction{}
    }
}
