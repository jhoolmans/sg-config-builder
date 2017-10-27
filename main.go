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
	location    Location
	LocationRef string `yaml:"location"`
}

/*Engine represents a dcc application with shotgun configurations.
Most importantly in a pipeline configuration this will also hold the
source location of this engine.
*/
type Engine struct {
	location      Location
	LocationRef   string `yaml:"location"`
	Apps          map[string]App
	Configuration map[string]interface{}
}

/*Includes represents a list of files relative to itself.
Every yaml document is able to include different yaml documents.
*/
type Includes map[string]string

/*Environment is a single environment that controls the engines
available and it's configuration. This means we can have several
engines that have the same location and configurations, but they
are part of different environments.
*/
type Environment struct {
	name        string
	Description string
	Engines     []Engine
	Frameworks  []string
}

/*Templates represents the templates document holding all the
global templates. Keys are mapped to interface{} since they
can be a string, map, or even a map of maps. Same with Paths.
*/
type Templates struct {
	Keys    map[string]interface{}
	Paths   map[string]interface{}
	Strings map[string]string
}

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
		location: mayaLocation,
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
	data, err := ioutil.ReadFile("yaml_files/engine-locations.yml")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	// Named locations (eg. tk-maya)
	var nl map[string]Location

	err = yaml.Unmarshal(data, &nl)
	if err != nil {
		log.Fatalf("Could not read yaml document: %v", err)
	}

	fmt.Printf("--- Locations:\n%v\n\n", nl)

	//
	// LocationStore
	store := NewLocationStore("engines")
	for _, l := range nl {
		store.addLocation(l)
	}

	fmt.Printf("--- Location store:\n%v:\n%v\n\n", store.path(), store)

	//
	// Now print them back out like we read them
	d, err = yaml.Marshal(&store.Locations)
	if err != nil {
		log.Fatalf("Error converting to yaml: %v", err)
	}

	fmt.Println("Output:\n", string(d))

	err = ioutil.WriteFile("yaml_files/engine-locations-output.yml", d, 0777)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}
