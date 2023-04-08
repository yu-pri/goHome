package home

const (
	internalSensorID        = "28-0315a14596ff"
	reverseSensorID         = "28-0115a32c38ff"
	entryRoomSensorID       = "28-0115a3752aff"
	heaterSensorID          = "28-0115a32295ff"
	_waterBoilerSensorID    = "28-0415a1cbb3ff"
	_outsideSensorID        = "28-0315a1a9cfff"
	_recuperatorSensorID    = "28-0315a1d318ff"
)

/*
HeaterPumpRunThreshold Run heater pump, stop heater pumb if temp 5 degree less
*/
const HeaterPumpRunThreshold = 30

/*
CommandOnPumpr1 command to toggle
*/
const CommandOnPumpr1 = "onPumpr1"

/*
CommandOffPumpr1 command to toggle
*/
const CommandOffPumpr1 = "offPumpr1"

/*
ElectroOnFrom - timezone for electricity metter - from
*/
const ElectroOnFrom = 23

/*
ElectroOnTo - timezone for electricity metter - from
*/
const ElectroOnTo = 7
