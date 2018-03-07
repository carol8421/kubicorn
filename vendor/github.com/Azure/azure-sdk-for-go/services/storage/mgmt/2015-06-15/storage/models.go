package storage

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"net/http"
)

// AccountStatus enumerates the values for account status.
type AccountStatus string

const (
	// Available ...
	Available AccountStatus = "Available"
	// Unavailable ...
	Unavailable AccountStatus = "Unavailable"
)

// AccountType enumerates the values for account type.
type AccountType string

const (
	// PremiumLRS ...
	PremiumLRS AccountType = "Premium_LRS"
	// StandardGRS ...
	StandardGRS AccountType = "Standard_GRS"
	// StandardLRS ...
	StandardLRS AccountType = "Standard_LRS"
	// StandardRAGRS ...
	StandardRAGRS AccountType = "Standard_RAGRS"
	// StandardZRS ...
	StandardZRS AccountType = "Standard_ZRS"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// ResolvingDNS ...
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// Reason enumerates the values for reason.
type Reason string

const (
	// AccountNameInvalid ...
	AccountNameInvalid Reason = "AccountNameInvalid"
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
)

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// Bytes ...
	Bytes UsageUnit = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UsageUnit = "BytesPerSecond"
	// Count ...
	Count UsageUnit = "Count"
	// CountsPerSecond ...
	CountsPerSecond UsageUnit = "CountsPerSecond"
	// Percent ...
	Percent UsageUnit = "Percent"
	// Seconds ...
	Seconds UsageUnit = "Seconds"
)

// Account the storage account.
type Account struct {
	autorest.Response `json:"-"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags               *map[string]*string `json:"tags,omitempty"`
	*AccountProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for Account struct.
func (a *Account) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties AccountProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		a.AccountProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		a.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		a.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		a.Type = &typeVar
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		a.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		a.Tags = &tags
	}

	return nil
}

// AccountCheckNameAvailabilityParameters the parameters used to check the availabity of the storage account name.
type AccountCheckNameAvailabilityParameters struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// AccountCreateParameters the parameters to provide for the account.
type AccountCreateParameters struct {
	// Location - The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
	Location *string `json:"location,omitempty"`
	// Tags - A list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*AccountPropertiesCreateParameters `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for AccountCreateParameters struct.
func (acp *AccountCreateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		acp.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		acp.Tags = &tags
	}

	v = m["properties"]
	if v != nil {
		var properties AccountPropertiesCreateParameters
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		acp.AccountPropertiesCreateParameters = &properties
	}

	return nil
}

// AccountKeys the access keys for the storage account.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// Key1 - The value of key 1.
	Key1 *string `json:"key1,omitempty"`
	// Key2 - The value of key 2.
	Key2 *string `json:"key2,omitempty"`
}

// AccountListResult the list storage accounts operation response.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of storage accounts and their properties.
	Value *[]Account `json:"value,omitempty"`
}

// AccountProperties properties of the storage account.
type AccountProperties struct {
	// ProvisioningState - The status of the storage account at the time the operation was called. Possible values include: 'Creating', 'ResolvingDNS', 'Succeeded'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// AccountType - The type of the storage account. Possible values include: 'StandardLRS', 'StandardZRS', 'StandardGRS', 'StandardRAGRS', 'PremiumLRS'
	AccountType AccountType `json:"accountType,omitempty"`
	// PrimaryEndpoints - The URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints *Endpoints `json:"primaryEndpoints,omitempty"`
	// PrimaryLocation - The location of the primary data center for the storage account.
	PrimaryLocation *string `json:"primaryLocation,omitempty"`
	// StatusOfPrimary - The status indicating whether the primary location of the storage account is available or unavailable. Possible values include: 'Available', 'Unavailable'
	StatusOfPrimary AccountStatus `json:"statusOfPrimary,omitempty"`
	// LastGeoFailoverTime - The timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime *date.Time `json:"lastGeoFailoverTime,omitempty"`
	// SecondaryLocation - The location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`
	// StatusOfSecondary - The status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS. Possible values include: 'Available', 'Unavailable'
	StatusOfSecondary AccountStatus `json:"statusOfSecondary,omitempty"`
	// CreationTime - The creation date and time of the storage account in UTC.
	CreationTime *date.Time `json:"creationTime,omitempty"`
	// CustomDomain - The custom domain the user assigned to this storage account.
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
	// SecondaryEndpoints - The URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints *Endpoints `json:"secondaryEndpoints,omitempty"`
}

