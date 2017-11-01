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
	Engines     map[string]Engine
	Frameworks  []string `yaml:"frameworks,omitempty"`
}

/*NewEnvironment returns a new Environment object. This func takes
care of the initialization of it's members.
*/
func NewEnvironment(name string) Environment {
	e := Environment{
		name: name,
	}
	e.Engines = make(map[string]Engine)
	return e
}

/*AddEngine adds an engine by the name of the supplied engine.
 */
func (e *Environment) AddEngine(engine *Engine) {
	e.Engines[engine.name] = *engine
}

/*Path returns the path of the named environment.
 */
func (e *Environment) Path() string {
	return fmt.Sprintf("env/%v.yml", e.name)
}

/*Include includes a location
 */
func (e *Environment) Include(ls LocationStore) {
	val, err := filepath.Rel(e.Path(), ls.Path())
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range e.Includes {
		if b == val {
			return
		}
	}
	e.Includes = append(e.Includes, val)
}
