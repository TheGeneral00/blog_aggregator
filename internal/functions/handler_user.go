package functions

import(
        "context"
        "fmt"
        "time"
        "github.com/TheGeneral00/blog_aggregator/internal/database"
        "github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
        users, err := s.db.GetUsers(context.Background())
        if err != nil{
                return fmt.Errorf("couldn't list users: %w", err)
        }
        for _, user := range users {
                if user.Name == s.config.CurrentUserName {
                        fmt.Printf("* %v (current)\n", user.Name)
                        continue 
                }
                fmt.Printf("* %v\n", user.Name)
        }
        return nil
}

func printUser(user database.User) {
        fmt.Printf(" * ID: %v\n", user.ID)
        fmt.Printf(" * Name: %v\n", user.Name)
}
