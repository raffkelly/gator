package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("invalid arguments for 'reset' command")
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("users and all linked tables reset")
	return nil
}
