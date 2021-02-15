package automation

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DscNodeConfigurationClient is the automation Client
type DscNodeConfigurationClient struct {
	BaseClient
}

// NewDscNodeConfigurationClient creates an instance of the DscNodeConfigurationClient client.
func NewDscNodeConfigurationClient(subscriptionID string) DscNodeConfigurationClient {
	return NewDscNodeConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDscNodeConfigurationClientWithBaseURI creates an instance of the DscNodeConfigurationClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDscNodeConfigurationClientWithBaseURI(baseURI string, subscriptionID string) DscNodeConfigurationClient {
	return DscNodeConfigurationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create the node configuration identified by node configuration name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeConfigurationName - the Dsc node configuration name.
// parameters - the create or update parameters for configuration.
func (client DscNodeConfigurationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, parameters DscNodeConfigurationCreateOrUpdateParameters) (result DscNodeConfigurationCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeConfigurationClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DscNodeConfigurationCreateOrUpdateParametersProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.DscNodeConfigurationCreateOrUpdateParametersProperties.Source", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.DscNodeConfigurationCreateOrUpdateParametersProperties.Source.Hash", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.DscNodeConfigurationCreateOrUpdateParametersProperties.Source.Hash.Algorithm", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.DscNodeConfigurationCreateOrUpdateParametersProperties.Source.Hash.Value", Name: validation.Null, Rule: true, Chain: nil},
						}},
					}},
					{Target: "parameters.DscNodeConfigurationCreateOrUpdateParametersProperties.Configuration", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("automation.DscNodeConfigurationClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, automationAccountName, nodeConfigurationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DscNodeConfigurationClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, parameters DscNodeConfigurationCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeConfigurationName": autorest.Encode("path", nodeConfigurationName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeConfigurationClient) CreateOrUpdateSender(req *http.Request) (future DscNodeConfigurationCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DscNodeConfigurationClient) (dnc DscNodeConfiguration, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("automation.DscNodeConfigurationCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		dnc.Response.Response, err = future.GetResult(sender)
		if dnc.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && dnc.Response.Response.StatusCode != http.StatusNoContent {
			dnc, err = client.CreateOrUpdateResponder(dnc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationCreateOrUpdateFuture", "Result", dnc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DscNodeConfigurationClient) CreateOrUpdateResponder(resp *http.Response) (result DscNodeConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the Dsc node configurations by node configuration.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeConfigurationName - the Dsc node configuration name.
func (client DscNodeConfigurationClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeConfigurationClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.DscNodeConfigurationClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, automationAccountName, nodeConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DscNodeConfigurationClient) DeletePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeConfigurationName": autorest.Encode("path", nodeConfigurationName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeConfigurationClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DscNodeConfigurationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the Dsc node configurations by node configuration.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeConfigurationName - the Dsc node configuration name.
func (client DscNodeConfigurationClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string) (result DscNodeConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeConfigurationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.DscNodeConfigurationClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, nodeConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DscNodeConfigurationClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeConfigurationName": autorest.Encode("path", nodeConfigurationName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DscNodeConfigurationClient) GetResponder(resp *http.Response) (result DscNodeConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of dsc node configurations.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// filter - the filter to apply on the operation.
// skip - the number of rows to skip.
// top - the number of rows to take.
// inlinecount - return total rows.
func (client DscNodeConfigurationClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, skip *int32, top *int32, inlinecount string) (result DscNodeConfigurationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeConfigurationClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.dnclr.Response.Response != nil {
				sc = result.dnclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.DscNodeConfigurationClient", "ListByAutomationAccount", err.Error())
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, resourceGroupName, automationAccountName, filter, skip, top, inlinecount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.dnclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.dnclr, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "ListByAutomationAccount", resp, "Failure responding to request")
		return
	}
	if result.dnclr.hasNextLink() && result.dnclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client DscNodeConfigurationClient) ListByAutomationAccountPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, skip *int32, top *int32, inlinecount string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(inlinecount) > 0 {
		queryParameters["$inlinecount"] = autorest.Encode("query", inlinecount)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeConfigurationClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client DscNodeConfigurationClient) ListByAutomationAccountResponder(resp *http.Response) (result DscNodeConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAutomationAccountNextResults retrieves the next set of results, if any.
func (client DscNodeConfigurationClient) listByAutomationAccountNextResults(ctx context.Context, lastResults DscNodeConfigurationListResult) (result DscNodeConfigurationListResult, err error) {
	req, err := lastResults.dscNodeConfigurationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client DscNodeConfigurationClient) ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, skip *int32, top *int32, inlinecount string) (result DscNodeConfigurationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeConfigurationClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAutomationAccount(ctx, resourceGroupName, automationAccountName, filter, skip, top, inlinecount)
	return
}
