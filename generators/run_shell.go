package generators

import (
	shell "github.com/airobots/ntgoutil/shell"
	"github.com/lxc/distrobuilder/image"
	"github.com/lxc/distrobuilder/shared"
)

type RunShellGenerator struct{}

func (g RunShellGenerator) RunLXC(cacheDir, sourceDir string, img *image.LXCImage, defFile shared.DefinitionFile) error {
	err := g.Run(cacheDir, sourceDir, defFile)
	if err != nil {
		return err
	}
	
	return nil
}

func (g RunShellGenerator) RunLXD(cacheDir, sourceDir string, img *image.LXDImage, defFile shared.DefinitionFile) error {
	return g.Run(cacheDir, sourceDir, defFile)
}

func (g RunShellGenerator) Run(cacheDir, sourceDir string, defFile shared.DefinitionFile) error {
	return shell.RunScript(defFile.Content, []string{"_distrobuild_dir="+sourceDir})
}