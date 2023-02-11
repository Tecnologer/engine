package engine

import "fmt"

type Metadata interface {
	//GetName returns the name of the game version
	GetName() string
	//GetMaxPlayer returns an integer that represents the maximum of players to play
	GetMaxPlayer() int
	//GetMinPlayer returns an integer that represents the minimum amount of players to start a game
	GetMinPlayer() int
	//GetVersion returns the build version of the plugin
	GetVersion() string
	//String returns a representative's string of the metadata
	String() string
}

func MetadataToString(m Metadata) string {
	return fmt.Sprintf("%s [players: %d-%d] (version: %s)",
		m.GetName(),
		m.GetMinPlayer(),
		m.GetMaxPlayer(),
		m.GetVersion(),
	)
}
