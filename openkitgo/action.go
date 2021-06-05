package openkitgo

import "time"

type Action interface {
	ReportEvent(eventName string) Action
	ReportEventAt(eventName string, timestamp time.Time) Action

	ReportValue(valueName string, value interface{}) Action
	ReportValueAt(valueName string, value interface{}, timestamp time.Time) Action

	ReportError(errorName string, causeName string, causeDescription string, causeStack string) Action
	ReportErrorAt(errorName string, causeName string, causeDescription string, causeStack string, timestamp time.Time) Action

	// TODO TraceWebRequest()
	// TODO TraceWebRequestAt()

	LeaveAction() Action
	LeaveActionAt(timestamp time.Time) Action

	CancelAction() Action
	CancelActionAt(timestamp time.Time) Action

	GetDuration() time.Duration
}
