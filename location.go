package main

import "fmt"

/*Location is the source location of a shotgun managed tool. There are
several valid types to choose from, such as git, git_branch, app_store,
dev etc.*/
type Location struct {
	refName      string
	Version      string
	LocationType string `yaml:"type"`
	Name_        string `yaml:"name,omitempty"`
	Path         string `yaml:"path,omitempty"`
}

func (l *Location) Name() string {
	if l.refName != "" {
		return l.refName
	}
	if l.Name_ != "" {
		return l.Name_
	}
	return "unknown"
}

/*LocationStore groups locations and handle the output path of itself.
 */
type LocationStore struct {
	name      string
	Locations map[string]Location
}

func (ls *LocationStore) Namespace() string {
	return fmt.Sprintf("locations.%s", ls.name)
}

func (ls *LocationStore) Path() string {
	return fmt.Sprintf("env/locations/%s.yml", ls.name)
}

func (ls *LocationStore) AddLocation(l Location) {
	locationName := fmt.Sprintf("%s.%s.location", ls.Namespace(), l.Name())
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
