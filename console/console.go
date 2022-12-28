package console

import (
	"ch/kirari/animeApi/setups"
)

func Console(doConsole string, doSeed string) {
	println("Console mode is enabled")
	switch doConsole {
	case "seed":
		switch doSeed {
		case "all":
			setups.SeedDatabase()
			setups.SeedSearch()
		case "database":
			setups.SeedDatabase()
		case "search":
			setups.SeedSearch()
		default:
			println("correct parameter -seed was not provided (options: all, database, search)")
		}
	}
}
