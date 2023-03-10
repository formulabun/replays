/*
 * GoBun File Info
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PlayerEntry struct {

	Name string `json:"Name,omitempty"`

	Skin string `json:"Skin,omitempty"`

	Color string `json:"Color,omitempty"`

	Score float32 `json:"Score,omitempty"`

	Speed float32 `json:"Speed,omitempty"`

	Weight float32 `json:"Weight,omitempty"`
}

// AssertPlayerEntryRequired checks if the required fields are not zero-ed
func AssertPlayerEntryRequired(obj PlayerEntry) error {
	return nil
}

// AssertRecursePlayerEntryRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PlayerEntry (e.g. [][]PlayerEntry), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePlayerEntryRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPlayerEntry, ok := obj.(PlayerEntry)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPlayerEntryRequired(aPlayerEntry)
	})
}
