package domain

import (
	"github.com/fusion44/raspiblitz-backend/graph/model"
)

func (d *Domain) PushUpdatedDeviceInfo(i *model.UpdatedDeviceInfo) {
	info := d.InfoRepo.UpdateDeviceInfo(i)
	for _, v := range d.SetupRepo.SetupEventObservers {
		v.Channel <- &model.DeviceInfo{
			Version:   info.Version,
			SetupStep: info.SetupStep,
			BaseImage: info.BaseImage,
			CPU:       info.CPU,
			IsDocker:  info.IsDocker,
			State:     info.State,
			Chain:     info.Chain,
			Network:   info.Network,
			Message:   info.Message,
			HostName:  info.HostName,
		}
	}
}
