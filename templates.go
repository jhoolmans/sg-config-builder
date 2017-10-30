package main

/*Templates represents the templates document holding all the
global templates. Keys are mapped to interface{} since they
can be a string, map, or even a map of maps. Same with Paths.
*/
type Templates struct {
	Keys    map[string]interface{}
	Paths   map[string]interface{}
	Strings map[string]string
}