// AccountPropertiesCreateParameters the parameters used to create the storage account.
type AccountPropertiesCreateParameters struct {
	// AccountType - The sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType. Possible values include: 'StandardLRS', 'StandardZRS', 'StandardGRS', 'StandardRAGRS', 'PremiumLRS'
	AccountType AccountType `json:"accountType,omitempty"`
}

// AccountPropertiesUpdateParameters the parameters used when updating a storage account.
type AccountPropertiesUpdateParameters struct {
	// AccountType - The account type. Note that StandardZRS and PremiumLRS accounts cannot be changed to other account types, and other account types cannot be changed to StandardZRS or PremiumLRS. Possible values include: 'StandardLRS', 'StandardZRS', 'StandardGRS', 'StandardRAGRS', 'PremiumLRS'
	AccountType AccountType `json:"accountType,omitempty"`
	// CustomDomain - User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
}

// AccountRegenerateKeyParameters the parameters used to regenerate the storage account key.
type AccountRegenerateKeyParameters struct {
	KeyName *string `json:"keyName,omitempty"`
}

// AccountsCreateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type AccountsCreateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future AccountsCreateFuture) Result(client AccountsClient) (a Account, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return a, autorest.NewError("storage.AccountsCreateFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		a, err = client.CreateResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	a, err = client.CreateResponder(resp)
	return
}

// AccountUpdateParameters the parameters to update on the account.
type AccountUpdateParameters struct {
	// Tags - Resource tags
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*AccountPropertiesUpdateParameters `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for AccountUpdateParameters struct.
func (aup *AccountUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		aup.Tags = &tags
	}

	v = m["properties"]
	if v != nil {
		var properties AccountPropertiesUpdateParameters
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		aup.AccountPropertiesUpdateParameters = &properties
	}

	return nil
}

// CheckNameAvailabilityResult the CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - Boolean value that indicates whether the name is available for you to use. If true, the name is available. If false, the name has already been taken or is invalid and cannot be used.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - The reason that a storage account name could not be used. The Reason element is only returned if NameAvailable is false. Possible values include: 'AccountNameInvalid', 'AlreadyExists'
	Reason Reason `json:"reason,omitempty"`
	// Message - The error message explaining the Reason value in more detail.
	Message *string `json:"message,omitempty"`
}

// CustomDomain the custom domain assigned to this storage account. This can be set via Update.
type CustomDomain struct {
	// Name - The custom domain name. Name is the CNAME source.
	Name *string `json:"name,omitempty"`
	// UseSubDomain - Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates
	UseSubDomain *bool `json:"useSubDomain,omitempty"`
}

// Endpoints the URIs that are used to perform a retrieval of a public blob, queue or table object.
type Endpoints struct {
	// Blob - The blob endpoint.
	Blob *string `json:"blob,omitempty"`
	// Queue - The queue endpoint.
	Queue *string `json:"queue,omitempty"`
	// Table - The table endpoint.
	Table *string `json:"table,omitempty"`
	// File - The file endpoint.
	File *string `json:"file,omitempty"`
}

// Resource describes a storage resource.
type Resource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags *map[string]*string `json:"tags,omitempty"`
}

// Usage describes Storage Resource Usage.
type Usage struct {
	// Unit - The unit of measurement. Possible values include: 'Count', 'Bytes', 'Seconds', 'Percent', 'CountsPerSecond', 'BytesPerSecond'
	Unit UsageUnit `json:"unit,omitempty"`
	// CurrentValue - The current count of the allocated resources in the subscription.
	CurrentValue *int32 `json:"currentValue,omitempty"`
	// Limit - The maximum count of the resources that can be allocated in the subscription.
	Limit *int32 `json:"limit,omitempty"`
	// Name - The name of the type of usage.
	Name *UsageName `json:"name,omitempty"`
}

// UsageListResult the List Usages operation response.
type UsageListResult struct {
	autorest.Response `json:"-"`
	// Value - The list Storage Resource Usages.
	Value *[]Usage `json:"value,omitempty"`
}

// UsageName the Usage Names.
type UsageName struct {
	// Value - A string describing the resource name.
	Value *string `json:"value,omitempty"`
	// LocalizedValue - A localized string describing the resource name.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}
