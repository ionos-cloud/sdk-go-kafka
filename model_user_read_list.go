/*
 * Kafka as a Service API
 *
 * An managed Apache Kafka cluster is designed to be highly fault-tolerant and scalable, allowing large volumes of data to be ingested, stored, and processed in real-time. By distributing data across multiple brokers, Kafka achieves high throughput and low latency, making it suitable for applications requiring real-time data processing and analytics.
 *
 * API version: 1.2.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// UserReadList struct for UserReadList
type UserReadList struct {
	// ID of the list of User resources.
	Id *string `json:"id"`
	// The type of the resource.
	Type *string `json:"type"`
	// The URL of the list of User resources.
	Href *string `json:"href"`
	// The list of User resources.
	Items *[]UserRead `json:"items,omitempty"`
}

// NewUserReadList instantiates a new UserReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserReadList(id string, type_ string, href string) *UserReadList {
	this := UserReadList{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href

	return &this
}

// NewUserReadListWithDefaults instantiates a new UserReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserReadListWithDefaults() *UserReadList {
	this := UserReadList{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserReadList) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *UserReadList) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *UserReadList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserReadList) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *UserReadList) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *UserReadList) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserReadList) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *UserReadList) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *UserReadList) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []UserRead will be returned
func (o *UserReadList) GetItems() *[]UserRead {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserReadList) GetItemsOk() (*[]UserRead, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *UserReadList) SetItems(v []UserRead) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *UserReadList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o UserReadList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableUserReadList struct {
	value *UserReadList
	isSet bool
}

func (v NullableUserReadList) Get() *UserReadList {
	return v.value
}

func (v *NullableUserReadList) Set(val *UserReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableUserReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableUserReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserReadList(val *UserReadList) *NullableUserReadList {
	return &NullableUserReadList{value: val, isSet: true}
}

func (v NullableUserReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
