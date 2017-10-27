package main

import "fmt"

/*Location is the source location of a shotgun managed tool. There are
several valid types to choose from, such as git, git_branch, app_store,
dev etc.*/
type Location struct {
	refName      string
	Version      string
	LocationType string `yaml:"type"`
	Name         string `yaml:"name,omitempty"`
	Path         string `yaml:"path,omitempty"`
}

func (l *Location) name() string {
	if l.refName != "" {
		return l.refName
	}
	if l.Name != "" {
		return l.Name
	}
	return "unknown"
}

/*LocationStore groups locations and handle the output path of itself.
 */
type LocationStore struct {
	name      string
	Locations map[string]Location
}

func (ls *LocationStore) namespace() string {
	return fmt.Sprintf("locations.%s", ls.name)
}

func (ls *LocationStore) path() string {
	return fmt.Sprintf("env/locations/%s.yml", ls.name)
}

func (ls *LocationStore) addLocation(l Location) {
	locationName := fmt.Sprintf("%s.%s.location", ls.namespace(), l.name())
	l.refName = locationName
	ls.Locations[locationName] = l
}

/*NewLocationStore initializes a new LocationStore object.
 */
func NewLocationStore(name string) LocationStore {
	ls := LocationStore{name: name}
	ls.Locations = make(map[string]Location)
	return ls
}
