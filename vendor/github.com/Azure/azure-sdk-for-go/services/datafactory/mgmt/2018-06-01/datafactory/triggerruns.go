package datafactory

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

// TriggerRunsClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that
// interact with Azure Data Factory V2 services.
type TriggerRunsClient struct {
	BaseClient
}

// NewTriggerRunsClient creates an instance of the TriggerRunsClient client.
func NewTriggerRunsClient(subscriptionID string) TriggerRunsClient {
	return NewTriggerRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTriggerRunsClientWithBaseURI creates an instance of the TriggerRunsClient client.
func NewTriggerRunsClientWithBaseURI(baseURI string, subscriptionID string) TriggerRunsClient {
	return TriggerRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// QueryByFactory query trigger runs.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// filterParameters - parameters to filter the pipeline run.
func (client TriggerRunsClient) QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters) (result TriggerRunsQueryResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: filterParameters,
			Constraints: []validation.Constraint{{Target: "filterParameters.LastUpdatedAfter", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "filterParameters.LastUpdatedBefore", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.TriggerRunsClient", "QueryByFactory", err.Error())
	}

	req, err := client.QueryByFactoryPreparer(ctx, resourceGroupName, factoryName, filterParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "QueryByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "QueryByFactory", resp, "Failure sending request")
		return
	}

	result, err = client.QueryByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.TriggerRunsClient", "QueryByFactory", resp, "Failure responding to request")
	}

	return
}

// QueryByFactoryPreparer prepares the QueryByFactory request.
func (client TriggerRunsClient) QueryByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/queryTriggerRuns", pathParameters),
		autorest.WithJSON(filterParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryByFactorySender sends the QueryByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerRunsClient) QueryByFactorySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// QueryByFactoryResponder handles the response to the QueryByFactory request. The method always
// closes the http.Response Body.
func (client TriggerRunsClient) QueryByFactoryResponder(resp *http.Response) (result TriggerRunsQueryResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}