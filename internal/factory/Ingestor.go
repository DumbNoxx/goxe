package factory

import (
	"fmt"
	"path/filepath"

	"github.com/DumbNoxx/goxe/pkg/ingestor"
)

func GetIngestor(route string) (manager ingestor.File, err error) {
	extension := filepath.Ext(route)
	switch extension {
	case ".log", ".txt":
		return &ingestor.NormalizedManager{}, nil
	default:
		return nil, fmt.Errorf("[Goxe] unsupported file extension: %s", extension)
	}
}
