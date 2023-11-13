/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"encoding/json"
	"time"
)

// checks if the RpmPackageLangpacksResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmPackageLangpacksResponse{}

// RpmPackageLangpacksResponse PackageLangpacks serializer.
type RpmPackageLangpacksResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Langpacks matches.
	Matches map[string]interface{} `json:"matches"`
	// Langpacks digest.
	Digest NullableString `json:"digest"`
	AdditionalProperties map[string]interface{}
}

type _RpmPackageLangpacksResponse RpmPackageLangpacksResponse

// NewRpmPackageLangpacksResponse instantiates a new RpmPackageLangpacksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmPackageLangpacksResponse(matches map[string]interface{}, digest NullableString) *RpmPackageLangpacksResponse {
	this := RpmPackageLangpacksResponse{}
	this.Matches = matches
	this.Digest = digest
	return &this
}

// NewRpmPackageLangpacksResponseWithDefaults instantiates a new RpmPackageLangpacksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmPackageLangpacksResponseWithDefaults() *RpmPackageLangpacksResponse {
	this := RpmPackageLangpacksResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RpmPackageLangpacksResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageLangpacksResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RpmPackageLangpacksResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RpmPackageLangpacksResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RpmPackageLangpacksResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackageLangpacksResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RpmPackageLangpacksResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RpmPackageLangpacksResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetMatches returns the Matches field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *RpmPackageLangpacksResponse) GetMatches() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageLangpacksResponse) GetMatchesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Matches) {
		return map[string]interface{}{}, false
	}
	return o.Matches, true
}

// SetMatches sets field value
func (o *RpmPackageLangpacksResponse) SetMatches(v map[string]interface{}) {
	o.Matches = v
}

// GetDigest returns the Digest field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RpmPackageLangpacksResponse) GetDigest() string {
	if o == nil || o.Digest.Get() == nil {
		var ret string
		return ret
	}

	return *o.Digest.Get()
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmPackageLangpacksResponse) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Digest.Get(), o.Digest.IsSet()
}

// SetDigest sets field value
func (o *RpmPackageLangpacksResponse) SetDigest(v string) {
	o.Digest.Set(&v)
}

func (o RpmPackageLangpacksResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmPackageLangpacksResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if o.Matches != nil {
		toSerialize["matches"] = o.Matches
	}
	toSerialize["digest"] = o.Digest.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RpmPackageLangpacksResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRpmPackageLangpacksResponse := _RpmPackageLangpacksResponse{}

	if err = json.Unmarshal(bytes, &varRpmPackageLangpacksResponse); err == nil {
		*o = RpmPackageLangpacksResponse(varRpmPackageLangpacksResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "matches")
		delete(additionalProperties, "digest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRpmPackageLangpacksResponse struct {
	value *RpmPackageLangpacksResponse
	isSet bool
}

func (v NullableRpmPackageLangpacksResponse) Get() *RpmPackageLangpacksResponse {
	return v.value
}

func (v *NullableRpmPackageLangpacksResponse) Set(val *RpmPackageLangpacksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmPackageLangpacksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmPackageLangpacksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmPackageLangpacksResponse(val *RpmPackageLangpacksResponse) *NullableRpmPackageLangpacksResponse {
	return &NullableRpmPackageLangpacksResponse{value: val, isSet: true}
}

func (v NullableRpmPackageLangpacksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmPackageLangpacksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

