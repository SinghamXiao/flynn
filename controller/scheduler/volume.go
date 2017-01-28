package main

import (
	ct "github.com/flynn/flynn/controller/types"
	"github.com/flynn/flynn/host/volume"
)

type Volume struct {
	ct.Volume
}

func (v *Volume) Info() *volume.Info {
	return &volume.Info{
		ID:   v.ID,
		Type: v.Type,
		Meta: v.Meta,
	}
}

func NewVolume(info *volume.Info, hostID string) *Volume {
	return &Volume{
		Volume: ct.Volume{
			VolumeReq: ct.VolumeReq{
				Path:         info.Meta["flynn-controller.path"],
				DeleteOnStop: info.Meta["flynn-controller.delete_on_stop"] == "true",
			},
			ID:        info.ID,
			HostID:    hostID,
			Type:      info.Type,
			AppID:     info.Meta["flynn-controller.app"],
			ReleaseID: info.Meta["flynn-controller.release"],
			JobType:   info.Meta["flynn-controller.type"],
			Meta:      info.Meta,
			CreatedAt: &info.CreatedAt,
		},
	}
}

type VolumeEvent struct {
	Type   VolumeEventType
	Volume *Volume
}

type VolumeEventType string

const (
	VolumeEventTypeCreate     VolumeEventType = "create"
	VolumeEventTypeDestroy    VolumeEventType = "destroy"
	VolumeEventTypeController VolumeEventType = "controller"
)
