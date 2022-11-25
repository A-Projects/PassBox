package folder

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"passbox.server/domain"
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideController,
	ProvideService,

	wire.Bind(new(domain.FolderController), new(*controller)),
	wire.Bind(new(domain.FolderService), new(*service)),
)

func ProvideController(s domain.FolderService) *controller {
	return &controller{
		folderService: s,
	}
}

func ProvideService(log *logrus.Logger) *service {
	return &service{
		logger: log,
	}
}
