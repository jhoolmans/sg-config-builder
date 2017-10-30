package main

import (
	"io/ioutil"
	"strings"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestLocationMarshal(t *testing.T) {
	//
	// Example of creating an Engine object with a git Locations
	ml := Location{
		Version:      "v0.8.3",
		LocationType: "git",
		LocationName: "tk-maya",
		Path:         "https://github.com/shotgunsoftware/tk-maya.git",
	}

	// Dump the mayaEngine formatted as yaml
	//
	d, err := yaml.Marshal(&ml)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	// Check for path option
	//
	if !strings.Contains(string(d), "path:") {
		t.Error("Location.Type not exposed.")
	}

	if strings.Contains(string(d), "locationpath:") {
		t.Error("LocationType exposed instead of 'type'.")
	}

	//
	// Appstore instance
	appstore := Location{
		Version:      "v0.1.0",
		LocationType: "app_store",
		LocationName: "tk-maya",
	}

	// Dump the mayaEngine formatted as yaml
	//
	d, err = yaml.Marshal(&appstore)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if strings.Contains(string(d), "path:") {
		t.Error("Type should be omitted when empty.")
	}
}

func TestLocationName(t *testing.T) {
	loc := Location{
		Version:      "v1.0.0b",
		LocationType: "app_store",
	}
	if loc.Name() != "unknown" {
		t.Errorf("Expected loc.name to be 'unknown', it's '%v' instead.", loc.Name())
	}

	// Override (app) location name
	loc.LocationName = "Hallo"
	if loc.Name() != "Hallo" {
		t.Errorf("Expected loc.name to be 'Hallo', it's '%v' instead.", loc.Name())
	}

	// Set refName to be used when printing in Yaml
	loc.refName = "Hoi"
	if loc.Name() != "Hoi" {
		t.Errorf("Expected loc.name to be 'Hoi', it's '%v' instead.", loc.Name())
	}
}

func TestLocationStoreMarshal(t *testing.T) {
	inputLocations := "yaml_files/engine-locations.yml"

	//
	// Example of loading an engines yaml file
	data, err := ioutil.ReadFile(inputLocations)
	if err != nil {
		t.Errorf("Could not read file: %v", err)
	}

	// Named locations (eg. tk-maya)
	var nl map[string]Location

	// Extracts locations into object
	err = yaml.Unmarshal(data, &nl)
	if err != nil {
		t.Errorf("Could not read yaml document: %v", err)
	}

	if len(nl) == 0 {
		t.Error("Could not read locations from engine-locations.yml")
	}

	//
	// LocationStore
	store := NewLocationStore("engines")
	for _, l := range nl {
		store.AddLocation(&l)
	}

	if store.Path() != "env/locations/engines.yml" {
		t.Error("Invalid store path.")
	}

	if len(store.Locations) == 0 {
		t.Error("Locations were not stored in LocationStore..")
	}

	//
	// Marshal LocationStore
	d, err := store.ToYaml()
	if err != nil {
		t.Errorf("Error converting store.Locations to yaml: %v", err)
	}

	if !strings.Contains(string(d), "locations.engines.") {
		t.Error("No locations were dumped with yaml.")
	}
}
