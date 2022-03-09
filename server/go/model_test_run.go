/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TestRun struct {

	// ID
	Id string `json:"id,omitempty"`
}

// AssertTestRunRequired checks if the required fields are not zero-ed
func AssertTestRunRequired(obj TestRun) error {
	return nil
}

// AssertRecurseTestRunRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestRun (e.g. [][]TestRun), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestRunRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestRun, ok := obj.(TestRun)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestRunRequired(aTestRun)
	})
}