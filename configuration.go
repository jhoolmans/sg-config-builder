/*
Configurations are a collection of Setting objects that describe the
attribute type it contains.
*/
package main

type Configuration map[string]SettingKey

type SettingTypedValue struct {
	Type  string                       `json:"type"`
	Items map[string]SettingTypedValue `json:"items,omitempty"`
}

type SettingKey struct {
	Type         string             `json:"type"`
	Description  string             `json:"description"`
	DefaultValue interface{}        `json:"default_value"`
	AllowsEmpty  bool               `json:"allows_empty,omitempty"`
	Values       *SettingTypedValue `json:"values,omitempty"`
}

type AppInfo struct {
	AppName       string
	Configuration Configuration
}
