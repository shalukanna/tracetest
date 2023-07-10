/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LinterResourceRule struct {
	Id string `json:"id,omitempty"`

	Weight int32 `json:"weight,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	ErrorDescription string `json:"errorDescription,omitempty"`

	Tips []string `json:"tips,omitempty"`

	ErrorLevel string `json:"errorLevel,omitempty"`
}

// AssertLinterResourceRuleRequired checks if the required fields are not zero-ed
func AssertLinterResourceRuleRequired(obj LinterResourceRule) error {
	return nil
}

// AssertRecurseLinterResourceRuleRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LinterResourceRule (e.g. [][]LinterResourceRule), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLinterResourceRuleRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLinterResourceRule, ok := obj.(LinterResourceRule)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLinterResourceRuleRequired(aLinterResourceRule)
	})
}