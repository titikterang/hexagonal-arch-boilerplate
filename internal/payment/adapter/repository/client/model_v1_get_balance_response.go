/*
Fast Campus Pay (Wallet API Endpoint)

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
Contact: azwar.nrst@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package repository

import (
	"encoding/json"
)

// checks if the V1GetBalanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GetBalanceResponse{}

// V1GetBalanceResponse struct for V1GetBalanceResponse
type V1GetBalanceResponse struct {
	UserId  *string  `json:"userId,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
}

// NewV1GetBalanceResponse instantiates a new V1GetBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GetBalanceResponse() *V1GetBalanceResponse {
	this := V1GetBalanceResponse{}
	return &this
}

// NewV1GetBalanceResponseWithDefaults instantiates a new V1GetBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GetBalanceResponseWithDefaults() *V1GetBalanceResponse {
	this := V1GetBalanceResponse{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *V1GetBalanceResponse) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GetBalanceResponse) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *V1GetBalanceResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *V1GetBalanceResponse) SetUserId(v string) {
	o.UserId = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *V1GetBalanceResponse) GetBalance() float64 {
	if o == nil || IsNil(o.Balance) {
		var ret float64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GetBalanceResponse) GetBalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *V1GetBalanceResponse) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float64 and assigns it to the Balance field.
func (o *V1GetBalanceResponse) SetBalance(v float64) {
	o.Balance = &v
}

func (o V1GetBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GetBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	return toSerialize, nil
}

type NullableV1GetBalanceResponse struct {
	value *V1GetBalanceResponse
	isSet bool
}

func (v NullableV1GetBalanceResponse) Get() *V1GetBalanceResponse {
	return v.value
}

func (v *NullableV1GetBalanceResponse) Set(val *V1GetBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GetBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GetBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GetBalanceResponse(val *V1GetBalanceResponse) *NullableV1GetBalanceResponse {
	return &NullableV1GetBalanceResponse{value: val, isSet: true}
}

func (v NullableV1GetBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GetBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}