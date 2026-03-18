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
	region := terraform.Output(t, opts, "region")
	require.NotEmpty(t, region, "region output should be set")

	// Verify Terraform outputs
	id := terraform.Output(t, opts, "id")
	name := terraform.Output(t, opts, "name")
	arn := terraform.Output(t, opts, "arn")
	state := terraform.Output(t, opts, "state")
	scheduleNames := terraform.OutputList(t, opts, "schedule_names")

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
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	opts := ctx.TerratestTerraformOptions()

	// Get region from Terraform output (matches where resources were created)
	region := terraform.Output(t, opts, "region")
	require.NotEmpty(t, region, "region output should be set")

	// Verify Terraform outputs (read-only)
	id := terraform.Output(t, opts, "id")
	name := terraform.Output(t, opts, "name")
	arn := terraform.Output(t, opts, "arn")
	state := terraform.Output(t, opts, "state")
	scheduleNames := terraform.OutputList(t, opts, "schedule_names")

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
