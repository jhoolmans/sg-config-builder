package main

/*Engine represents a dcc application with shotgun configurations.
Most importantly in a pipeline configuration this will also hold the
source location of this engine.
*/
type Engine struct {
	name          string
	location      Location
	LocationRef   string `json:"location"`
	Apps          map[string]App
	configuration Configuration
}

/*NewEngine constructs an Engine based on a location. It can be renamed
after building. If not, it will default to the location's name.
*/
func NewEngine(l *Location) Engine {
	e := Engine{}
	e.name = l.LocationName

	e.Apps = make(map[string]App)
	e.configuration = make(Configuration)

	// Set location as usual
	e.SetLocation(l)
	return e
}

/*SetLocation sets the Location and creates the proper reference.
Not that this should be called _after_ the location has been added
to the LocationStore to get the right reference link.
*/
func (e *Engine) SetLocation(l *Location) {
	e.location = *l
	e.LocationRef = "@" + l.Name()
}
