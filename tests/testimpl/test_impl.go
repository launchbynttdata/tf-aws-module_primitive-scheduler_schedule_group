package testimpl

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	schedulertypes "github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	opts := ctx.TerratestTerraformOptions()

	// Get region from Terraform output (matches where resources were created)
	region := terraform.OutputContext(t, context.Background(), opts, "region")
	require.NotEmpty(t, region, "region output should be set")

	// Verify Terraform outputs
	id := terraform.OutputContext(t, context.Background(), opts, "id")
	name := terraform.OutputContext(t, context.Background(), opts, "name")
	arn := terraform.OutputContext(t, context.Background(), opts, "arn")
	state := terraform.OutputContext(t, context.Background(), opts, "state")
	scheduleNames := terraform.OutputListContext(t, context.Background(), opts, "schedule_names")

	require.NotEmpty(t, id, "id output should be set")
	require.NotEmpty(t, name, "name output should be set")
	require.NotEmpty(t, arn, "arn output should be set")
	require.NotEmpty(t, state, "state output should be set")
	require.Len(t, scheduleNames, 2, "schedule_names should have 2 entries")

	// Verify id equals name (documented behavior)
	assert.Equal(t, name, id, "id should equal name for schedule group")

	// Verify state is ACTIVE
	assert.Equal(t, "ACTIVE", state, "schedule group state should be ACTIVE")

	// Verify via AWS API
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	require.NoError(t, err)

	schedulerClient := scheduler.NewFromConfig(cfg)
	getOut, err := schedulerClient.GetScheduleGroup(context.Background(), &scheduler.GetScheduleGroupInput{
		Name: aws.String(name),
	})
	require.NoError(t, err)
	require.NotNil(t, getOut)

	assert.Equal(t, name, aws.ToString(getOut.Name), "API name should match Terraform output")
	assert.Equal(t, arn, aws.ToString(getOut.Arn), "API ARN should match Terraform output")
	assert.Equal(t, schedulertypes.ScheduleGroupStateActive, getOut.State, "API state should be ACTIVE")

	// Verify both schedules exist and are in the group
	for _, scheduleName := range scheduleNames {
		schedOut, err := schedulerClient.GetSchedule(context.Background(), &scheduler.GetScheduleInput{
			Name:      aws.String(scheduleName),
			GroupName: aws.String(name),
		})
		require.NoError(t, err, "schedule %s should exist", scheduleName)
		assert.Equal(t, name, aws.ToString(schedOut.GroupName), "schedule %s should be in the schedule group", scheduleName)
	}

	// Write-path: update a schedule state and verify (functional test only)
	firstSchedule := scheduleNames[0]
	schedOut, err := schedulerClient.GetSchedule(context.Background(), &scheduler.GetScheduleInput{
		Name:      aws.String(firstSchedule),
		GroupName: aws.String(name),
	})
	require.NoError(t, err, "schedule %s should exist for write-path test", firstSchedule)

	_, err = schedulerClient.UpdateSchedule(context.Background(), &scheduler.UpdateScheduleInput{
		GroupName:          aws.String(name),
		Name:               aws.String(firstSchedule),
		FlexibleTimeWindow: schedOut.FlexibleTimeWindow,
		ScheduleExpression: schedOut.ScheduleExpression,
		Target:             schedOut.Target,
		State:              schedulertypes.ScheduleStateDisabled,
	})
	require.NoError(t, err, "UpdateSchedule to DISABLED should succeed")

	schedOut, err = schedulerClient.GetSchedule(context.Background(), &scheduler.GetScheduleInput{
		Name:      aws.String(firstSchedule),
		GroupName: aws.String(name),
	})
	require.NoError(t, err)
	assert.Equal(t, schedulertypes.ScheduleStateDisabled, schedOut.State, "schedule state should be DISABLED after update")

	_, err = schedulerClient.UpdateSchedule(context.Background(), &scheduler.UpdateScheduleInput{
		GroupName:          aws.String(name),
		Name:               aws.String(firstSchedule),
		FlexibleTimeWindow: schedOut.FlexibleTimeWindow,
		ScheduleExpression: schedOut.ScheduleExpression,
		Target:             schedOut.Target,
		State:              schedulertypes.ScheduleStateEnabled,
	})
	require.NoError(t, err, "UpdateSchedule back to ENABLED should succeed")
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	opts := ctx.TerratestTerraformOptions()

	// Get region from Terraform output (matches where resources were created)
	region := terraform.OutputContext(t, context.Background(), opts, "region")
	require.NotEmpty(t, region, "region output should be set")

	// Verify Terraform outputs (read-only)
	id := terraform.OutputContext(t, context.Background(), opts, "id")
	name := terraform.OutputContext(t, context.Background(), opts, "name")
	arn := terraform.OutputContext(t, context.Background(), opts, "arn")
	state := terraform.OutputContext(t, context.Background(), opts, "state")
	scheduleNames := terraform.OutputListContext(t, context.Background(), opts, "schedule_names")

	require.NotEmpty(t, id, "id output should be set")
	require.NotEmpty(t, name, "name output should be set")
	require.NotEmpty(t, arn, "arn output should be set")
	require.NotEmpty(t, state, "state output should be set")
	require.Len(t, scheduleNames, 2, "schedule_names should have 2 entries")

	assert.Equal(t, name, id, "id should equal name for schedule group")
	assert.Equal(t, "ACTIVE", state, "schedule group state should be ACTIVE")

	// Verify via AWS API (read-only)
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	require.NoError(t, err)

	schedulerClient := scheduler.NewFromConfig(cfg)
	getOut, err := schedulerClient.GetScheduleGroup(context.Background(), &scheduler.GetScheduleGroupInput{
		Name: aws.String(name),
	})
	require.NoError(t, err)
	require.NotNil(t, getOut)

	assert.Equal(t, name, aws.ToString(getOut.Name), "API name should match Terraform output")
	assert.Equal(t, arn, aws.ToString(getOut.Arn), "API ARN should match Terraform output")
	assert.Equal(t, schedulertypes.ScheduleGroupStateActive, getOut.State, "API state should be ACTIVE")

	// Verify both schedules exist and are in the group
	for _, scheduleName := range scheduleNames {
		schedOut, err := schedulerClient.GetSchedule(context.Background(), &scheduler.GetScheduleInput{
			Name:      aws.String(scheduleName),
			GroupName: aws.String(name),
		})
		require.NoError(t, err, "schedule %s should exist", scheduleName)
		assert.Equal(t, name, aws.ToString(schedOut.GroupName), "schedule %s should be in the schedule group", scheduleName)
	}
}
