package notificationhubs

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// NameClient is the azure NotificationHub client
type NameClient struct {
	BaseClient
}

// NewNameClient creates an instance of the NameClient client.
func NewNameClient(subscriptionID string) NameClient {
	return NewNameClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNameClientWithBaseURI creates an instance of the NameClient client.
func NewNameClientWithBaseURI(baseURI string, subscriptionID string) NameClient {
	return NameClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckAvailability checks the availability of the given service namespace across all Azure subscriptions. This is
// useful because the domain name is created based on the service namespace name.
//
// parameters is the namespace name.
func (client NameClient) CheckAvailability(ctx context.Context, parameters CheckNameAvailabilityRequestParameters) (result CheckNameAvailabilityResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "notificationhubs.NameClient", "CheckAvailability")
	}

	req, err := client.CheckAvailabilityPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationhubs.NameClient", "CheckAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "notificationhubs.NameClient", "CheckAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notificationhubs.NameClient", "CheckAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckAvailabilityPreparer prepares the CheckAvailability request.
func (client NameClient) CheckAvailabilityPreparer(ctx context.Context, parameters CheckNameAvailabilityRequestParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.NotificationHubs/checkNameAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckAvailabilitySender sends the CheckAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client NameClient) CheckAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CheckAvailabilityResponder handles the response to the CheckAvailability request. The method always
// closes the http.Response Body.
func (client NameClient) CheckAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
