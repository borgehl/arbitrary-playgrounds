package contract

import "time"

type Event struct {
	Name string
	Date time.Time
}

type VenueInfo struct {
	Name string
	Link string
	Tags []string
}

type PluginOutput struct {
	VenueInfo VenueInfo
	Events    []Event
}
