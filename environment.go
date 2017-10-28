package main

import (
	"fmt"
	"log"
	"path/filepath"
)

/*Environment is a single environment that controls the engines
available and it's configuration. This means we can have several
engines that have the same location and configurations, but they
are part of different environments.
*/
type Environment struct {
	name        string
	Includes    []string `yaml:"includes,omitempty"`
	Description string
	Engines     []Engine
	Frameworks  []string `yaml:"frameworks,omitempty"`
}

/*Path returns the path of the named environment.
 */
func (e *Environment) Path() string {
	return fmt.Sprintf("env/%v.yml", e.name)
}

/*Include includes a location
 */
func (e *Environment) Include(ls LocationStore) string {
	val, err := filepath.Rel(e.Path(), ls.Path())
	if err != nil {
		log.Fatal(err)
	}
	return val
}
