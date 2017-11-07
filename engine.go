package main

import (
	"encoding/json"
)

/*Engine represents a dcc application with shotgun configurations.
Most importantly in a pipeline configuration this will also hold the
source location of this engine.
*/
type Engine struct {
	name          string
	location      Location
	LocationRef   string `json:"location"`
	Apps          map[string]App
	configuration map[string]interface{}
}

/*NewEngine constructs an Engine based on a location. It can be renamed
after building. If not, it will default to the location's name.
*/
func NewEngine(l *Location) Engine {
	e := Engine{}
	e.Init()
	e.name = l.LocationName

	// Set location as usual
	e.SetLocation(l)
	return e
}

func (e *Engine) Init() {
	e.Apps = make(map[string]App)
	e.configuration = make(map[string]interface{})
}

/*SetLocation sets the Location and creates the proper reference.
Not that this should be called _after_ the location has been added
to the LocationStore to get the right reference link.
*/
func (e *Engine) SetLocation(l *Location) {
	e.location = *l
	e.LocationRef = "@" + l.Name()
}

/*MarshalJSON override. This basically flattens the configuration to be
on the same level as the other keys.
*/
func (e Engine) MarshalJSON() ([]byte, error) {
	rawData := make(map[string]interface{})
	rawData["location"] = e.LocationRef
	rawData["apps"] = e.Apps

	// loop through all configuration keys and add them to rawData
	for k, v := range e.configuration {
		rawData[k] = v
	}
	return json.Marshal(rawData)
}

/*UnmarshalJSON override. Reads all the data into a map[string]interface
and sets anything that isn't location/apps into configuration.

After unmarshal-ing to an Engine, the Engine won't have a valid Location
reference object. In order to restore the connection, we have to relink
it through an 'engine' LocationStore.
*/
func (e *Engine) UnmarshalJSON(b []byte) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	if val, ok := data["apps"]; ok {
		// check to see if value is an empty map
		if len(val.(map[string]interface{})) != 0 {
			e.Apps = val.(map[string]App)
		}
		delete(data, "apps")
	}
	if val, ok := data["location"]; ok {
		e.LocationRef = val.(string)
		delete(data, "location")
	}
	for k, v := range data {
		e.configuration[k] = v
	}
	return nil
}
