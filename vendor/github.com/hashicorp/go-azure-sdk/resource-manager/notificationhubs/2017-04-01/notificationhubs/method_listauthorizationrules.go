package notificationhubs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAuthorizationRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SharedAccessAuthorizationRuleResource
}

type ListAuthorizationRulesCompleteResult struct {
	Items []SharedAccessAuthorizationRuleResource
}

// ListAuthorizationRules ...
func (c NotificationHubsClient) ListAuthorizationRules(ctx context.Context, id NotificationHubId) (result ListAuthorizationRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/authorizationRules", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]SharedAccessAuthorizationRuleResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthorizationRulesComplete retrieves all the results into a single object
func (c NotificationHubsClient) ListAuthorizationRulesComplete(ctx context.Context, id NotificationHubId) (ListAuthorizationRulesCompleteResult, error) {
	return c.ListAuthorizationRulesCompleteMatchingPredicate(ctx, id, SharedAccessAuthorizationRuleResourceOperationPredicate{})
}

// ListAuthorizationRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NotificationHubsClient) ListAuthorizationRulesCompleteMatchingPredicate(ctx context.Context, id NotificationHubId, predicate SharedAccessAuthorizationRuleResourceOperationPredicate) (result ListAuthorizationRulesCompleteResult, err error) {
	items := make([]SharedAccessAuthorizationRuleResource, 0)

	resp, err := c.ListAuthorizationRules(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListAuthorizationRulesCompleteResult{
		Items: items,
	}
	return
}
