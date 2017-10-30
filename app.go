package main

/*App represents a single tool which belongs to a single or multiple
engines. It stores information such as the source location and
configuration specific for a context.
*/
type App struct {
	location    Location
	LocationRef string `yaml:"location"`
}
