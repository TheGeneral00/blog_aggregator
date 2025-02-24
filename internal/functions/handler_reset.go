package functions

import(
        "context"
        "fmt"
)

func handlerReset(s *state, cmd command) error {
        err := s.db.Reset(context.Background())
        if err != nil {
                fmt.Errorf("couldn't delet users: %v\n", err)
        }
        fmt.Println("Database reset successfully!")
        return nil
}
