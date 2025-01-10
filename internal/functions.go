package functions

import (
        "fmt"
        "time"
        "context"
	"database/sql"
	"github.com/TheGeneral00/blog_aggregator/internal/config"
	"github.com/TheGeneral00/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type state struct{
        config *config.Config
        db *database.Queries
}

type command struct{
        name string
        args []string
}

type commands struct{
        available map[string]func(s *state, cmd command) int
}

func NewState(cfg *config.Config) *state {
        return &state{config: cfg}
}


func NewCommands() *commands {
        return &commands{
                available: map[string]func(*state, command) int{
                "login": handlerLogin,
                "register": handlerRegister,
                },
        }
}

func NewCommand(args []string) command {
        var command command
        command.name = args[1]
        command.args = args[2:]
        return command
}
   

func handlerLogin(s *state, cmd command) int{
        if len(cmd.args)==0{
                fmt.Printf("Missing username for login")
                return 1
        } 
        err := s.config.SetUser(cmd.args[0])
        if err != nil{
                fmt.Printf("Failed to login user: '%v'", err)
                return 1
        }  
        fmt.Printf("User set to %v\n", cmd.args[0])
        return 0
}

func handlerRegister(s *state, cmd command) int{
        if len(cmd.args)==0{
                fmt.Printf("Missing username to register")
                return 1
        }
        tmpContext := context.Background()
        _, err := s.db.GetUser(tmpContext, cmd.args[0])
        if err == sql.ErrNoRows { 
                now := time.Now()
                params := database.CreateUserParams{ 
                        ID: uuid.New(),
                        CreatedAt: now,
                        UpdatedAt: now,
                        Name: cmd.args[0],
                }
                user, err := s.db.CreateUser(tmpContext, params)
                if err != nil {
                        fmt.Printf("Failed to create user: '%v'", err)
                        return 1
                }
                s.config.SetUser(cmd.args[0])
                fmt.Printf("Created user: '%+v'\n", user)
                return 0
        } else if err != nil {
                fmt.Printf("Request failed with error: %v\n", err)
                return 1
        } else {
                fmt.Printf("User '%s' already exists\n", cmd.args[0])
                return 1
        }
}

func (c *commands) Run(s *state, command command) int{
        val, ok := c.available[command.name]
        if !ok {
                fmt.Printf("Command %v is not available", command.name)
                return 1
        }
        return val(s, command)
}

func (s *state) GetDBURL() string { return s.config.GetDBURL()}

func (s *state) GetConfig() config.Config {
        return *s.config
}

func (s *state) SetDB(db *sql.DB) {
        s.db = database.New(db)
}

func (s *state) CreateUser(ctx context.Context, params database.CreateUserParams) ( database.User, error ){
        return s.db.CreateUser(ctx, params)
}

func (s *state) GetUser(ctx context.Context, name string) ( database.User, error ){
        return s.db.GetUser(ctx, name)
}
