package connectionmonitors

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionMonitorsClient struct {
	Client *resourcemanager.Client
}

func NewConnectionMonitorsClientWithBaseURI(api environments.Api) (*ConnectionMonitorsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "connectionmonitors", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectionMonitorsClient: %+v", err)
	}

	return &ConnectionMonitorsClient{
		Client: client,
	}, nil
}
