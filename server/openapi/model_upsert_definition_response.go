/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UpsertDefinitionResponse struct {

	// resource ID
	Id string `json:"id,omitempty"`

	// resource type
	Type string `json:"type,omitempty"`
}

// AssertUpsertDefinitionResponseRequired checks if the required fields are not zero-ed
func AssertUpsertDefinitionResponseRequired(obj UpsertDefinitionResponse) error {
	return nil
}

// AssertRecurseUpsertDefinitionResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UpsertDefinitionResponse (e.g. [][]UpsertDefinitionResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUpsertDefinitionResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUpsertDefinitionResponse, ok := obj.(UpsertDefinitionResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUpsertDefinitionResponseRequired(aUpsertDefinitionResponse)
	})
}