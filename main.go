package main

import (
	"fmt"
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
	name          string
	location      Location
	LocationRef   string `yaml:"location"`
	Apps          map[string]App
	configuration map[string]interface{}
}

func (e *Engine) SetLocation(l *Location) {
	e.location = *l
	e.LocationRef = "@" + l.Name()
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

	// Simple Scenario:
	//  One environment, a project.
	// two engines: nuke and maya.
	// Each engine has it's apps and configurations

	// Goals:
	//  Functioning environment with included engines.
	// Engines have several apps with their locations
	// referenced from a different file also.
	eStore := NewLocationStore("engines")
	aStore := NewLocationStore("apps")

	projectEnv := NewEnvironment("project")
	projectEnv.Description = "Apps and Engines when launching with a project only context."

	projectEnv.Include(eStore)
	projectEnv.Include(aStore)

	// Engine locations
	mayaLocation := Location{
		LocationName: "tk-maya",
		Version:      "v1.0.0",
		LocationType: "app_store",
	}
	eStore.AddLocation(&mayaLocation)

	nukeLocation := Location{
		LocationName: "tk-nuke",
		Version:      "v1.0.0",
		LocationType: "app_store",
	}
	eStore.AddLocation(&nukeLocation)

	// Add some engines to play with
	mayaEngine := Engine{
		name: "tk-maya",
	}
	mayaEngine.SetLocation(&mayaLocation)

	nukeEngine := Engine{
		name: "tk-nuke",
	}
	nukeEngine.SetLocation(&nukeLocation)

	projectEnv.Engines[mayaEngine.name] = mayaEngine
	projectEnv.Engines[nukeEngine.name] = nukeEngine

	fmt.Println(projectEnv.Path(), "\n-------")

	d, err := yaml.Marshal(projectEnv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d))
}
