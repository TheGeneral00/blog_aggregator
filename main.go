package main

import(
    "fmt"
    "github.com/TheGeneral00/blog_aggregator/internal/config"
)

func main() {
    cfg, err := config.Read()
    if err != nil {
        fmt.Println(err)
    }
    cfg.SetUser("TheGeneral00")
    cfg, err = config.Read()
    fmt.Printf("%+v\n", cfg)
}
