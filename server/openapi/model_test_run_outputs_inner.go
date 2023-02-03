/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TestRunOutputsInner struct {
	Name string `json:"name,omitempty"`

	SpanId string `json:"spanId,omitempty"`

	Value string `json:"value,omitempty"`
}

// AssertTestRunOutputsInnerRequired checks if the required fields are not zero-ed
func AssertTestRunOutputsInnerRequired(obj TestRunOutputsInner) error {
	return nil
}

// AssertRecurseTestRunOutputsInnerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestRunOutputsInner (e.g. [][]TestRunOutputsInner), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestRunOutputsInnerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestRunOutputsInner, ok := obj.(TestRunOutputsInner)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestRunOutputsInnerRequired(aTestRunOutputsInner)
	})
}
