//
// MinIO Object Storage (c) 2021 MinIO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package madmin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// ClusterLinkResult - Result type for the cluster link command.
type ClusterLinkResult struct {
	Status                      string   `json:"status"`
	UsedBucketNames             []string `json:"usedBucketNames,omitempty"`
	UsedServiceAccountIDs       []string `json:"usedServiceAccountIDs,omitempty"`
	UsedSTSAccountIDs           []string `json:"usedSTSAccountIDs,omitempty"`
	UsedPolicyNames             []string `json:"usedPolicyNames,omitempty"`
	ExistingSTSPolicyMappings   []string `json:"existingSTSPolicyMappings,omitempty"`
	ExistingGroupPolicyMappings []string `json:"existingGroupPolicyMappings,omitempty"`
}

// ClusterLink sends the request to link with another cluster.
func (adm *AdminClient) ClusterLink(ctx context.Context) (ClusterLinkResult, error) {
	reqData := requestData{
		relPath: adminAPIPrefix + "/cluster/link",
	}

	var res ClusterLinkResult
	resp, err := adm.executeMethod(ctx, http.MethodPut, reqData)
	defer closeResponse(resp)
	if err != nil {
		return res, err
	}

	if resp.StatusCode != http.StatusOK {
		return res, httpRespToErrorResponse(resp)
	}

	data, err := DecryptData(adm.getSecretKey(), resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	return res, err
}

func (adm *AdminClient) ClusterLeave(ctx context.Context) error {
	return errors.New("Not implemented")
}

func (adm *AdminClient) ClusterInfo(ctx context.Context) error {
	return errors.New("Not implemented")
}
