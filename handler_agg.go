package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/raffkelly/gator/internal/database"
)

func handlerAgg(_ *state, _ command) error {
	feedUrl := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

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
