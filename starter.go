package scheduleStarter

import (
	"github.com/dennesshen/photon-core-starter/core"
	"github.com/dennesshen/photon-scheduled-starter/schedule"
)

func init() {
	core.RegisterAddModule(schedule.Start)
	core.RegisterShutdownAddModule(schedule.Shutdown)
}
