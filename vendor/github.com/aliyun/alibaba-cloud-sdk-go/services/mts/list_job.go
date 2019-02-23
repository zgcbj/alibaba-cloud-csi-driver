package mts

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

// ListJob invokes the mts.ListJob API synchronously
// api document: https://help.aliyun.com/api/mts/listjob.html
func (client *Client) ListJob(request *ListJobRequest) (response *ListJobResponse, err error) {
	response = CreateListJobResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobWithChan invokes the mts.ListJob API asynchronously
// api document: https://help.aliyun.com/api/mts/listjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobWithChan(request *ListJobRequest) (<-chan *ListJobResponse, <-chan error) {
	responseChan := make(chan *ListJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJob(request)
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

// ListJobWithCallback invokes the mts.ListJob API asynchronously
// api document: https://help.aliyun.com/api/mts/listjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobWithCallback(request *ListJobRequest, callback func(response *ListJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobResponse
		var err error
		defer close(result)
		response, err = client.ListJob(request)
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

// ListJobRequest is the request struct for api ListJob
type ListJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	NextPageToken              string           `position:"Query" name:"NextPageToken"`
	StartOfJobCreatedTimeRange string           `position:"Query" name:"StartOfJobCreatedTimeRange"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	MaximumPageSize            requests.Integer `position:"Query" name:"MaximumPageSize"`
	State                      string           `position:"Query" name:"State"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	EndOfJobCreatedTimeRange   string           `position:"Query" name:"EndOfJobCreatedTimeRange"`
	PipelineId                 string           `position:"Query" name:"PipelineId"`
}

// ListJobResponse is the response struct for api ListJob
type ListJobResponse struct {
	*responses.BaseResponse
	RequestId     string           `json:"RequestId" xml:"RequestId"`
	NextPageToken string           `json:"NextPageToken" xml:"NextPageToken"`
	JobList       JobListInListJob `json:"JobList" xml:"JobList"`
}

// CreateListJobRequest creates a request to invoke ListJob API
func CreateListJobRequest() (request *ListJobRequest) {
	request = &ListJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListJob", "mts", "openAPI")
	return
}

// CreateListJobResponse creates a response to parse from ListJob response
func CreateListJobResponse() (response *ListJobResponse) {
	response = &ListJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}