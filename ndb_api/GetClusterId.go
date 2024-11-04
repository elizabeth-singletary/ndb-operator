package ndb_api

import (
	"context"
	"fmt"

	"github.com/nutanix-cloud-native/ndb-operator/ndb_client"
)

func GetClusterByName(ctx context.Context, ndbClient ndb_client.NDBClientHTTPInterface, clusterName string) *ClusterResponse {

	url := fmt.Sprintf("clusters/name/%s", clusterName)
	clusterResp := &ClusterResponse{}

	_, err := sendRequest(ctx, ndbClient, "GET", url, nil, clusterResp)
	if err != nil {
		return nil
	}

	return clusterResp
}
