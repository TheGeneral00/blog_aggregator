package main

import (
	"fmt"
	"os"
        "database/sql"
	"github.com/TheGeneral00/blog_aggregator/internal"
        "github.com/TheGeneral00/blog_aggregator/internal/config"
        _ "github.com/lib/pq"
)

func main() {
        state := functions.NewState(config.Read())
        commands := functions.NewCommands()
        db, err := sql.Open("postgres", state.GetDBURL())
        if err != nil {
                fmt.Println("Failed to open dbURL")
        }
        state.SetDB(db)
        args := os.Args
        if len(args)<2{
                fmt.Println("Not enough arguments provided")
                os.Exit(1)
        }
        command := functions.NewCommand(args)
        if commands.Run(state, command) != 0 {
                os.Exit(1)
        }
}
