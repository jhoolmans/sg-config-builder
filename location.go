package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

/*Location is the source location of a shotgun managed tool. There are
several valid types to choose from, such as git, git_branch, app_store,
dev etc.*/
type Location struct {
	refName      string
	Version      string
	LocationType string `yaml:"type"`
	LocationName string `yaml:"name,omitempty"`
	Path         string `yaml:"path,omitempty"`
}

/*Name returns the reference name for the location. It prioritizes the
refName attribute over LocationName. LocationName is still used by the
LocationStore to update it's refName.
*/
func (l *Location) Name() string {
	if l.refName != "" {
		return l.refName
	}
	if l.LocationName != "" {
		return l.LocationName
	}
	return "unknown"
}

/*LocationStore groups locations and handle the output path of itself.
 */
type LocationStore struct {
	name      string
	Locations map[string]*Location
}

/*Namespace returns a prefix used for keys inside the LocationStore.
This allow for readability when included.
*/
func (ls *LocationStore) Namespace() string {
	return fmt.Sprintf("locations.%s", ls.name)
}

/*Path returns a relative path based on where it's executed. Running
the commands from $HOME will result in $HOME/env/locations/`name`.yml.
*/
func (ls *LocationStore) Path() string {
	return fmt.Sprintf("env/locations/%s.yml", ls.name)
}

/*AddLocation adds the given location to the Locations map and renames
the location with a prefix and suffix.
*/
func (ls *LocationStore) AddLocation(l *Location) {
	locationName := fmt.Sprintf("%s.%s.location", ls.Namespace(), l.Name())
	l.refName = locationName
	ls.Locations[locationName] = l
}

/*ToYaml returns a valid yaml output for Shotgun.
 */
func (ls *LocationStore) ToYaml() (string, error) {
	d, err := yaml.Marshal(&ls.Locations)
	return string(d), err
}

/*NewLocationStore initializes a new LocationStore object.
 */
func NewLocationStore(name string) LocationStore {
	ls := LocationStore{name: name}
	ls.Locations = make(map[string]*Location)
	return ls
}
