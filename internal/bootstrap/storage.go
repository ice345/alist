package bootstrap

import (
	"context"

	"github.com/NewAlist/alist/internal/conf"
	"github.com/NewAlist/alist/internal/db"
	"github.com/NewAlist/alist/internal/model"
	"github.com/NewAlist/alist/internal/op"
	"github.com/NewAlist/alist/pkg/utils"
)

func LoadStorages() {
	storages, err := db.GetEnabledStorages()
	if err != nil {
		utils.Log.Fatalf("failed get enabled storages: %+v", err)
	}
	go func(storages []model.Storage) {
		for i := range storages {
			err := op.LoadStorage(context.Background(), storages[i])
			if err != nil {
				utils.Log.Errorf("failed get enabled storages: %+v", err)
			} else {
				utils.Log.Infof("success load storage: [%s], driver: [%s], order: [%d]",
					storages[i].MountPath, storages[i].Driver, storages[i].Order)
			}
		}
		conf.StoragesLoaded = true
	}(storages)
}
