// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package qbusinessiface provides an interface to enable mocking the QBusiness service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package qbusinessiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/qbusiness"
)

// QBusinessAPI provides an interface to enable mocking the
// qbusiness.QBusiness service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// QBusiness.
//	func myFunc(svc qbusinessiface.QBusinessAPI) bool {
//	    // Make svc.BatchDeleteDocument request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := qbusiness.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockQBusinessClient struct {
//	    qbusinessiface.QBusinessAPI
//	}
//	func (m *mockQBusinessClient) BatchDeleteDocument(input *qbusiness.BatchDeleteDocumentInput) (*qbusiness.BatchDeleteDocumentOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockQBusinessClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type QBusinessAPI interface {
	BatchDeleteDocument(*qbusiness.BatchDeleteDocumentInput) (*qbusiness.BatchDeleteDocumentOutput, error)
	BatchDeleteDocumentWithContext(aws.Context, *qbusiness.BatchDeleteDocumentInput, ...request.Option) (*qbusiness.BatchDeleteDocumentOutput, error)
	BatchDeleteDocumentRequest(*qbusiness.BatchDeleteDocumentInput) (*request.Request, *qbusiness.BatchDeleteDocumentOutput)

	BatchPutDocument(*qbusiness.BatchPutDocumentInput) (*qbusiness.BatchPutDocumentOutput, error)
	BatchPutDocumentWithContext(aws.Context, *qbusiness.BatchPutDocumentInput, ...request.Option) (*qbusiness.BatchPutDocumentOutput, error)
	BatchPutDocumentRequest(*qbusiness.BatchPutDocumentInput) (*request.Request, *qbusiness.BatchPutDocumentOutput)

	ChatSync(*qbusiness.ChatSyncInput) (*qbusiness.ChatSyncOutput, error)
	ChatSyncWithContext(aws.Context, *qbusiness.ChatSyncInput, ...request.Option) (*qbusiness.ChatSyncOutput, error)
	ChatSyncRequest(*qbusiness.ChatSyncInput) (*request.Request, *qbusiness.ChatSyncOutput)

	CreateApplication(*qbusiness.CreateApplicationInput) (*qbusiness.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *qbusiness.CreateApplicationInput, ...request.Option) (*qbusiness.CreateApplicationOutput, error)
	CreateApplicationRequest(*qbusiness.CreateApplicationInput) (*request.Request, *qbusiness.CreateApplicationOutput)

	CreateIndex(*qbusiness.CreateIndexInput) (*qbusiness.CreateIndexOutput, error)
	CreateIndexWithContext(aws.Context, *qbusiness.CreateIndexInput, ...request.Option) (*qbusiness.CreateIndexOutput, error)
	CreateIndexRequest(*qbusiness.CreateIndexInput) (*request.Request, *qbusiness.CreateIndexOutput)

	CreatePlugin(*qbusiness.CreatePluginInput) (*qbusiness.CreatePluginOutput, error)
	CreatePluginWithContext(aws.Context, *qbusiness.CreatePluginInput, ...request.Option) (*qbusiness.CreatePluginOutput, error)
	CreatePluginRequest(*qbusiness.CreatePluginInput) (*request.Request, *qbusiness.CreatePluginOutput)

	CreateRetriever(*qbusiness.CreateRetrieverInput) (*qbusiness.CreateRetrieverOutput, error)
	CreateRetrieverWithContext(aws.Context, *qbusiness.CreateRetrieverInput, ...request.Option) (*qbusiness.CreateRetrieverOutput, error)
	CreateRetrieverRequest(*qbusiness.CreateRetrieverInput) (*request.Request, *qbusiness.CreateRetrieverOutput)

	CreateUser(*qbusiness.CreateUserInput) (*qbusiness.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *qbusiness.CreateUserInput, ...request.Option) (*qbusiness.CreateUserOutput, error)
	CreateUserRequest(*qbusiness.CreateUserInput) (*request.Request, *qbusiness.CreateUserOutput)

	CreateWebExperience(*qbusiness.CreateWebExperienceInput) (*qbusiness.CreateWebExperienceOutput, error)
	CreateWebExperienceWithContext(aws.Context, *qbusiness.CreateWebExperienceInput, ...request.Option) (*qbusiness.CreateWebExperienceOutput, error)
	CreateWebExperienceRequest(*qbusiness.CreateWebExperienceInput) (*request.Request, *qbusiness.CreateWebExperienceOutput)

	DeleteApplication(*qbusiness.DeleteApplicationInput) (*qbusiness.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *qbusiness.DeleteApplicationInput, ...request.Option) (*qbusiness.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*qbusiness.DeleteApplicationInput) (*request.Request, *qbusiness.DeleteApplicationOutput)

	DeleteChatControlsConfiguration(*qbusiness.DeleteChatControlsConfigurationInput) (*qbusiness.DeleteChatControlsConfigurationOutput, error)
	DeleteChatControlsConfigurationWithContext(aws.Context, *qbusiness.DeleteChatControlsConfigurationInput, ...request.Option) (*qbusiness.DeleteChatControlsConfigurationOutput, error)
	DeleteChatControlsConfigurationRequest(*qbusiness.DeleteChatControlsConfigurationInput) (*request.Request, *qbusiness.DeleteChatControlsConfigurationOutput)

	DeleteConversation(*qbusiness.DeleteConversationInput) (*qbusiness.DeleteConversationOutput, error)
	DeleteConversationWithContext(aws.Context, *qbusiness.DeleteConversationInput, ...request.Option) (*qbusiness.DeleteConversationOutput, error)
	DeleteConversationRequest(*qbusiness.DeleteConversationInput) (*request.Request, *qbusiness.DeleteConversationOutput)

	DeleteDataSource(*qbusiness.DeleteDataSourceInput) (*qbusiness.DeleteDataSourceOutput, error)
	DeleteDataSourceWithContext(aws.Context, *qbusiness.DeleteDataSourceInput, ...request.Option) (*qbusiness.DeleteDataSourceOutput, error)
	DeleteDataSourceRequest(*qbusiness.DeleteDataSourceInput) (*request.Request, *qbusiness.DeleteDataSourceOutput)

	DeleteGroup(*qbusiness.DeleteGroupInput) (*qbusiness.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *qbusiness.DeleteGroupInput, ...request.Option) (*qbusiness.DeleteGroupOutput, error)
	DeleteGroupRequest(*qbusiness.DeleteGroupInput) (*request.Request, *qbusiness.DeleteGroupOutput)

	DeleteIndex(*qbusiness.DeleteIndexInput) (*qbusiness.DeleteIndexOutput, error)
	DeleteIndexWithContext(aws.Context, *qbusiness.DeleteIndexInput, ...request.Option) (*qbusiness.DeleteIndexOutput, error)
	DeleteIndexRequest(*qbusiness.DeleteIndexInput) (*request.Request, *qbusiness.DeleteIndexOutput)

	DeletePlugin(*qbusiness.DeletePluginInput) (*qbusiness.DeletePluginOutput, error)
	DeletePluginWithContext(aws.Context, *qbusiness.DeletePluginInput, ...request.Option) (*qbusiness.DeletePluginOutput, error)
	DeletePluginRequest(*qbusiness.DeletePluginInput) (*request.Request, *qbusiness.DeletePluginOutput)

	DeleteRetriever(*qbusiness.DeleteRetrieverInput) (*qbusiness.DeleteRetrieverOutput, error)
	DeleteRetrieverWithContext(aws.Context, *qbusiness.DeleteRetrieverInput, ...request.Option) (*qbusiness.DeleteRetrieverOutput, error)
	DeleteRetrieverRequest(*qbusiness.DeleteRetrieverInput) (*request.Request, *qbusiness.DeleteRetrieverOutput)

	DeleteUser(*qbusiness.DeleteUserInput) (*qbusiness.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *qbusiness.DeleteUserInput, ...request.Option) (*qbusiness.DeleteUserOutput, error)
	DeleteUserRequest(*qbusiness.DeleteUserInput) (*request.Request, *qbusiness.DeleteUserOutput)

	DeleteWebExperience(*qbusiness.DeleteWebExperienceInput) (*qbusiness.DeleteWebExperienceOutput, error)
	DeleteWebExperienceWithContext(aws.Context, *qbusiness.DeleteWebExperienceInput, ...request.Option) (*qbusiness.DeleteWebExperienceOutput, error)
	DeleteWebExperienceRequest(*qbusiness.DeleteWebExperienceInput) (*request.Request, *qbusiness.DeleteWebExperienceOutput)

	GetApplication(*qbusiness.GetApplicationInput) (*qbusiness.GetApplicationOutput, error)
	GetApplicationWithContext(aws.Context, *qbusiness.GetApplicationInput, ...request.Option) (*qbusiness.GetApplicationOutput, error)
	GetApplicationRequest(*qbusiness.GetApplicationInput) (*request.Request, *qbusiness.GetApplicationOutput)

	GetChatControlsConfiguration(*qbusiness.GetChatControlsConfigurationInput) (*qbusiness.GetChatControlsConfigurationOutput, error)
	GetChatControlsConfigurationWithContext(aws.Context, *qbusiness.GetChatControlsConfigurationInput, ...request.Option) (*qbusiness.GetChatControlsConfigurationOutput, error)
	GetChatControlsConfigurationRequest(*qbusiness.GetChatControlsConfigurationInput) (*request.Request, *qbusiness.GetChatControlsConfigurationOutput)

	GetChatControlsConfigurationPages(*qbusiness.GetChatControlsConfigurationInput, func(*qbusiness.GetChatControlsConfigurationOutput, bool) bool) error
	GetChatControlsConfigurationPagesWithContext(aws.Context, *qbusiness.GetChatControlsConfigurationInput, func(*qbusiness.GetChatControlsConfigurationOutput, bool) bool, ...request.Option) error

	GetDataSource(*qbusiness.GetDataSourceInput) (*qbusiness.GetDataSourceOutput, error)
	GetDataSourceWithContext(aws.Context, *qbusiness.GetDataSourceInput, ...request.Option) (*qbusiness.GetDataSourceOutput, error)
	GetDataSourceRequest(*qbusiness.GetDataSourceInput) (*request.Request, *qbusiness.GetDataSourceOutput)

	GetGroup(*qbusiness.GetGroupInput) (*qbusiness.GetGroupOutput, error)
	GetGroupWithContext(aws.Context, *qbusiness.GetGroupInput, ...request.Option) (*qbusiness.GetGroupOutput, error)
	GetGroupRequest(*qbusiness.GetGroupInput) (*request.Request, *qbusiness.GetGroupOutput)

	GetIndex(*qbusiness.GetIndexInput) (*qbusiness.GetIndexOutput, error)
	GetIndexWithContext(aws.Context, *qbusiness.GetIndexInput, ...request.Option) (*qbusiness.GetIndexOutput, error)
	GetIndexRequest(*qbusiness.GetIndexInput) (*request.Request, *qbusiness.GetIndexOutput)

	GetPlugin(*qbusiness.GetPluginInput) (*qbusiness.GetPluginOutput, error)
	GetPluginWithContext(aws.Context, *qbusiness.GetPluginInput, ...request.Option) (*qbusiness.GetPluginOutput, error)
	GetPluginRequest(*qbusiness.GetPluginInput) (*request.Request, *qbusiness.GetPluginOutput)

	GetRetriever(*qbusiness.GetRetrieverInput) (*qbusiness.GetRetrieverOutput, error)
	GetRetrieverWithContext(aws.Context, *qbusiness.GetRetrieverInput, ...request.Option) (*qbusiness.GetRetrieverOutput, error)
	GetRetrieverRequest(*qbusiness.GetRetrieverInput) (*request.Request, *qbusiness.GetRetrieverOutput)

	GetUser(*qbusiness.GetUserInput) (*qbusiness.GetUserOutput, error)
	GetUserWithContext(aws.Context, *qbusiness.GetUserInput, ...request.Option) (*qbusiness.GetUserOutput, error)
	GetUserRequest(*qbusiness.GetUserInput) (*request.Request, *qbusiness.GetUserOutput)

	GetWebExperience(*qbusiness.GetWebExperienceInput) (*qbusiness.GetWebExperienceOutput, error)
	GetWebExperienceWithContext(aws.Context, *qbusiness.GetWebExperienceInput, ...request.Option) (*qbusiness.GetWebExperienceOutput, error)
	GetWebExperienceRequest(*qbusiness.GetWebExperienceInput) (*request.Request, *qbusiness.GetWebExperienceOutput)

	ListApplications(*qbusiness.ListApplicationsInput) (*qbusiness.ListApplicationsOutput, error)
	ListApplicationsWithContext(aws.Context, *qbusiness.ListApplicationsInput, ...request.Option) (*qbusiness.ListApplicationsOutput, error)
	ListApplicationsRequest(*qbusiness.ListApplicationsInput) (*request.Request, *qbusiness.ListApplicationsOutput)

	ListApplicationsPages(*qbusiness.ListApplicationsInput, func(*qbusiness.ListApplicationsOutput, bool) bool) error
	ListApplicationsPagesWithContext(aws.Context, *qbusiness.ListApplicationsInput, func(*qbusiness.ListApplicationsOutput, bool) bool, ...request.Option) error

	ListConversations(*qbusiness.ListConversationsInput) (*qbusiness.ListConversationsOutput, error)
	ListConversationsWithContext(aws.Context, *qbusiness.ListConversationsInput, ...request.Option) (*qbusiness.ListConversationsOutput, error)
	ListConversationsRequest(*qbusiness.ListConversationsInput) (*request.Request, *qbusiness.ListConversationsOutput)

	ListConversationsPages(*qbusiness.ListConversationsInput, func(*qbusiness.ListConversationsOutput, bool) bool) error
	ListConversationsPagesWithContext(aws.Context, *qbusiness.ListConversationsInput, func(*qbusiness.ListConversationsOutput, bool) bool, ...request.Option) error

	ListDataSourceSyncJobs(*qbusiness.ListDataSourceSyncJobsInput) (*qbusiness.ListDataSourceSyncJobsOutput, error)
	ListDataSourceSyncJobsWithContext(aws.Context, *qbusiness.ListDataSourceSyncJobsInput, ...request.Option) (*qbusiness.ListDataSourceSyncJobsOutput, error)
	ListDataSourceSyncJobsRequest(*qbusiness.ListDataSourceSyncJobsInput) (*request.Request, *qbusiness.ListDataSourceSyncJobsOutput)

	ListDataSourceSyncJobsPages(*qbusiness.ListDataSourceSyncJobsInput, func(*qbusiness.ListDataSourceSyncJobsOutput, bool) bool) error
	ListDataSourceSyncJobsPagesWithContext(aws.Context, *qbusiness.ListDataSourceSyncJobsInput, func(*qbusiness.ListDataSourceSyncJobsOutput, bool) bool, ...request.Option) error

	ListDataSources(*qbusiness.ListDataSourcesInput) (*qbusiness.ListDataSourcesOutput, error)
	ListDataSourcesWithContext(aws.Context, *qbusiness.ListDataSourcesInput, ...request.Option) (*qbusiness.ListDataSourcesOutput, error)
	ListDataSourcesRequest(*qbusiness.ListDataSourcesInput) (*request.Request, *qbusiness.ListDataSourcesOutput)

	ListDataSourcesPages(*qbusiness.ListDataSourcesInput, func(*qbusiness.ListDataSourcesOutput, bool) bool) error
	ListDataSourcesPagesWithContext(aws.Context, *qbusiness.ListDataSourcesInput, func(*qbusiness.ListDataSourcesOutput, bool) bool, ...request.Option) error

	ListDocuments(*qbusiness.ListDocumentsInput) (*qbusiness.ListDocumentsOutput, error)
	ListDocumentsWithContext(aws.Context, *qbusiness.ListDocumentsInput, ...request.Option) (*qbusiness.ListDocumentsOutput, error)
	ListDocumentsRequest(*qbusiness.ListDocumentsInput) (*request.Request, *qbusiness.ListDocumentsOutput)

	ListDocumentsPages(*qbusiness.ListDocumentsInput, func(*qbusiness.ListDocumentsOutput, bool) bool) error
	ListDocumentsPagesWithContext(aws.Context, *qbusiness.ListDocumentsInput, func(*qbusiness.ListDocumentsOutput, bool) bool, ...request.Option) error

	ListGroups(*qbusiness.ListGroupsInput) (*qbusiness.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *qbusiness.ListGroupsInput, ...request.Option) (*qbusiness.ListGroupsOutput, error)
	ListGroupsRequest(*qbusiness.ListGroupsInput) (*request.Request, *qbusiness.ListGroupsOutput)

	ListGroupsPages(*qbusiness.ListGroupsInput, func(*qbusiness.ListGroupsOutput, bool) bool) error
	ListGroupsPagesWithContext(aws.Context, *qbusiness.ListGroupsInput, func(*qbusiness.ListGroupsOutput, bool) bool, ...request.Option) error

	ListIndices(*qbusiness.ListIndicesInput) (*qbusiness.ListIndicesOutput, error)
	ListIndicesWithContext(aws.Context, *qbusiness.ListIndicesInput, ...request.Option) (*qbusiness.ListIndicesOutput, error)
	ListIndicesRequest(*qbusiness.ListIndicesInput) (*request.Request, *qbusiness.ListIndicesOutput)

	ListIndicesPages(*qbusiness.ListIndicesInput, func(*qbusiness.ListIndicesOutput, bool) bool) error
	ListIndicesPagesWithContext(aws.Context, *qbusiness.ListIndicesInput, func(*qbusiness.ListIndicesOutput, bool) bool, ...request.Option) error

	ListMessages(*qbusiness.ListMessagesInput) (*qbusiness.ListMessagesOutput, error)
	ListMessagesWithContext(aws.Context, *qbusiness.ListMessagesInput, ...request.Option) (*qbusiness.ListMessagesOutput, error)
	ListMessagesRequest(*qbusiness.ListMessagesInput) (*request.Request, *qbusiness.ListMessagesOutput)

	ListMessagesPages(*qbusiness.ListMessagesInput, func(*qbusiness.ListMessagesOutput, bool) bool) error
	ListMessagesPagesWithContext(aws.Context, *qbusiness.ListMessagesInput, func(*qbusiness.ListMessagesOutput, bool) bool, ...request.Option) error

	ListPlugins(*qbusiness.ListPluginsInput) (*qbusiness.ListPluginsOutput, error)
	ListPluginsWithContext(aws.Context, *qbusiness.ListPluginsInput, ...request.Option) (*qbusiness.ListPluginsOutput, error)
	ListPluginsRequest(*qbusiness.ListPluginsInput) (*request.Request, *qbusiness.ListPluginsOutput)

	ListPluginsPages(*qbusiness.ListPluginsInput, func(*qbusiness.ListPluginsOutput, bool) bool) error
	ListPluginsPagesWithContext(aws.Context, *qbusiness.ListPluginsInput, func(*qbusiness.ListPluginsOutput, bool) bool, ...request.Option) error

	ListRetrievers(*qbusiness.ListRetrieversInput) (*qbusiness.ListRetrieversOutput, error)
	ListRetrieversWithContext(aws.Context, *qbusiness.ListRetrieversInput, ...request.Option) (*qbusiness.ListRetrieversOutput, error)
	ListRetrieversRequest(*qbusiness.ListRetrieversInput) (*request.Request, *qbusiness.ListRetrieversOutput)

	ListRetrieversPages(*qbusiness.ListRetrieversInput, func(*qbusiness.ListRetrieversOutput, bool) bool) error
	ListRetrieversPagesWithContext(aws.Context, *qbusiness.ListRetrieversInput, func(*qbusiness.ListRetrieversOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*qbusiness.ListTagsForResourceInput) (*qbusiness.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *qbusiness.ListTagsForResourceInput, ...request.Option) (*qbusiness.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*qbusiness.ListTagsForResourceInput) (*request.Request, *qbusiness.ListTagsForResourceOutput)

	ListWebExperiences(*qbusiness.ListWebExperiencesInput) (*qbusiness.ListWebExperiencesOutput, error)
	ListWebExperiencesWithContext(aws.Context, *qbusiness.ListWebExperiencesInput, ...request.Option) (*qbusiness.ListWebExperiencesOutput, error)
	ListWebExperiencesRequest(*qbusiness.ListWebExperiencesInput) (*request.Request, *qbusiness.ListWebExperiencesOutput)

	ListWebExperiencesPages(*qbusiness.ListWebExperiencesInput, func(*qbusiness.ListWebExperiencesOutput, bool) bool) error
	ListWebExperiencesPagesWithContext(aws.Context, *qbusiness.ListWebExperiencesInput, func(*qbusiness.ListWebExperiencesOutput, bool) bool, ...request.Option) error

	PutFeedback(*qbusiness.PutFeedbackInput) (*qbusiness.PutFeedbackOutput, error)
	PutFeedbackWithContext(aws.Context, *qbusiness.PutFeedbackInput, ...request.Option) (*qbusiness.PutFeedbackOutput, error)
	PutFeedbackRequest(*qbusiness.PutFeedbackInput) (*request.Request, *qbusiness.PutFeedbackOutput)

	PutGroup(*qbusiness.PutGroupInput) (*qbusiness.PutGroupOutput, error)
	PutGroupWithContext(aws.Context, *qbusiness.PutGroupInput, ...request.Option) (*qbusiness.PutGroupOutput, error)
	PutGroupRequest(*qbusiness.PutGroupInput) (*request.Request, *qbusiness.PutGroupOutput)

	StartDataSourceSyncJob(*qbusiness.StartDataSourceSyncJobInput) (*qbusiness.StartDataSourceSyncJobOutput, error)
	StartDataSourceSyncJobWithContext(aws.Context, *qbusiness.StartDataSourceSyncJobInput, ...request.Option) (*qbusiness.StartDataSourceSyncJobOutput, error)
	StartDataSourceSyncJobRequest(*qbusiness.StartDataSourceSyncJobInput) (*request.Request, *qbusiness.StartDataSourceSyncJobOutput)

	StopDataSourceSyncJob(*qbusiness.StopDataSourceSyncJobInput) (*qbusiness.StopDataSourceSyncJobOutput, error)
	StopDataSourceSyncJobWithContext(aws.Context, *qbusiness.StopDataSourceSyncJobInput, ...request.Option) (*qbusiness.StopDataSourceSyncJobOutput, error)
	StopDataSourceSyncJobRequest(*qbusiness.StopDataSourceSyncJobInput) (*request.Request, *qbusiness.StopDataSourceSyncJobOutput)

	TagResource(*qbusiness.TagResourceInput) (*qbusiness.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *qbusiness.TagResourceInput, ...request.Option) (*qbusiness.TagResourceOutput, error)
	TagResourceRequest(*qbusiness.TagResourceInput) (*request.Request, *qbusiness.TagResourceOutput)

	UntagResource(*qbusiness.UntagResourceInput) (*qbusiness.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *qbusiness.UntagResourceInput, ...request.Option) (*qbusiness.UntagResourceOutput, error)
	UntagResourceRequest(*qbusiness.UntagResourceInput) (*request.Request, *qbusiness.UntagResourceOutput)

	UpdateApplication(*qbusiness.UpdateApplicationInput) (*qbusiness.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *qbusiness.UpdateApplicationInput, ...request.Option) (*qbusiness.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*qbusiness.UpdateApplicationInput) (*request.Request, *qbusiness.UpdateApplicationOutput)

	UpdateChatControlsConfiguration(*qbusiness.UpdateChatControlsConfigurationInput) (*qbusiness.UpdateChatControlsConfigurationOutput, error)
	UpdateChatControlsConfigurationWithContext(aws.Context, *qbusiness.UpdateChatControlsConfigurationInput, ...request.Option) (*qbusiness.UpdateChatControlsConfigurationOutput, error)
	UpdateChatControlsConfigurationRequest(*qbusiness.UpdateChatControlsConfigurationInput) (*request.Request, *qbusiness.UpdateChatControlsConfigurationOutput)

	UpdateDataSource(*qbusiness.UpdateDataSourceInput) (*qbusiness.UpdateDataSourceOutput, error)
	UpdateDataSourceWithContext(aws.Context, *qbusiness.UpdateDataSourceInput, ...request.Option) (*qbusiness.UpdateDataSourceOutput, error)
	UpdateDataSourceRequest(*qbusiness.UpdateDataSourceInput) (*request.Request, *qbusiness.UpdateDataSourceOutput)

	UpdateIndex(*qbusiness.UpdateIndexInput) (*qbusiness.UpdateIndexOutput, error)
	UpdateIndexWithContext(aws.Context, *qbusiness.UpdateIndexInput, ...request.Option) (*qbusiness.UpdateIndexOutput, error)
	UpdateIndexRequest(*qbusiness.UpdateIndexInput) (*request.Request, *qbusiness.UpdateIndexOutput)

	UpdatePlugin(*qbusiness.UpdatePluginInput) (*qbusiness.UpdatePluginOutput, error)
	UpdatePluginWithContext(aws.Context, *qbusiness.UpdatePluginInput, ...request.Option) (*qbusiness.UpdatePluginOutput, error)
	UpdatePluginRequest(*qbusiness.UpdatePluginInput) (*request.Request, *qbusiness.UpdatePluginOutput)

	UpdateRetriever(*qbusiness.UpdateRetrieverInput) (*qbusiness.UpdateRetrieverOutput, error)
	UpdateRetrieverWithContext(aws.Context, *qbusiness.UpdateRetrieverInput, ...request.Option) (*qbusiness.UpdateRetrieverOutput, error)
	UpdateRetrieverRequest(*qbusiness.UpdateRetrieverInput) (*request.Request, *qbusiness.UpdateRetrieverOutput)

	UpdateUser(*qbusiness.UpdateUserInput) (*qbusiness.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *qbusiness.UpdateUserInput, ...request.Option) (*qbusiness.UpdateUserOutput, error)
	UpdateUserRequest(*qbusiness.UpdateUserInput) (*request.Request, *qbusiness.UpdateUserOutput)

	UpdateWebExperience(*qbusiness.UpdateWebExperienceInput) (*qbusiness.UpdateWebExperienceOutput, error)
	UpdateWebExperienceWithContext(aws.Context, *qbusiness.UpdateWebExperienceInput, ...request.Option) (*qbusiness.UpdateWebExperienceOutput, error)
	UpdateWebExperienceRequest(*qbusiness.UpdateWebExperienceInput) (*request.Request, *qbusiness.UpdateWebExperienceOutput)
}

var _ QBusinessAPI = (*qbusiness.QBusiness)(nil)