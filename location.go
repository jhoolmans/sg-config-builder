package main

/*Location is the source location of a shotgun managed tool. There are
several valid types to choose from, such as git, git_branch, app_store,
dev etc.*/
type Location struct {
	Version      string
	LocationType string `yaml:"type"`
	Name         string
	Path         string `yaml:"path,omitempty"`
}
