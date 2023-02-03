/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TestSpecs struct {
	Specs []TestSpecsSpecsInner `json:"specs,omitempty"`
}

// AssertTestSpecsRequired checks if the required fields are not zero-ed
func AssertTestSpecsRequired(obj TestSpecs) error {
	for _, el := range obj.Specs {
		if err := AssertTestSpecsSpecsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseTestSpecsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestSpecs (e.g. [][]TestSpecs), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestSpecsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestSpecs, ok := obj.(TestSpecs)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestSpecsRequired(aTestSpecs)
	})
}
