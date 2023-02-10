package engine

import "fmt"

type Metadata interface {
	GetName() string
	GetMaxPlayer() int
	GetMinPlayer() int
	GetVersion() string
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
