package factory

import (
	"github.com/DumbNoxx/goxe/pkg/exporter"
)

func GetShipper(destination string) exporter.Shipper {
	switch destination {
	case "socket":
		return &exporter.JsonManager{}
	default:
		return &exporter.JsonManager{}
	}
}
