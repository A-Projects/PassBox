//go:build wireinject
// +build wireinject

package folder

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func Wire(log *logrus.Logger) *controller {
	panic(wire.Build(ProviderSet))
}
