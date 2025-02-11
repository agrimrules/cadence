// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by thriftrw-plugin-yarpc
// @generated

package historyserviceclient

import (
	context "context"
	reflect "reflect"

	wire "go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"

	history "github.com/uber/cadence/.gen/go/history"
	replicator "github.com/uber/cadence/.gen/go/replicator"
	shared "github.com/uber/cadence/.gen/go/shared"
)

// Interface is a client for the HistoryService service.
type Interface interface {
	CloseShard(
		ctx context.Context,
		Request *shared.CloseShardRequest,
		opts ...yarpc.CallOption,
	) error

	DescribeHistoryHost(
		ctx context.Context,
		Request *shared.DescribeHistoryHostRequest,
		opts ...yarpc.CallOption,
	) (*shared.DescribeHistoryHostResponse, error)

	DescribeMutableState(
		ctx context.Context,
		Request *history.DescribeMutableStateRequest,
		opts ...yarpc.CallOption,
	) (*history.DescribeMutableStateResponse, error)

	DescribeQueue(
		ctx context.Context,
		Request *shared.DescribeQueueRequest,
		opts ...yarpc.CallOption,
	) (*shared.DescribeQueueResponse, error)

	DescribeWorkflowExecution(
		ctx context.Context,
		DescribeRequest *history.DescribeWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) (*shared.DescribeWorkflowExecutionResponse, error)

	GetCrossClusterTasks(
		ctx context.Context,
		Request *shared.GetCrossClusterTasksRequest,
		opts ...yarpc.CallOption,
	) (*shared.GetCrossClusterTasksResponse, error)

	GetDLQReplicationMessages(
		ctx context.Context,
		Request *replicator.GetDLQReplicationMessagesRequest,
		opts ...yarpc.CallOption,
	) (*replicator.GetDLQReplicationMessagesResponse, error)

	GetFailoverInfo(
		ctx context.Context,
		Request *history.GetFailoverInfoRequest,
		opts ...yarpc.CallOption,
	) (*history.GetFailoverInfoResponse, error)

	GetMutableState(
		ctx context.Context,
		GetRequest *history.GetMutableStateRequest,
		opts ...yarpc.CallOption,
	) (*history.GetMutableStateResponse, error)

	GetReplicationMessages(
		ctx context.Context,
		Request *replicator.GetReplicationMessagesRequest,
		opts ...yarpc.CallOption,
	) (*replicator.GetReplicationMessagesResponse, error)

	MergeDLQMessages(
		ctx context.Context,
		Request *replicator.MergeDLQMessagesRequest,
		opts ...yarpc.CallOption,
	) (*replicator.MergeDLQMessagesResponse, error)

	NotifyFailoverMarkers(
		ctx context.Context,
		Request *history.NotifyFailoverMarkersRequest,
		opts ...yarpc.CallOption,
	) error

	PollMutableState(
		ctx context.Context,
		PollRequest *history.PollMutableStateRequest,
		opts ...yarpc.CallOption,
	) (*history.PollMutableStateResponse, error)

	PurgeDLQMessages(
		ctx context.Context,
		Request *replicator.PurgeDLQMessagesRequest,
		opts ...yarpc.CallOption,
	) error

	QueryWorkflow(
		ctx context.Context,
		QueryRequest *history.QueryWorkflowRequest,
		opts ...yarpc.CallOption,
	) (*history.QueryWorkflowResponse, error)

	ReadDLQMessages(
		ctx context.Context,
		Request *replicator.ReadDLQMessagesRequest,
		opts ...yarpc.CallOption,
	) (*replicator.ReadDLQMessagesResponse, error)

	ReapplyEvents(
		ctx context.Context,
		ReapplyEventsRequest *history.ReapplyEventsRequest,
		opts ...yarpc.CallOption,
	) error

	RecordActivityTaskHeartbeat(
		ctx context.Context,
		HeartbeatRequest *history.RecordActivityTaskHeartbeatRequest,
		opts ...yarpc.CallOption,
	) (*shared.RecordActivityTaskHeartbeatResponse, error)

	RecordActivityTaskStarted(
		ctx context.Context,
		AddRequest *history.RecordActivityTaskStartedRequest,
		opts ...yarpc.CallOption,
	) (*history.RecordActivityTaskStartedResponse, error)

	RecordChildExecutionCompleted(
		ctx context.Context,
		CompletionRequest *history.RecordChildExecutionCompletedRequest,
		opts ...yarpc.CallOption,
	) error

	RecordDecisionTaskStarted(
		ctx context.Context,
		AddRequest *history.RecordDecisionTaskStartedRequest,
		opts ...yarpc.CallOption,
	) (*history.RecordDecisionTaskStartedResponse, error)

	RefreshWorkflowTasks(
		ctx context.Context,
		Request *history.RefreshWorkflowTasksRequest,
		opts ...yarpc.CallOption,
	) error

	RemoveSignalMutableState(
		ctx context.Context,
		RemoveRequest *history.RemoveSignalMutableStateRequest,
		opts ...yarpc.CallOption,
	) error

	RemoveTask(
		ctx context.Context,
		Request *shared.RemoveTaskRequest,
		opts ...yarpc.CallOption,
	) error

	ReplicateEventsV2(
		ctx context.Context,
		ReplicateV2Request *history.ReplicateEventsV2Request,
		opts ...yarpc.CallOption,
	) error

	RequestCancelWorkflowExecution(
		ctx context.Context,
		CancelRequest *history.RequestCancelWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) error

	ResetQueue(
		ctx context.Context,
		Request *shared.ResetQueueRequest,
		opts ...yarpc.CallOption,
	) error

	ResetStickyTaskList(
		ctx context.Context,
		ResetRequest *history.ResetStickyTaskListRequest,
		opts ...yarpc.CallOption,
	) (*history.ResetStickyTaskListResponse, error)

	ResetWorkflowExecution(
		ctx context.Context,
		ResetRequest *history.ResetWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) (*shared.ResetWorkflowExecutionResponse, error)

	RespondActivityTaskCanceled(
		ctx context.Context,
		CanceledRequest *history.RespondActivityTaskCanceledRequest,
		opts ...yarpc.CallOption,
	) error

	RespondActivityTaskCompleted(
		ctx context.Context,
		CompleteRequest *history.RespondActivityTaskCompletedRequest,
		opts ...yarpc.CallOption,
	) error

	RespondActivityTaskFailed(
		ctx context.Context,
		FailRequest *history.RespondActivityTaskFailedRequest,
		opts ...yarpc.CallOption,
	) error

	RespondCrossClusterTasksCompleted(
		ctx context.Context,
		Request *shared.RespondCrossClusterTasksCompletedRequest,
		opts ...yarpc.CallOption,
	) (*shared.RespondCrossClusterTasksCompletedResponse, error)

	RespondDecisionTaskCompleted(
		ctx context.Context,
		CompleteRequest *history.RespondDecisionTaskCompletedRequest,
		opts ...yarpc.CallOption,
	) (*history.RespondDecisionTaskCompletedResponse, error)

	RespondDecisionTaskFailed(
		ctx context.Context,
		FailedRequest *history.RespondDecisionTaskFailedRequest,
		opts ...yarpc.CallOption,
	) error

	ScheduleDecisionTask(
		ctx context.Context,
		ScheduleRequest *history.ScheduleDecisionTaskRequest,
		opts ...yarpc.CallOption,
	) error

	SignalWithStartWorkflowExecution(
		ctx context.Context,
		SignalWithStartRequest *history.SignalWithStartWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) (*shared.StartWorkflowExecutionResponse, error)

	SignalWorkflowExecution(
		ctx context.Context,
		SignalRequest *history.SignalWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) error

	StartWorkflowExecution(
		ctx context.Context,
		StartRequest *history.StartWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) (*shared.StartWorkflowExecutionResponse, error)

	SyncActivity(
		ctx context.Context,
		SyncActivityRequest *history.SyncActivityRequest,
		opts ...yarpc.CallOption,
	) error

	SyncShardStatus(
		ctx context.Context,
		SyncShardStatusRequest *history.SyncShardStatusRequest,
		opts ...yarpc.CallOption,
	) error

	TerminateWorkflowExecution(
		ctx context.Context,
		TerminateRequest *history.TerminateWorkflowExecutionRequest,
		opts ...yarpc.CallOption,
	) error
}

