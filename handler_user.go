package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/raffkelly/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	_, err := s.db.GetUser(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to: %v\n", s.cfg.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	_, err := s.db.GetUser(context.Background(), cmd.Arguments[0])

	if err == sql.ErrNoRows {
		// User does not exist, proceed to create
		args := database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.Arguments[0],
		}
		createdUser, innerErr := s.db.CreateUser(context.Background(), args)
		if innerErr != nil {
			return innerErr
		}
		err = s.cfg.SetUser(cmd.Arguments[0])
		if err != nil {
			return err
		}
		log.Printf("User %v successfully created\n", cmd.Arguments[0])
		log.Printf("%+v\n", createdUser)
	} else if err != nil {
		// Handle unexpected errors (like a database connection issue)
		return fmt.Errorf("failed to check user: %w", err)
	} else {
		// User already exists
		fmt.Println("A user with that name already exists!")
		os.Exit(1) // Exit with error code 1
	}

	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("no arguments allowed for command: users")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if s.cfg.CurrentUserName == user.Name {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Println("* " + user.Name)
		}
	}
	return nil
}
