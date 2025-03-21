package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/raffkelly/gator/internal/config"
	"github.com/raffkelly/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	// read the config file in the home directory
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// set the state of the program based on the config
	currentState := &state{
		cfg: &cfg,
	}

	// open a connection to my database and get queries
	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	currentState.db = dbQueries

	// register the login command in the commands struct
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("at least one argument required")
	}

	cmd := command{
		Name:      args[1],
		Arguments: args[2:],
	}

	err = cmds.run(currentState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
