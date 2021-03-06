// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package internal

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/talos-systems/go-blockdevice/blockdevice/util"

	"github.com/talos-systems/talos/pkg/machinery/api/storage"
)

// Server implements storage.StorageService.
// TODO: this is not a full blown service yet, it's used as the common base in the machine and the maintenance services.
type Server struct{}

// Disks implements machine.MaintenanceService.
func (s *Server) Disks(ctx context.Context, in *empty.Empty) (reply *storage.DisksResponse, err error) {
	disks, err := util.GetDisks()
	if err != nil {
		return nil, err
	}

	diskList := make([]*storage.Disk, len(disks))

	for i, disk := range disks {
		diskList[i] = &storage.Disk{
			DeviceName: disk.DeviceName,
			Model:      disk.Model,
			Size:       disk.Size,
		}
	}

	reply = &storage.DisksResponse{
		Disks: diskList,
	}

	return reply, nil
}
