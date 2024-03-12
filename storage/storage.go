package storage

import (
	"github.com/rom2049/vehicle-server/storage/vehiclestore"
)

type Store interface {
	Vehicle() vehiclestore.Store
}
