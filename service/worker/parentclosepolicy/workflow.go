// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package parentclosepolicy

import (
	"context"
	"time"

	"go.uber.org/cadence"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"

	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/log/tag"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/common/types"
)

type (
	contextKey string
)

const (
	processorContextKey contextKey = "processorContext"
	// processorTaskListName is the tasklist name
	processorTaskListName = "cadence-sys-processor-parent-close-policy"
	// processorWFTypeName is the workflow type
	processorWFTypeName   = "cadence-sys-parent-close-policy-workflow"
	processorActivityName = "cadence-sys-parent-close-policy-activity"
	infiniteDuration      = 20 * 365 * 24 * time.Hour
	processorChannelName  = "ParentClosePolicyProcessorChannelName"
)

type (
	// RequestDetail defines detail of each workflow to process
	RequestDetail struct {
		DomainName string
		WorkflowID string
		RunID      string
		Policy     types.ParentClosePolicy
	}

	// Request defines the request for parent close policy
	Request struct {
		Executions []RequestDetail

		// DEPRECATED: the following field is deprecated since childworkflow
		// might in a different domain, use the DomainName field in RequestDetail
		DomainName string
	}
)

var (
	retryPolicy = cadence.RetryPolicy{
		InitialInterval:    10 * time.Second,
		BackoffCoefficient: 1.7,
		MaximumInterval:    5 * time.Minute,
		ExpirationInterval: infiniteDuration,
	}

	activityOptions = workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    5 * time.Minute,
		RetryPolicy:            &retryPolicy,
	}
)

func init() {
	workflow.RegisterWithOptions(ProcessorWorkflow, workflow.RegisterOptions{Name: processorWFTypeName})
	activity.RegisterWithOptions(ProcessorActivity, activity.RegisterOptions{Name: processorActivityName})
}

// ProcessorWorkflow is the workflow that performs actions for ParentClosePolicy
func ProcessorWorkflow(ctx workflow.Context) error {
	requestCh := workflow.GetSignalChannel(ctx, processorChannelName)
	for {
		var request Request
		if !requestCh.ReceiveAsync(&request) {
			// no more request
			break
		}

		opt := workflow.WithActivityOptions(ctx, activityOptions)
		_ = workflow.ExecuteActivity(opt, processorActivityName, request).Get(ctx, nil)
	}
	return nil
}

// ProcessorActivity is activity for processing batch operation
func ProcessorActivity(ctx context.Context, request Request) error {
	processor := ctx.Value(processorContextKey).(*Processor)
	for _, execution := range request.Executions {
		domainName := execution.DomainName
		if domainName == "" {
			// for backward compatibility
			domainName = request.DomainName
		}
		var err error
		switch execution.Policy {
		case types.ParentClosePolicyAbandon:
			//no-op
			continue
		case types.ParentClosePolicyTerminate:
			err = processor.frontendClient.TerminateWorkflowExecution(ctx, &types.TerminateWorkflowExecutionRequest{
				Domain: domainName,
				WorkflowExecution: &types.WorkflowExecution{
					WorkflowID: execution.WorkflowID,
					RunID:      execution.RunID,
				},
				Reason:   "by parent close policy",
				Identity: processorWFTypeName,
			})
		case types.ParentClosePolicyRequestCancel:
			err = processor.frontendClient.RequestCancelWorkflowExecution(ctx, &types.RequestCancelWorkflowExecutionRequest{
				Domain: domainName,
				WorkflowExecution: &types.WorkflowExecution{
					WorkflowID: execution.WorkflowID,
					RunID:      execution.RunID,
				},
				Identity: processorWFTypeName,
			})
		}

		if err != nil {
			switch err.(type) {
			case *types.EntityNotExistsError, *types.CancellationAlreadyRequestedError:
				err = nil
			}
		}

		if err != nil {
			processor.metricsClient.IncCounter(metrics.ParentClosePolicyProcessorScope, metrics.ParentClosePolicyProcessorFailures)
			getActivityLogger(ctx).Error("failed to process parent close policy", tag.Error(err))
			return err
		}
		processor.metricsClient.IncCounter(metrics.ParentClosePolicyProcessorScope, metrics.ParentClosePolicyProcessorSuccess)
	}
	return nil
}

func getActivityLogger(ctx context.Context) log.Logger {
	processor := ctx.Value(processorContextKey).(*Processor)
	wfInfo := activity.GetInfo(ctx)
	return processor.logger.WithTags(
		tag.WorkflowID(wfInfo.WorkflowExecution.ID),
		tag.WorkflowRunID(wfInfo.WorkflowExecution.RunID),
		tag.WorkflowDomainName(wfInfo.WorkflowDomain),
	)
}
