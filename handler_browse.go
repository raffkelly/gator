package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/raffkelly/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	postLimit := 2
	var err error
	if len(cmd.Arguments) == 1 {
		postLimit, err = strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return err
		}
	}
	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(postLimit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil {
		return err
	}
	fmt.Printf("Here are the %v most recent posts for user: %v\n", postLimit, user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.Feedname)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("	%v\n", post.Description.String)
		fmt.Printf("========================\n")

	}
	return nil
}
