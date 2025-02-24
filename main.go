package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/TheGeneral00/blog_aggregator/internal/config"
	"github.com/TheGeneral00/blog_aggregator/internal/database"
	"github.com/TheGeneral00/blog_aggregator/internal/functions"
	_ "github.com/lib/pq"
)

func main() {
        cfg, err := config.Read()
        if err != nil {
                log.Fatal(err)
        }
        if err != nil {
                log.Fatal(err)
        }
        commands := functions.NewCommands()
        db, err := sql.Open("postgres", config.DBURL)
        if err != nil {
                log.Fatal("Failed to open dbURL")
        }
        defer db.Close()
        dbQuerries := database.New(db)

        programState, err := functions.NewState(&cfg, dbQuerries)
        if err != nil {
                log.Fatal(err)
        }

        args := os.Args
        if len(args)<2{
                log.Fatal("Not enough arguments provided")
        }
        command := functions.NewCommand(os.Args[1], os.Args[2:])
        err = commands.Run(programState, command)
        if err != nil {
                log.Fatal(err)
        }
}
