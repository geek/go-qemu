// Code generated by "stringer -type=Status -output=string.gen.go"; DO NOT EDIT

package qemu

import "fmt"

const _Status_name = "StatusDebugStatusInMigrateStatusInternalErrorStatusIOErrorStatusPausedStatusPostMigrateStatusPreLaunchStatusFinishMigrateStatusRestoreVMStatusRunningStatusSaveVMStatusShutdownStatusSuspendedStatusWatchdogStatusGuestPanicked"

var _Status_index = [...]uint8{0, 11, 26, 45, 58, 70, 87, 102, 121, 136, 149, 161, 175, 190, 204, 223}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
