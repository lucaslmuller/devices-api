package modules

import (
	"github.com/lucaslmuller/technical-test/internal/infrastructure"
)

func SetupModules(res *infrastructure.Resources) {
	setupDeviceModule(res)
}
