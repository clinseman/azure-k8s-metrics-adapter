package servicebus

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

// MigrationConfigsClient is the azure Service Bus client
type MigrationConfigsClient struct {
	BaseClient
}

// NewMigrationConfigsClient creates an instance of the MigrationConfigsClient client.
func NewMigrationConfigsClient(subscriptionID string) MigrationConfigsClient {
	return NewMigrationConfigsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMigrationConfigsClientWithBaseURI creates an instance of the MigrationConfigsClient client.
func NewMigrationConfigsClientWithBaseURI(baseURI string, subscriptionID string) MigrationConfigsClient {
	return MigrationConfigsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CompleteMigration this operation Completes Migration of entities by pointing the connection strings to Premium
// namespace and any enties created after the operation will be under Premium Namespace. CompleteMigration operation
// will fail when entity migration is in-progress.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
func (client MigrationConfigsClient) CompleteMigration(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.MigrationConfigsClient", "CompleteMigration", err.Error())
	}

	req, err := client.CompleteMigrationPreparer(ctx, resourceGroupName, namespaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "CompleteMigration", nil, "Failure preparing request")
		return
	}

	resp, err := client.CompleteMigrationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "CompleteMigration", resp, "Failure sending request")
		return
	}

	result, err = client.CompleteMigrationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "CompleteMigration", resp, "Failure responding to request")
	}

	return
}

// CompleteMigrationPreparer prepares the CompleteMigration request.
func (client MigrationConfigsClient) CompleteMigrationPreparer(ctx context.Context, resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configName":        autorest.Encode("path", "$default"),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}/upgrade", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CompleteMigrationSender sends the CompleteMigration request. The method will close the
// http.Response Body if it receives an error.
func (client MigrationConfigsClient) CompleteMigrationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CompleteMigrationResponder handles the response to the CompleteMigration request. The method always
// closes the http.Response Body.
func (client MigrationConfigsClient) CompleteMigrationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateAndStartMigration creates Migration configuration and starts migration of enties from Standard to Premium
// namespace
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
// parameters - parameters required to create Migration Configuration
func (client MigrationConfigsClient) CreateAndStartMigration(ctx context.Context, resourceGroupName string, namespaceName string, parameters MigrationConfigProperties) (result MigrationConfigsCreateAndStartMigrationFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.MigrationConfigPropertiesProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.MigrationConfigPropertiesProperties.TargetNamespace", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.MigrationConfigPropertiesProperties.PostMigrationName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("servicebus.MigrationConfigsClient", "CreateAndStartMigration", err.Error())
	}

	req, err := client.CreateAndStartMigrationPreparer(ctx, resourceGroupName, namespaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "CreateAndStartMigration", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateAndStartMigrationSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "CreateAndStartMigration", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateAndStartMigrationPreparer prepares the CreateAndStartMigration request.
func (client MigrationConfigsClient) CreateAndStartMigrationPreparer(ctx context.Context, resourceGroupName string, namespaceName string, parameters MigrationConfigProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configName":        autorest.Encode("path", "$default"),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateAndStartMigrationSender sends the CreateAndStartMigration request. The method will close the
// http.Response Body if it receives an error.
func (client MigrationConfigsClient) CreateAndStartMigrationSender(req *http.Request) (future MigrationConfigsCreateAndStartMigrationFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateAndStartMigrationResponder handles the response to the CreateAndStartMigration request. The method always
// closes the http.Response Body.
func (client MigrationConfigsClient) CreateAndStartMigrationResponder(resp *http.Response) (result MigrationConfigProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a MigrationConfiguration
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
func (client MigrationConfigsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.MigrationConfigsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, namespaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MigrationConfigsClient) DeletePreparer(ctx context.Context, resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configName":        autorest.Encode("path", "$default"),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MigrationConfigsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MigrationConfigsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves Migration Config
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
func (client MigrationConfigsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string) (result MigrationConfigProperties, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.MigrationConfigsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, namespaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MigrationConfigsClient) GetPreparer(ctx context.Context, resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configName":        autorest.Encode("path", "$default"),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MigrationConfigsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MigrationConfigsClient) GetResponder(resp *http.Response) (result MigrationConfigProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all migrationConfigurations
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
func (client MigrationConfigsClient) List(ctx context.Context, resourceGroupName string, namespaceName string) (result MigrationConfigListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.MigrationConfigsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, namespaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "List", resp, "Failure sending request")
		return
	}

	result.mclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MigrationConfigsClient) ListPreparer(ctx context.Context, resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MigrationConfigsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MigrationConfigsClient) ListResponder(resp *http.Response) (result MigrationConfigListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MigrationConfigsClient) listNextResults(lastResults MigrationConfigListResult) (result MigrationConfigListResult, err error) {
	req, err := lastResults.migrationConfigListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MigrationConfigsClient) ListComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result MigrationConfigListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, namespaceName)
	return
}

// Revert this operation reverts Migration
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// namespaceName - the namespace name
func (client MigrationConfigsClient) Revert(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicebus.MigrationConfigsClient", "Revert", err.Error())
	}

	req, err := client.RevertPreparer(ctx, resourceGroupName, namespaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Revert", nil, "Failure preparing request")
		return
	}

	resp, err := client.RevertSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Revert", resp, "Failure sending request")
		return
	}

	result, err = client.RevertResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.MigrationConfigsClient", "Revert", resp, "Failure responding to request")
	}

	return
}

// RevertPreparer prepares the Revert request.
func (client MigrationConfigsClient) RevertPreparer(ctx context.Context, resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configName":        autorest.Encode("path", "$default"),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}/revert", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RevertSender sends the Revert request. The method will close the
// http.Response Body if it receives an error.
func (client MigrationConfigsClient) RevertSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// RevertResponder handles the response to the Revert request. The method always
// closes the http.Response Body.
func (client MigrationConfigsClient) RevertResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
