package main

import(
	"bufio"
    "fmt"
    "strings"
	"os"
	"github.com/coolarif123/pokedexcli/internal/pokecache"
	"time"
)

func startRepl() {
	var config Config
	config.PokemonCaught = make(map[string]Pokemon)
	config.Cache = pokecache.NewCache(5 * time.Second)
	config.initMap = false 
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands(&config)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		textString := strings.Fields(text)

		if len(text) == 0 {
			printPrompt()
			continue
		}

		if textString[0] == "explore" {
			text = textString[0]
			config.Area = textString[1]
		}

		if textString[0] == "catch" {
			text = textString[0]
			config.PokemonName = textString[1]
		}

		if textString[0] == "inspect" {
			text = textString[0]
			config.PokemonName = textString[1]
		}

		if cmd, ok := commands[text]; !ok {
			handleInvalidCmd(text)
		} else {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
		printPrompt()
	}
}

var cliName string = "pokedex"

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands(config *Config) map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    func() error {
				return commandHelp(config)
			},
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
		"map": {
			name:		 "map",
			description: "Displays the names of 20 locations in the Pokemon world. Each subsequent call to map will display the next 20 locations.",
			callback:	 func() error {
				config.Mapb = false
				return commandMap(config)
			},
		},
		"mapb": {
			name:		 "mapb",
		 	description: "Displays the names of the previous 20 locations in the Pokemon world.",
			callback:	 func() error {
				config.Mapb = true
				return commandMap(config)
			},
		},
		"explore": {
			name:		 "explore <area_name>",
		 	description: "explore <area_name> explores the pokemon available to be caught in that area",
			callback:	 func() error {
				return commandExplore(config)
			},
		},
		"catch": {
			name:		 "catch <pokemon_name>",
		 	description: "catch <pokemon_name> throws a pokeball and gives a chance to catch the pokemon",
			callback:	 func() error {
				return commandCatch(config)
			},
		},
		"inspect": {
			name:		 "inspect <pokemon_name>",
		 	description: "inspect <pokemon_name> reveals the stats of a caught pokemon",
			callback:	 func() error {
				return commandInspect(config)
			},
		},
		"pokedex": {
			name:		 "pokedex",
		 	description: "pokedex lists all of the pokemon you have caught",
			callback:	 func() error {
				return commandPokedex(config)
			},
		},
	}
}

type Config struct {
	Mapb        bool
	initMap     bool
	Area	    string
	PokemonName string 
	Previous    *string
	Next 	    *string
	Cache 	    *pokecache.Cache
	PokemonCaught map[string]Pokemon
}

// handleInvalidCmd attempts to recover from a bad command
func handleInvalidCmd(text string) {
    defer printUnknown(text)
}
 
// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
    output := strings.TrimSpace(text)
    output = strings.ToLower(output)
    return output
}