// New builds a new client for the HistoryService service.
//
// 	client := historyserviceclient.New(dispatcher.ClientConfig("historyservice"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "HistoryService",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	c thrift.Client
}

func (c client) CloseShard(
	ctx context.Context,
	_Request *shared.CloseShardRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_CloseShard_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_CloseShard_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_CloseShard_Helper.UnwrapResponse(&result)
	return
}

func (c client) DescribeHistoryHost(
	ctx context.Context,
	_Request *shared.DescribeHistoryHostRequest,
	opts ...yarpc.CallOption,
) (success *shared.DescribeHistoryHostResponse, err error) {

	args := history.HistoryService_DescribeHistoryHost_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_DescribeHistoryHost_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_DescribeHistoryHost_Helper.UnwrapResponse(&result)
	return
}

func (c client) DescribeMutableState(
	ctx context.Context,
	_Request *history.DescribeMutableStateRequest,
	opts ...yarpc.CallOption,
) (success *history.DescribeMutableStateResponse, err error) {

	args := history.HistoryService_DescribeMutableState_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_DescribeMutableState_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_DescribeMutableState_Helper.UnwrapResponse(&result)
	return
}

func (c client) DescribeQueue(
	ctx context.Context,
	_Request *shared.DescribeQueueRequest,
	opts ...yarpc.CallOption,
) (success *shared.DescribeQueueResponse, err error) {

	args := history.HistoryService_DescribeQueue_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_DescribeQueue_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_DescribeQueue_Helper.UnwrapResponse(&result)
	return
}

