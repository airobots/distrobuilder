package generators

import (
	"path"
	
	shell "github.com/airobots/ntgoutil/shell"
	"github.com/lxc/distrobuilder/image"
	"github.com/lxc/distrobuilder/shared"
)

type CopyGenerator struct{}

func (g CopyGenerator) RunLXC(cacheDir, sourceDir string, img *image.LXCImage, defFile shared.DefinitionFile) error {
	err := g.Run(cacheDir, sourceDir, defFile)
	if err != nil {
		return err
	}
	
	return nil
}

func (g CopyGenerator) RunLXD(cacheDir, sourceDir string, img *image.LXDImage, defFile shared.DefinitionFile) error {
	return g.Run(cacheDir, sourceDir, defFile)
}

func (g CopyGenerator) Run(cacheDir, sourceDir string, defFile shared.DefinitionFile) error {
	return shell.Copy(defFile.Content, path.Join(sourceDir, defFile.Path), defFile.Uid, defFile.Gid, defFile.FileMode, defFile.DirectoryMode)
}