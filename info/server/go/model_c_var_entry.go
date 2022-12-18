/*
 * GoBun Replay File Info
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CVarEntry struct {

	CVarId int32 `json:"CVarId,omitempty"`

	Value string `json:"Value,omitempty"`
}

// AssertCVarEntryRequired checks if the required fields are not zero-ed
func AssertCVarEntryRequired(obj CVarEntry) error {
	return nil
}

// AssertRecurseCVarEntryRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CVarEntry (e.g. [][]CVarEntry), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCVarEntryRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCVarEntry, ok := obj.(CVarEntry)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCVarEntryRequired(aCVarEntry)
	})
}
