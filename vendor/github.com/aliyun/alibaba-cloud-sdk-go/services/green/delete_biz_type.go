package green

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

// DeleteBizType invokes the green.DeleteBizType API synchronously
// api document: https://help.aliyun.com/api/green/deletebiztype.html
func (client *Client) DeleteBizType(request *DeleteBizTypeRequest) (response *DeleteBizTypeResponse, err error) {
	response = CreateDeleteBizTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBizTypeWithChan invokes the green.DeleteBizType API asynchronously
// api document: https://help.aliyun.com/api/green/deletebiztype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBizTypeWithChan(request *DeleteBizTypeRequest) (<-chan *DeleteBizTypeResponse, <-chan error) {
	responseChan := make(chan *DeleteBizTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBizType(request)
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

// DeleteBizTypeWithCallback invokes the green.DeleteBizType API asynchronously
// api document: https://help.aliyun.com/api/green/deletebiztype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBizTypeWithCallback(request *DeleteBizTypeRequest, callback func(response *DeleteBizTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBizTypeResponse
		var err error
		defer close(result)
		response, err = client.DeleteBizType(request)
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

// DeleteBizTypeRequest is the request struct for api DeleteBizType
type DeleteBizTypeRequest struct {
	*requests.RpcRequest
	SourceIp    string `position:"Query" name:"SourceIp"`
	BizTypeName string `position:"Query" name:"BizTypeName"`
}

// DeleteBizTypeResponse is the response struct for api DeleteBizType
type DeleteBizTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteBizTypeRequest creates a request to invoke DeleteBizType API
func CreateDeleteBizTypeRequest() (request *DeleteBizTypeRequest) {
	request = &DeleteBizTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DeleteBizType", "green", "openAPI")
	return
}

// CreateDeleteBizTypeResponse creates a response to parse from DeleteBizType response
func CreateDeleteBizTypeResponse() (response *DeleteBizTypeResponse) {
	response = &DeleteBizTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}