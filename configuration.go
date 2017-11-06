/*
Configurations are a collection of Setting objects that describe the
attribute type it contains.
*/
package main

/*Configuration is an alias for map[string]SettingKey.
 */
type Configuration map[string]SettingKey

/*SettingTypedValue is a structure that can be used recursively.
This can define values hold by a python dict object and can be of
any type supported by Shotgun/python.
For instance, 'hook' is a Shotgun type and 'dict' is a plain python
type.
*/
type SettingTypedValue struct {
	Type  string                       `json:"type"`
	Items map[string]SettingTypedValue `json:"items,omitempty"`
}

/*SettingKey defines a setting 'option' for apps. Internally the apps
or engines will read from this SettingKey and will use the default_value
if omitted.
*/
type SettingKey struct {
	Type         string             `json:"type"`
	Description  string             `json:"description"`
	DefaultValue interface{}        `json:"default_value"`
	AllowsEmpty  bool               `json:"allows_empty,omitempty"`
	Values       *SettingTypedValue `json:"values,omitempty"`
}

/*AppInfo is used internally to track the configuration for a specific
app. AppName can be used as an identifier.
*/
type AppInfo struct {
	AppName       string
	Configuration Configuration
}
