package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"

	"github.com/borgehl/arbitrary-playgrounds/go/extending_go_applications/exec/contract"
)

func main() {

	// each plugin would be named in a config file
	// at server startup
	var output bytes.Buffer
	cmd := exec.Command("./plugin1.x")
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	po := contract.PluginOutput{}
	buf, err := io.ReadAll(&output)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf, &po)
	if err != nil {
		panic(err)
	}

	// handle the data
	fmt.Printf("Venue info:%+v\n", po.VenueInfo)
	for _, event := range po.Events {
		fmt.Printf("\tevent: %+v\n", event)
	}

}
