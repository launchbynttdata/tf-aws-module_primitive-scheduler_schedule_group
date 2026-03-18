package testimpl

import "github.com/launchbynttdata/lcaf-component-terratest/types"

// SchedulerScheduleGroupConfig holds test configuration for the scheduler schedule group module.
type SchedulerScheduleGroupConfig struct {
	types.GenericTFModuleConfig
	// Tags expected on the schedule group (for verification).
	Tags map[string]string
}
