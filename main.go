package main

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
	location      Location
	LocationRef   string `yaml:"location"`
	Apps          map[string]App
	Configuration map[string]interface{}
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
	//fmt.Println("------\nConfig Builder\n------")

}
