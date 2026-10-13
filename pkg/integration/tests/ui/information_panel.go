package ui

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var InformationPanel = NewIntegrationTest(NewIntegrationTestArgs{
	Description: "The information panel shows only the version when no mode is active, even with mouse events enabled",
	SetupConfig: func(cfg *config.AppConfig) {
		cfg.GetUserConfig().Gui.MouseEvents = true
	},
	SetupRepo: func(shell *Shell) {},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Information().Content(Equals("unversioned"))
	},
})
