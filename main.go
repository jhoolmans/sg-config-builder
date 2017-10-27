package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

/*App represents a single tool which belongs to a single or multiple
engines. It stores information such as the source location and
configuration specific for a context.
*/
type App struct {
	Location Location
}

/*Engine represents a dcc application with shotgun configurations.
Most importantly in a pipeline configuration this will also hold the
source location of this engine.
*/
type Engine struct {
	Location      Location
	Apps          map[string]App
	Configuration map[string]interface{}
}

/*Engines represents multiple Engine objects.
 */
type Engines []Engine

func main() {
	fmt.Println("------\nConfig Builder\n------")

	//
	// Example of creating an Engine object with a git Locations
	mayaLocation := Location{
		Version:      "v0.8.3",
		LocationType: "app_store",
		Name:         "tk-maya",
		Path:         "https://github.com/shotgunsoftware/tk-maya.git",
	}
	mayaEngine := Engine{
		Location: mayaLocation,
	}

	// Dump the mayaEngine formatted as yaml
	//
	d, err := yaml.Marshal(&mayaEngine)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("-- engine dump:\n%s\n\n", string(d))

	//
	// Example of loading an engines yaml file
	data, err := ioutil.ReadFile("yaml_files/engines.yml")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	// Named locations (eg. common.engines.tk-maya.location)
	var l map[string]Location

	err = yaml.Unmarshal(data, &l)
	if err != nil {
		log.Fatalf("Could not read yaml document: %v", err)
	}

	fmt.Printf("--- Locations:\n%v\n\n", l)

	//
	// Now print them back out like we read them
	d, err = yaml.Marshal(&l)
	if err != nil {
		log.Fatalf("Error converting to yaml: %v", err)
	}

	fmt.Println("Output:\n", string(d))

	err = ioutil.WriteFile("yaml_files/engines-output.yml", d, 0777)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}
