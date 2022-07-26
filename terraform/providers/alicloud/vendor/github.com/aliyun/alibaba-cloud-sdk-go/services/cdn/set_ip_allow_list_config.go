package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SetIpAllowListConfig invokes the cdn.SetIpAllowListConfig API synchronously
func (client *Client) SetIpAllowListConfig(request *SetIpAllowListConfigRequest) (response *SetIpAllowListConfigResponse, err error) {
	response = CreateSetIpAllowListConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetIpAllowListConfigWithChan invokes the cdn.SetIpAllowListConfig API asynchronously
func (client *Client) SetIpAllowListConfigWithChan(request *SetIpAllowListConfigRequest) (<-chan *SetIpAllowListConfigResponse, <-chan error) {
	responseChan := make(chan *SetIpAllowListConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetIpAllowListConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SetIpAllowListConfigWithCallback invokes the cdn.SetIpAllowListConfig API asynchronously
func (client *Client) SetIpAllowListConfigWithCallback(request *SetIpAllowListConfigRequest, callback func(response *SetIpAllowListConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetIpAllowListConfigResponse
		var err error
		defer close(result)
		response, err = client.SetIpAllowListConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SetIpAllowListConfigRequest is the request struct for api SetIpAllowListConfig
type SetIpAllowListConfigRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	AllowIps      string           `position:"Query" name:"AllowIps"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// SetIpAllowListConfigResponse is the response struct for api SetIpAllowListConfig
type SetIpAllowListConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetIpAllowListConfigRequest creates a request to invoke SetIpAllowListConfig API
func CreateSetIpAllowListConfigRequest() (request *SetIpAllowListConfigRequest) {
	request = &SetIpAllowListConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetIpAllowListConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateSetIpAllowListConfigResponse creates a response to parse from SetIpAllowListConfig response
func CreateSetIpAllowListConfigResponse() (response *SetIpAllowListConfigResponse) {
	response = &SetIpAllowListConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}