func (c client) DescribeWorkflowExecution(
	ctx context.Context,
	_DescribeRequest *history.DescribeWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.DescribeWorkflowExecutionResponse, err error) {

	args := history.HistoryService_DescribeWorkflowExecution_Helper.Args(_DescribeRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_DescribeWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_DescribeWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) GetCrossClusterTasks(
	ctx context.Context,
	_Request *shared.GetCrossClusterTasksRequest,
	opts ...yarpc.CallOption,
) (success *shared.GetCrossClusterTasksResponse, err error) {

	args := history.HistoryService_GetCrossClusterTasks_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_GetCrossClusterTasks_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_GetCrossClusterTasks_Helper.UnwrapResponse(&result)
	return
}

func (c client) GetDLQReplicationMessages(
	ctx context.Context,
	_Request *replicator.GetDLQReplicationMessagesRequest,
	opts ...yarpc.CallOption,
) (success *replicator.GetDLQReplicationMessagesResponse, err error) {

	args := history.HistoryService_GetDLQReplicationMessages_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_GetDLQReplicationMessages_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_GetDLQReplicationMessages_Helper.UnwrapResponse(&result)
	return
}

func (c client) GetFailoverInfo(
	ctx context.Context,
	_Request *history.GetFailoverInfoRequest,
	opts ...yarpc.CallOption,
) (success *history.GetFailoverInfoResponse, err error) {

	args := history.HistoryService_GetFailoverInfo_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_GetFailoverInfo_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_GetFailoverInfo_Helper.UnwrapResponse(&result)
	return
}

func (c client) GetMutableState(
	ctx context.Context,
	_GetRequest *history.GetMutableStateRequest,
	opts ...yarpc.CallOption,
) (success *history.GetMutableStateResponse, err error) {

	args := history.HistoryService_GetMutableState_Helper.Args(_GetRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_GetMutableState_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_GetMutableState_Helper.UnwrapResponse(&result)
	return
}

func (c client) GetReplicationMessages(
	ctx context.Context,
	_Request *replicator.GetReplicationMessagesRequest,
	opts ...yarpc.CallOption,
) (success *replicator.GetReplicationMessagesResponse, err error) {

	args := history.HistoryService_GetReplicationMessages_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_GetReplicationMessages_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_GetReplicationMessages_Helper.UnwrapResponse(&result)
	return
}

func (c client) MergeDLQMessages(
	ctx context.Context,
	_Request *replicator.MergeDLQMessagesRequest,
	opts ...yarpc.CallOption,
) (success *replicator.MergeDLQMessagesResponse, err error) {

	args := history.HistoryService_MergeDLQMessages_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_MergeDLQMessages_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_MergeDLQMessages_Helper.UnwrapResponse(&result)
	return
}

func (c client) NotifyFailoverMarkers(
	ctx context.Context,
	_Request *history.NotifyFailoverMarkersRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_NotifyFailoverMarkers_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_NotifyFailoverMarkers_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_NotifyFailoverMarkers_Helper.UnwrapResponse(&result)
	return
}

func (c client) PollMutableState(
	ctx context.Context,
	_PollRequest *history.PollMutableStateRequest,
	opts ...yarpc.CallOption,
) (success *history.PollMutableStateResponse, err error) {

	args := history.HistoryService_PollMutableState_Helper.Args(_PollRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_PollMutableState_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_PollMutableState_Helper.UnwrapResponse(&result)
	return
}

func (c client) PurgeDLQMessages(
	ctx context.Context,
	_Request *replicator.PurgeDLQMessagesRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_PurgeDLQMessages_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_PurgeDLQMessages_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_PurgeDLQMessages_Helper.UnwrapResponse(&result)
	return
}

func (c client) QueryWorkflow(
	ctx context.Context,
	_QueryRequest *history.QueryWorkflowRequest,
	opts ...yarpc.CallOption,
) (success *history.QueryWorkflowResponse, err error) {

	args := history.HistoryService_QueryWorkflow_Helper.Args(_QueryRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_QueryWorkflow_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_QueryWorkflow_Helper.UnwrapResponse(&result)
	return
}

func (c client) ReadDLQMessages(
	ctx context.Context,
	_Request *replicator.ReadDLQMessagesRequest,
	opts ...yarpc.CallOption,
) (success *replicator.ReadDLQMessagesResponse, err error) {

	args := history.HistoryService_ReadDLQMessages_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ReadDLQMessages_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_ReadDLQMessages_Helper.UnwrapResponse(&result)
	return
}

func (c client) ReapplyEvents(
	ctx context.Context,
	_ReapplyEventsRequest *history.ReapplyEventsRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_ReapplyEvents_Helper.Args(_ReapplyEventsRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ReapplyEvents_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_ReapplyEvents_Helper.UnwrapResponse(&result)
	return
}

func (c client) RecordActivityTaskHeartbeat(
	ctx context.Context,
	_HeartbeatRequest *history.RecordActivityTaskHeartbeatRequest,
	opts ...yarpc.CallOption,
) (success *shared.RecordActivityTaskHeartbeatResponse, err error) {

	args := history.HistoryService_RecordActivityTaskHeartbeat_Helper.Args(_HeartbeatRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RecordActivityTaskHeartbeat_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RecordActivityTaskHeartbeat_Helper.UnwrapResponse(&result)
	return
}

func (c client) RecordActivityTaskStarted(
	ctx context.Context,
	_AddRequest *history.RecordActivityTaskStartedRequest,
	opts ...yarpc.CallOption,
) (success *history.RecordActivityTaskStartedResponse, err error) {

	args := history.HistoryService_RecordActivityTaskStarted_Helper.Args(_AddRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RecordActivityTaskStarted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RecordActivityTaskStarted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RecordChildExecutionCompleted(
	ctx context.Context,
	_CompletionRequest *history.RecordChildExecutionCompletedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RecordChildExecutionCompleted_Helper.Args(_CompletionRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RecordChildExecutionCompleted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RecordChildExecutionCompleted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RecordDecisionTaskStarted(
	ctx context.Context,
	_AddRequest *history.RecordDecisionTaskStartedRequest,
	opts ...yarpc.CallOption,
) (success *history.RecordDecisionTaskStartedResponse, err error) {

	args := history.HistoryService_RecordDecisionTaskStarted_Helper.Args(_AddRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RecordDecisionTaskStarted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RecordDecisionTaskStarted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RefreshWorkflowTasks(
	ctx context.Context,
	_Request *history.RefreshWorkflowTasksRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RefreshWorkflowTasks_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RefreshWorkflowTasks_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RefreshWorkflowTasks_Helper.UnwrapResponse(&result)
	return
}

func (c client) RemoveSignalMutableState(
	ctx context.Context,
	_RemoveRequest *history.RemoveSignalMutableStateRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RemoveSignalMutableState_Helper.Args(_RemoveRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RemoveSignalMutableState_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RemoveSignalMutableState_Helper.UnwrapResponse(&result)
	return
}

func (c client) RemoveTask(
	ctx context.Context,
	_Request *shared.RemoveTaskRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RemoveTask_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RemoveTask_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RemoveTask_Helper.UnwrapResponse(&result)
	return
}

func (c client) ReplicateEventsV2(
	ctx context.Context,
	_ReplicateV2Request *history.ReplicateEventsV2Request,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_ReplicateEventsV2_Helper.Args(_ReplicateV2Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ReplicateEventsV2_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_ReplicateEventsV2_Helper.UnwrapResponse(&result)
	return
}

func (c client) RequestCancelWorkflowExecution(
	ctx context.Context,
	_CancelRequest *history.RequestCancelWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RequestCancelWorkflowExecution_Helper.Args(_CancelRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RequestCancelWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RequestCancelWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) ResetQueue(
	ctx context.Context,
	_Request *shared.ResetQueueRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_ResetQueue_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ResetQueue_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_ResetQueue_Helper.UnwrapResponse(&result)
	return
}

func (c client) ResetStickyTaskList(
	ctx context.Context,
	_ResetRequest *history.ResetStickyTaskListRequest,
	opts ...yarpc.CallOption,
) (success *history.ResetStickyTaskListResponse, err error) {

	args := history.HistoryService_ResetStickyTaskList_Helper.Args(_ResetRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ResetStickyTaskList_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_ResetStickyTaskList_Helper.UnwrapResponse(&result)
	return
}

func (c client) ResetWorkflowExecution(
	ctx context.Context,
	_ResetRequest *history.ResetWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.ResetWorkflowExecutionResponse, err error) {

	args := history.HistoryService_ResetWorkflowExecution_Helper.Args(_ResetRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ResetWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_ResetWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondActivityTaskCanceled(
	ctx context.Context,
	_CanceledRequest *history.RespondActivityTaskCanceledRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RespondActivityTaskCanceled_Helper.Args(_CanceledRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondActivityTaskCanceled_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RespondActivityTaskCanceled_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondActivityTaskCompleted(
	ctx context.Context,
	_CompleteRequest *history.RespondActivityTaskCompletedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RespondActivityTaskCompleted_Helper.Args(_CompleteRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondActivityTaskCompleted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RespondActivityTaskCompleted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondActivityTaskFailed(
	ctx context.Context,
	_FailRequest *history.RespondActivityTaskFailedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RespondActivityTaskFailed_Helper.Args(_FailRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondActivityTaskFailed_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RespondActivityTaskFailed_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondCrossClusterTasksCompleted(
	ctx context.Context,
	_Request *shared.RespondCrossClusterTasksCompletedRequest,
	opts ...yarpc.CallOption,
) (success *shared.RespondCrossClusterTasksCompletedResponse, err error) {

	args := history.HistoryService_RespondCrossClusterTasksCompleted_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondCrossClusterTasksCompleted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RespondCrossClusterTasksCompleted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondDecisionTaskCompleted(
	ctx context.Context,
	_CompleteRequest *history.RespondDecisionTaskCompletedRequest,
	opts ...yarpc.CallOption,
) (success *history.RespondDecisionTaskCompletedResponse, err error) {

	args := history.HistoryService_RespondDecisionTaskCompleted_Helper.Args(_CompleteRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondDecisionTaskCompleted_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_RespondDecisionTaskCompleted_Helper.UnwrapResponse(&result)
	return
}

func (c client) RespondDecisionTaskFailed(
	ctx context.Context,
	_FailedRequest *history.RespondDecisionTaskFailedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_RespondDecisionTaskFailed_Helper.Args(_FailedRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_RespondDecisionTaskFailed_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_RespondDecisionTaskFailed_Helper.UnwrapResponse(&result)
	return
}

func (c client) ScheduleDecisionTask(
	ctx context.Context,
	_ScheduleRequest *history.ScheduleDecisionTaskRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_ScheduleDecisionTask_Helper.Args(_ScheduleRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_ScheduleDecisionTask_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_ScheduleDecisionTask_Helper.UnwrapResponse(&result)
	return
}

func (c client) SignalWithStartWorkflowExecution(
	ctx context.Context,
	_SignalWithStartRequest *history.SignalWithStartWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.StartWorkflowExecutionResponse, err error) {

	args := history.HistoryService_SignalWithStartWorkflowExecution_Helper.Args(_SignalWithStartRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_SignalWithStartWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_SignalWithStartWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) SignalWorkflowExecution(
	ctx context.Context,
	_SignalRequest *history.SignalWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_SignalWorkflowExecution_Helper.Args(_SignalRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_SignalWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_SignalWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) StartWorkflowExecution(
	ctx context.Context,
	_StartRequest *history.StartWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.StartWorkflowExecutionResponse, err error) {

	args := history.HistoryService_StartWorkflowExecution_Helper.Args(_StartRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_StartWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = history.HistoryService_StartWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}

func (c client) SyncActivity(
	ctx context.Context,
	_SyncActivityRequest *history.SyncActivityRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_SyncActivity_Helper.Args(_SyncActivityRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_SyncActivity_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_SyncActivity_Helper.UnwrapResponse(&result)
	return
}

func (c client) SyncShardStatus(
	ctx context.Context,
	_SyncShardStatusRequest *history.SyncShardStatusRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_SyncShardStatus_Helper.Args(_SyncShardStatusRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_SyncShardStatus_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_SyncShardStatus_Helper.UnwrapResponse(&result)
	return
}

func (c client) TerminateWorkflowExecution(
	ctx context.Context,
	_TerminateRequest *history.TerminateWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := history.HistoryService_TerminateWorkflowExecution_Helper.Args(_TerminateRequest)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result history.HistoryService_TerminateWorkflowExecution_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = history.HistoryService_TerminateWorkflowExecution_Helper.UnwrapResponse(&result)
	return
}
