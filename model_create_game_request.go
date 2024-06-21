/*
jabali-api

API Proposal for Jabali.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jabalisdkgo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateGameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGameRequest{}

// CreateGameRequest struct for CreateGameRequest
type CreateGameRequest struct {
	Name string `json:"name"`
	Displayname string `json:"displayname"`
	Url string `json:"url"`
}

type _CreateGameRequest CreateGameRequest

// NewCreateGameRequest instantiates a new CreateGameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGameRequest(name string, displayname string, url string) *CreateGameRequest {
	this := CreateGameRequest{}
	this.Name = name
	this.Displayname = displayname
	this.Url = url
	return &this
}

// NewCreateGameRequestWithDefaults instantiates a new CreateGameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGameRequestWithDefaults() *CreateGameRequest {
	this := CreateGameRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateGameRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGameRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGameRequest) SetName(v string) {
	o.Name = v
}

// GetDisplayname returns the Displayname field value
func (o *CreateGameRequest) GetDisplayname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value
// and a boolean to check if the value has been set.
func (o *CreateGameRequest) GetDisplaynameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Displayname, true
}

// SetDisplayname sets field value
func (o *CreateGameRequest) SetDisplayname(v string) {
	o.Displayname = v
}

// GetUrl returns the Url field value
func (o *CreateGameRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateGameRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateGameRequest) SetUrl(v string) {
	o.Url = v
}

func (o CreateGameRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["displayname"] = o.Displayname
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *CreateGameRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"displayname",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateGameRequest := _CreateGameRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGameRequest)

	if err != nil {
		return err
	}

	*o = CreateGameRequest(varCreateGameRequest)

	return err
}

type NullableCreateGameRequest struct {
	value *CreateGameRequest
	isSet bool
}

func (v NullableCreateGameRequest) Get() *CreateGameRequest {
	return v.value
}

func (v *NullableCreateGameRequest) Set(val *CreateGameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGameRequest(val *CreateGameRequest) *NullableCreateGameRequest {
	return &NullableCreateGameRequest{value: val, isSet: true}
}

func (v NullableCreateGameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


