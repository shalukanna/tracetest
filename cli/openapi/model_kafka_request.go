/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the KafkaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KafkaRequest{}

// KafkaRequest struct for KafkaRequest
type KafkaRequest struct {
	BrokerUrls      []string             `json:"brokerUrls,omitempty"`
	Topic           *string              `json:"topic,omitempty"`
	Authentication  *KafkaAuthentication `json:"authentication,omitempty"`
	SslVerification *bool                `json:"sslVerification,omitempty"`
	Headers         []KafkaMessageHeader `json:"headers,omitempty"`
	MessageKey      *string              `json:"messageKey,omitempty"`
	MessageValue    *string              `json:"messageValue,omitempty"`
}

// NewKafkaRequest instantiates a new KafkaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaRequest() *KafkaRequest {
	this := KafkaRequest{}
	var sslVerification bool = false
	this.SslVerification = &sslVerification
	return &this
}

// NewKafkaRequestWithDefaults instantiates a new KafkaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaRequestWithDefaults() *KafkaRequest {
	this := KafkaRequest{}
	var sslVerification bool = false
	this.SslVerification = &sslVerification
	return &this
}

// GetBrokerUrls returns the BrokerUrls field value if set, zero value otherwise.
func (o *KafkaRequest) GetBrokerUrls() []string {
	if o == nil || isNil(o.BrokerUrls) {
		var ret []string
		return ret
	}
	return o.BrokerUrls
}

// GetBrokerUrlsOk returns a tuple with the BrokerUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetBrokerUrlsOk() ([]string, bool) {
	if o == nil || isNil(o.BrokerUrls) {
		return nil, false
	}
	return o.BrokerUrls, true
}

// HasBrokerUrls returns a boolean if a field has been set.
func (o *KafkaRequest) HasBrokerUrls() bool {
	if o != nil && !isNil(o.BrokerUrls) {
		return true
	}

	return false
}

// SetBrokerUrls gets a reference to the given []string and assigns it to the BrokerUrls field.
func (o *KafkaRequest) SetBrokerUrls(v []string) {
	o.BrokerUrls = v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *KafkaRequest) GetTopic() string {
	if o == nil || isNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetTopicOk() (*string, bool) {
	if o == nil || isNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *KafkaRequest) HasTopic() bool {
	if o != nil && !isNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *KafkaRequest) SetTopic(v string) {
	o.Topic = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *KafkaRequest) GetAuthentication() KafkaAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret KafkaAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetAuthenticationOk() (*KafkaAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *KafkaRequest) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given KafkaAuthentication and assigns it to the Authentication field.
func (o *KafkaRequest) SetAuthentication(v KafkaAuthentication) {
	o.Authentication = &v
}

// GetSslVerification returns the SslVerification field value if set, zero value otherwise.
func (o *KafkaRequest) GetSslVerification() bool {
	if o == nil || isNil(o.SslVerification) {
		var ret bool
		return ret
	}
	return *o.SslVerification
}

// GetSslVerificationOk returns a tuple with the SslVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetSslVerificationOk() (*bool, bool) {
	if o == nil || isNil(o.SslVerification) {
		return nil, false
	}
	return o.SslVerification, true
}

// HasSslVerification returns a boolean if a field has been set.
func (o *KafkaRequest) HasSslVerification() bool {
	if o != nil && !isNil(o.SslVerification) {
		return true
	}

	return false
}

// SetSslVerification gets a reference to the given bool and assigns it to the SslVerification field.
func (o *KafkaRequest) SetSslVerification(v bool) {
	o.SslVerification = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *KafkaRequest) GetHeaders() []KafkaMessageHeader {
	if o == nil || isNil(o.Headers) {
		var ret []KafkaMessageHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetHeadersOk() ([]KafkaMessageHeader, bool) {
	if o == nil || isNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *KafkaRequest) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []KafkaMessageHeader and assigns it to the Headers field.
func (o *KafkaRequest) SetHeaders(v []KafkaMessageHeader) {
	o.Headers = v
}

// GetMessageKey returns the MessageKey field value if set, zero value otherwise.
func (o *KafkaRequest) GetMessageKey() string {
	if o == nil || isNil(o.MessageKey) {
		var ret string
		return ret
	}
	return *o.MessageKey
}

// GetMessageKeyOk returns a tuple with the MessageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetMessageKeyOk() (*string, bool) {
	if o == nil || isNil(o.MessageKey) {
		return nil, false
	}
	return o.MessageKey, true
}

// HasMessageKey returns a boolean if a field has been set.
func (o *KafkaRequest) HasMessageKey() bool {
	if o != nil && !isNil(o.MessageKey) {
		return true
	}

	return false
}

// SetMessageKey gets a reference to the given string and assigns it to the MessageKey field.
func (o *KafkaRequest) SetMessageKey(v string) {
	o.MessageKey = &v
}

// GetMessageValue returns the MessageValue field value if set, zero value otherwise.
func (o *KafkaRequest) GetMessageValue() string {
	if o == nil || isNil(o.MessageValue) {
		var ret string
		return ret
	}
	return *o.MessageValue
}

// GetMessageValueOk returns a tuple with the MessageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetMessageValueOk() (*string, bool) {
	if o == nil || isNil(o.MessageValue) {
		return nil, false
	}
	return o.MessageValue, true
}

// HasMessageValue returns a boolean if a field has been set.
func (o *KafkaRequest) HasMessageValue() bool {
	if o != nil && !isNil(o.MessageValue) {
		return true
	}

	return false
}

// SetMessageValue gets a reference to the given string and assigns it to the MessageValue field.
func (o *KafkaRequest) SetMessageValue(v string) {
	o.MessageValue = &v
}

func (o KafkaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KafkaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BrokerUrls) {
		toSerialize["brokerUrls"] = o.BrokerUrls
	}
	if !isNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !isNil(o.SslVerification) {
		toSerialize["sslVerification"] = o.SslVerification
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !isNil(o.MessageKey) {
		toSerialize["messageKey"] = o.MessageKey
	}
	if !isNil(o.MessageValue) {
		toSerialize["messageValue"] = o.MessageValue
	}
	return toSerialize, nil
}

type NullableKafkaRequest struct {
	value *KafkaRequest
	isSet bool
}

func (v NullableKafkaRequest) Get() *KafkaRequest {
	return v.value
}

func (v *NullableKafkaRequest) Set(val *KafkaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaRequest(val *KafkaRequest) *NullableKafkaRequest {
	return &NullableKafkaRequest{value: val, isSet: true}
}

func (v NullableKafkaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}