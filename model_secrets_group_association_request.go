/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the SecretsGroupAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretsGroupAssociationRequest{}

// SecretsGroupAssociationRequest Serializer for `SecretsGroupAssociation` objects.
type SecretsGroupAssociationRequest struct {
	AccessType AccessTypeEnum `json:"access_type"`
	SecretType SecretTypeEnum `json:"secret_type"`
	SecretsGroup BulkWritableCableRequestStatus `json:"secrets_group"`
	Secret BulkWritableCableRequestStatus `json:"secret"`
	AdditionalProperties map[string]interface{}
}

type _SecretsGroupAssociationRequest SecretsGroupAssociationRequest

// NewSecretsGroupAssociationRequest instantiates a new SecretsGroupAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretsGroupAssociationRequest(accessType AccessTypeEnum, secretType SecretTypeEnum, secretsGroup BulkWritableCableRequestStatus, secret BulkWritableCableRequestStatus) *SecretsGroupAssociationRequest {
	this := SecretsGroupAssociationRequest{}
	this.AccessType = accessType
	this.SecretType = secretType
	this.SecretsGroup = secretsGroup
	this.Secret = secret
	return &this
}

// NewSecretsGroupAssociationRequestWithDefaults instantiates a new SecretsGroupAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretsGroupAssociationRequestWithDefaults() *SecretsGroupAssociationRequest {
	this := SecretsGroupAssociationRequest{}
	return &this
}

// GetAccessType returns the AccessType field value
func (o *SecretsGroupAssociationRequest) GetAccessType() AccessTypeEnum {
	if o == nil {
		var ret AccessTypeEnum
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociationRequest) GetAccessTypeOk() (*AccessTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *SecretsGroupAssociationRequest) SetAccessType(v AccessTypeEnum) {
	o.AccessType = v
}

// GetSecretType returns the SecretType field value
func (o *SecretsGroupAssociationRequest) GetSecretType() SecretTypeEnum {
	if o == nil {
		var ret SecretTypeEnum
		return ret
	}

	return o.SecretType
}

// GetSecretTypeOk returns a tuple with the SecretType field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociationRequest) GetSecretTypeOk() (*SecretTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretType, true
}

// SetSecretType sets field value
func (o *SecretsGroupAssociationRequest) SetSecretType(v SecretTypeEnum) {
	o.SecretType = v
}

// GetSecretsGroup returns the SecretsGroup field value
func (o *SecretsGroupAssociationRequest) GetSecretsGroup() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.SecretsGroup
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociationRequest) GetSecretsGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsGroup, true
}

// SetSecretsGroup sets field value
func (o *SecretsGroupAssociationRequest) SetSecretsGroup(v BulkWritableCableRequestStatus) {
	o.SecretsGroup = v
}

// GetSecret returns the Secret field value
func (o *SecretsGroupAssociationRequest) GetSecret() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociationRequest) GetSecretOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *SecretsGroupAssociationRequest) SetSecret(v BulkWritableCableRequestStatus) {
	o.Secret = v
}

func (o SecretsGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretsGroupAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_type"] = o.AccessType
	toSerialize["secret_type"] = o.SecretType
	toSerialize["secrets_group"] = o.SecretsGroup
	toSerialize["secret"] = o.Secret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecretsGroupAssociationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_type",
		"secret_type",
		"secrets_group",
		"secret",
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

	varSecretsGroupAssociationRequest := _SecretsGroupAssociationRequest{}

	err = json.Unmarshal(data, &varSecretsGroupAssociationRequest)

	if err != nil {
		return err
	}

	*o = SecretsGroupAssociationRequest(varSecretsGroupAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_type")
		delete(additionalProperties, "secret_type")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecretsGroupAssociationRequest struct {
	value *SecretsGroupAssociationRequest
	isSet bool
}

func (v NullableSecretsGroupAssociationRequest) Get() *SecretsGroupAssociationRequest {
	return v.value
}

func (v *NullableSecretsGroupAssociationRequest) Set(val *SecretsGroupAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretsGroupAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretsGroupAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretsGroupAssociationRequest(val *SecretsGroupAssociationRequest) *NullableSecretsGroupAssociationRequest {
	return &NullableSecretsGroupAssociationRequest{value: val, isSet: true}
}

func (v NullableSecretsGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretsGroupAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


