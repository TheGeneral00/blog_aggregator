package functions

import (
	"context"
	"errors"
	"fmt"

	"github.com/TheGeneral00/blog_aggregator/internal/config"
	"github.com/TheGeneral00/blog_aggregator/internal/database"
)

type state struct {
        config *config.Config
        db *database.Queries
}

type command struct{
        Name string
        Args []string
}

type commands struct{
        registeredCommands map[string]func(*state, command) error
}

func NewState(config *config.Config, dbQuerries *database.Queries) (*state, error){
        if config == nil {
                return &state{}, fmt.Errorf("unable to create state")
        }
        return &state{config: config, db: dbQuerries}, nil
}

func NewCommands() *commands {
        return &commands{
                registeredCommands: map[string]func(*state, command) error{
                "login": handlerLogin,
                "register": handlerRegister,
                "reset": handlerReset,
                "users": handlerListUsers,
                "agg": middlewareLoggedIn(handlerAggregate),
                "addfeed": middlewareLoggedIn(handlerAddFeed),
                "feeds": handlerFeeds,
                "follow": middlewareLoggedIn(handlerFollow),
                "following": middlewareLoggedIn(handlerFollowing),
                "unfollow": middlewareLoggedIn(handlerUnfollow),
                "browse": middlewareLoggedIn(handlerBrowse),
                },
        }
}

func NewCommand(name string, args []string) command{
        return command{Name: name, Args: args}
}

func (c *commands) Run(s *state, cmd command) error {
        f, ok := c.registeredCommands[cmd.Name]
        if !ok {
                return errors.New("command not found")
        }
        return f(s, cmd)
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
        return func(s *state, cmd command) error {

                user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
                if err != nil {
                        return err
                }

                return handler(s, cmd, user)
        }
}
