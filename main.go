package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

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

	fmt.Println(eStore.Path(), "\n------")

	d, err := eStore.ToYaml()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d))

	fmt.Println(projectEnv.Path(), "\n-------")

	m, err := yaml.Marshal(projectEnv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(m))
}
