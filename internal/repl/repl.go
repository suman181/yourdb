// TODO: Implement this file
package repl

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/suman181/yourdb/internal/db"
)

// Start starts the REPL loop.
func Start() {
    // Open (or create) the database file.
    database, err := db.NewDB("testdata/example.db")
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer database.Close()

    reader := bufio.NewScanner(os.Stdin)
    fmt.Print("db> ")
    for reader.Scan() {
        line := strings.TrimSpace(reader.Text())
        if strings.EqualFold(line, "exit") {
            fmt.Println("Exiting.")
            break
        }
        // Execute the command via the database API.
        result, err := database.Exec(line)
        if err != nil {
            fmt.Println("Error:", err)
        } else if result != "" {
            fmt.Println(result)
        }
        fmt.Print("db> ")
    }
}
