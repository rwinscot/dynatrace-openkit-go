package openkitgo

import (
	"github.com/rwinscot/dynatrace-openkit-go/openkitgo/core"
	"github.com/rwinscot/dynatrace-openkit-go/openkitgo/interfaces"
)

func NewOpenKitBuilder(endpointURL string, applicationID string, deviceID int64) interfaces.OpenKitBuilder {
	return core.NewOpenKitBuilder(endpointURL, applicationID, deviceID)
}
