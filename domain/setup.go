package domain

import "github.com/fusion44/raspiblitz-backend/graph/model"

func (d *Domain) PushSetupEvent(i *model.PushSetupEventMessage) {
	for _, v := range d.SetupRepo.Observers {
		v.Channel <- &model.SetupInfoEvent{State: i.State, Message: i.Message}
	}
}
