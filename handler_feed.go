package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/raffkelly/gator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf(`usage: addfeed "name" "url"`)
	}
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	currentUserId := currentUser.ID
	feedTitle := cmd.Arguments[0]
	feedUrl := cmd.Arguments[1]

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedTitle,
		Url:       feedUrl,
		UserID:    currentUserId,
	}
	newFeed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	fmt.Println("Feed successfully added.")
	fmt.Printf("%+v\n", newFeed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("no arguments allowed for feeds command")
	}
	var creator string
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		creator, err = s.db.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("===========\n")
		fmt.Printf("Feed Name: %v\n", feed.Name)
		fmt.Printf("Feed URL: %v\n", feed.Url)
		fmt.Printf("Created by User: %v\n", creator)
	}
	return nil
}
