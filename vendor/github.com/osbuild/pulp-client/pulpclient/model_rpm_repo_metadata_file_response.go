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

// checks if the RpmRepoMetadataFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmRepoMetadataFileResponse{}

// RpmRepoMetadataFileResponse RepoMetadataFile serializer.
type RpmRepoMetadataFileResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The MD5 checksum if available.
	Md5 *string `json:"md5,omitempty"`
	// The SHA-1 checksum if available.
	Sha1 *string `json:"sha1,omitempty"`
	// The SHA-224 checksum if available.
	Sha224 *string `json:"sha224,omitempty"`
	// The SHA-256 checksum if available.
	Sha256 *string `json:"sha256,omitempty"`
	// The SHA-384 checksum if available.
	Sha384 *string `json:"sha384,omitempty"`
	// The SHA-512 checksum if available.
	Sha512 *string `json:"sha512,omitempty"`
	// Artifact file representing the physical content
	Artifact *string `json:"artifact,omitempty"`
	// Relative path of the file.
	RelativePath string `json:"relative_path"`
	// Metadata type.
	DataType string `json:"data_type"`
	// Checksum type for the file.
	ChecksumType string `json:"checksum_type"`
	// Checksum for the file.
	Checksum string `json:"checksum"`
	AdditionalProperties map[string]interface{}
}

type _RpmRepoMetadataFileResponse RpmRepoMetadataFileResponse

// NewRpmRepoMetadataFileResponse instantiates a new RpmRepoMetadataFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmRepoMetadataFileResponse(relativePath string, dataType string, checksumType string, checksum string) *RpmRepoMetadataFileResponse {
	this := RpmRepoMetadataFileResponse{}
	this.RelativePath = relativePath
	this.DataType = dataType
	this.ChecksumType = checksumType
	this.Checksum = checksum
	return &this
}

// NewRpmRepoMetadataFileResponseWithDefaults instantiates a new RpmRepoMetadataFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmRepoMetadataFileResponseWithDefaults() *RpmRepoMetadataFileResponse {
	this := RpmRepoMetadataFileResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RpmRepoMetadataFileResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RpmRepoMetadataFileResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *RpmRepoMetadataFileResponse) SetMd5(v string) {
	o.Md5 = &v
}

// GetSha1 returns the Sha1 field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetSha1() string {
	if o == nil || IsNil(o.Sha1) {
		var ret string
		return ret
	}
	return *o.Sha1
}

// GetSha1Ok returns a tuple with the Sha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetSha1Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha1) {
		return nil, false
	}
	return o.Sha1, true
}

// HasSha1 returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasSha1() bool {
	if o != nil && !IsNil(o.Sha1) {
		return true
	}

	return false
}

// SetSha1 gets a reference to the given string and assigns it to the Sha1 field.
func (o *RpmRepoMetadataFileResponse) SetSha1(v string) {
	o.Sha1 = &v
}

// GetSha224 returns the Sha224 field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetSha224() string {
	if o == nil || IsNil(o.Sha224) {
		var ret string
		return ret
	}
	return *o.Sha224
}

// GetSha224Ok returns a tuple with the Sha224 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetSha224Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha224) {
		return nil, false
	}
	return o.Sha224, true
}

// HasSha224 returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasSha224() bool {
	if o != nil && !IsNil(o.Sha224) {
		return true
	}

	return false
}

// SetSha224 gets a reference to the given string and assigns it to the Sha224 field.
func (o *RpmRepoMetadataFileResponse) SetSha224(v string) {
	o.Sha224 = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetSha256() string {
	if o == nil || IsNil(o.Sha256) {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha256) {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasSha256() bool {
	if o != nil && !IsNil(o.Sha256) {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *RpmRepoMetadataFileResponse) SetSha256(v string) {
	o.Sha256 = &v
}

// GetSha384 returns the Sha384 field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetSha384() string {
	if o == nil || IsNil(o.Sha384) {
		var ret string
		return ret
	}
	return *o.Sha384
}

// GetSha384Ok returns a tuple with the Sha384 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetSha384Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha384) {
		return nil, false
	}
	return o.Sha384, true
}

// HasSha384 returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasSha384() bool {
	if o != nil && !IsNil(o.Sha384) {
		return true
	}

	return false
}

