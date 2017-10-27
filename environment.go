package main

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
	Frameworks  []string
}
