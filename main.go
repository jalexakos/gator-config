package main

import (
	"database/sql"
	"fmt"
	"os"

	command "github.com/jalexakos/gator-config/internal/command"
	config "github.com/jalexakos/gator-config/internal/config"
	"github.com/jalexakos/gator-config/internal/database"
	login "github.com/jalexakos/gator-config/internal/login"
	"github.com/jalexakos/gator-config/internal/register"
	"github.com/jalexakos/gator-config/internal/reset"
	"github.com/jalexakos/gator-config/users"

	_ "github.com/lib/pq"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}
	dbURL := configFile.DbUrl
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	dbQueries := database.New(db)
	state := config.State{
		Cfg: configFile,
		Db:  dbQueries,
	}
	commands := command.Commands{
		Handlers: make(map[string]func(*config.State, command.Command) error),
	}
	commands.Register("login", login.HandlerLogin)
	commands.Register("register", register.HandlerRegister)
	commands.Register("reset", reset.HandlerReset)
	commands.Register("users", users.HandlerUsers)
	cmds := os.Args
	if len(cmds) < 2 {
		fmt.Errorf("Please provide a command")
		os.Exit(1)
	}
	cmdName := cmds[1]
	cmdArgs := cmds[2:]
	cmd := command.Command{
		Name: cmdName,
		Args: cmdArgs,
	}
	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
}