// SetSha384 gets a reference to the given string and assigns it to the Sha384 field.
func (o *RpmRepoMetadataFileResponse) SetSha384(v string) {
	o.Sha384 = &v
}

// GetSha512 returns the Sha512 field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetSha512() string {
	if o == nil || IsNil(o.Sha512) {
		var ret string
		return ret
	}
	return *o.Sha512
}

// GetSha512Ok returns a tuple with the Sha512 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetSha512Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha512) {
		return nil, false
	}
	return o.Sha512, true
}

// HasSha512 returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasSha512() bool {
	if o != nil && !IsNil(o.Sha512) {
		return true
	}

	return false
}

// SetSha512 gets a reference to the given string and assigns it to the Sha512 field.
func (o *RpmRepoMetadataFileResponse) SetSha512(v string) {
	o.Sha512 = &v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *RpmRepoMetadataFileResponse) GetArtifact() string {
	if o == nil || IsNil(o.Artifact) {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetArtifactOk() (*string, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *RpmRepoMetadataFileResponse) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *RpmRepoMetadataFileResponse) SetArtifact(v string) {
	o.Artifact = &v
}

// GetRelativePath returns the RelativePath field value
func (o *RpmRepoMetadataFileResponse) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *RpmRepoMetadataFileResponse) SetRelativePath(v string) {
	o.RelativePath = v
}

// GetDataType returns the DataType field value
func (o *RpmRepoMetadataFileResponse) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *RpmRepoMetadataFileResponse) SetDataType(v string) {
	o.DataType = v
}

// GetChecksumType returns the ChecksumType field value
func (o *RpmRepoMetadataFileResponse) GetChecksumType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChecksumType
}

// GetChecksumTypeOk returns a tuple with the ChecksumType field value
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetChecksumTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChecksumType, true
}

// SetChecksumType sets field value
func (o *RpmRepoMetadataFileResponse) SetChecksumType(v string) {
	o.ChecksumType = v
}

// GetChecksum returns the Checksum field value
func (o *RpmRepoMetadataFileResponse) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *RpmRepoMetadataFileResponse) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value
func (o *RpmRepoMetadataFileResponse) SetChecksum(v string) {
	o.Checksum = v
}

func (o RpmRepoMetadataFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmRepoMetadataFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !IsNil(o.Sha1) {
		toSerialize["sha1"] = o.Sha1
	}
	if !IsNil(o.Sha224) {
		toSerialize["sha224"] = o.Sha224
	}
	if !IsNil(o.Sha256) {
		toSerialize["sha256"] = o.Sha256
	}
	if !IsNil(o.Sha384) {
		toSerialize["sha384"] = o.Sha384
	}
	if !IsNil(o.Sha512) {
		toSerialize["sha512"] = o.Sha512
	}
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	toSerialize["relative_path"] = o.RelativePath
	toSerialize["data_type"] = o.DataType
	toSerialize["checksum_type"] = o.ChecksumType
	toSerialize["checksum"] = o.Checksum

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RpmRepoMetadataFileResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRpmRepoMetadataFileResponse := _RpmRepoMetadataFileResponse{}

	if err = json.Unmarshal(bytes, &varRpmRepoMetadataFileResponse); err == nil {
		*o = RpmRepoMetadataFileResponse(varRpmRepoMetadataFileResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "md5")
		delete(additionalProperties, "sha1")
		delete(additionalProperties, "sha224")
		delete(additionalProperties, "sha256")
		delete(additionalProperties, "sha384")
		delete(additionalProperties, "sha512")
		delete(additionalProperties, "artifact")
		delete(additionalProperties, "relative_path")
		delete(additionalProperties, "data_type")
		delete(additionalProperties, "checksum_type")
		delete(additionalProperties, "checksum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRpmRepoMetadataFileResponse struct {
	value *RpmRepoMetadataFileResponse
	isSet bool
}

func (v NullableRpmRepoMetadataFileResponse) Get() *RpmRepoMetadataFileResponse {
	return v.value
}

func (v *NullableRpmRepoMetadataFileResponse) Set(val *RpmRepoMetadataFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmRepoMetadataFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmRepoMetadataFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmRepoMetadataFileResponse(val *RpmRepoMetadataFileResponse) *NullableRpmRepoMetadataFileResponse {
	return &NullableRpmRepoMetadataFileResponse{value: val, isSet: true}
}

func (v NullableRpmRepoMetadataFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmRepoMetadataFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

