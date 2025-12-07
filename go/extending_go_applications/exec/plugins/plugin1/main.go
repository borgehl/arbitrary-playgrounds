package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/borgehl/arbitrary-playgrounds/go/extending_go_applications/exec/contract"
)

func main() {

	output := contract.PluginOutput{
		VenueInfo: contract.VenueInfo{
			Name: "Plugin Venue Test",
			Link: "example.com",
			Tags: []string{"music", "recuring"},
		},
		Events: make([]contract.Event, 0),
	}
	today := time.Now()

	for i := range 5 {
		output.Events = append(output.Events, contract.Event{
			Name: fmt.Sprintf("Event %d", i),
			Date: today.AddDate(0, 0, 1),
		})
	}

	buf, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	_, err = os.Stdout.Write(buf)
	if err != nil {
		panic(err)
	}
}
