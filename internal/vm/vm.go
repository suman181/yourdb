// TODO: Implement this file
// internal/vm/vm.go
package vm

import (
    "errors"
    "fmt"

    "github.com/suman181/yourdb/internal/bytecode"
    "github.com/suman181/yourdb/internal/db"
)

// VM executes bytecode instructions using a DB.
type VM struct {
    db *db.DB
}

// NewVM creates a VM bound to a database.
func NewVM(database *db.DB) *VM {
    return &VM{db: database}
}

// Run executes instructions and returns the SELECT result (if any).
func (v *VM) Run(ins []bytecode.Instruction) (string, error) {
    var output string
    for _, inst := range ins {
        switch inst.Op {
        case bytecode.OpInsert:
            if err := v.db.Insert(inst.Key, inst.Value); err != nil {
                return "", err
            }
        case bytecode.OpSelect:
            val, err := v.db.Select(inst.Key)
            if err != nil {
                return "", err
            }
            output = val
        case bytecode.OpDelete:
            if err := v.db.Delete(inst.Key); err != nil {
                return "", err
            }
        case bytecode.OpHalt:
            return output, nil
        default:
            return "", errors.New(fmt.Sprintf("unknown opcode: %v", inst.Op))
        }
    }
    return output, nil
}
