package main

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

/*SetLocation sets the Location and creates the proper reference.
Not that this should be called _after_ the location has been added
to the LocationStore to get the right reference link.
*/
func (e *Engine) SetLocation(l *Location) {
	e.location = *l
	e.LocationRef = "@" + l.Name()
}
