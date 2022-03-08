package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

func (csClient *CSClient) UpdateUserGroup(ctx context.Context, req *CreateUserGroupRequest) (*Group, error) {
	method := "PUT"
	url := fmt.Sprintf("%s/user/groups", csClient.config.URL)
	bodyAsBytes, err := marshalCreateUserGroupRequest(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyAsBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}
	httpResp, err := csClient.signV2AndDo(req.AuthToken, httpReq, bodyAsBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to %s to %s: %s", method, url, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			_ = fmt.Errorf("failed to Close response body  %s", err)
		}
	}(httpResp.Body)

	var readUserGroupResp []Group
	if err := csClient.unmarshalJSONBody(httpResp.Body, &readUserGroupResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response body sdjhskdhskdskdskdksdkskjd: %s", err)
	}
	return &readUserGroupResp[0], err
}
