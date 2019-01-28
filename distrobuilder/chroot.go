package main

import (
	"fmt"

	lxd "github.com/lxc/lxd/shared"

	"github.com/lxc/distrobuilder/managers"
	"github.com/lxc/distrobuilder/shared"
)

func managePackages(def shared.DefinitionPackages, actions []shared.DefinitionAction,
	release string) error {
	var err error

	manager := managers.Get(def.Manager)
	if manager == nil {
		return fmt.Errorf("Couldn't get manager")
	}

	// Handle repositories actions
	if def.Repositories != nil && len(def.Repositories) > 0 {
		if manager.RepoHandler == nil {
			return fmt.Errorf("No repository handler present")
		}

		for _, repo := range def.Repositories {
			err = manager.RepoHandler(repo)
			if err != nil {
				return fmt.Errorf("Error for repository %s: %s", repo.Name, err)
			}
		}
	}

	err = manager.Refresh()
	if err != nil {
		return err
	}

	if def.Update {
		err = manager.Update()
		if err != nil {
			return err
		}

		// Run post update hook
		for _, action := range actions {
			err = shared.RunScript(action.Action)
			if err != nil {
				return fmt.Errorf("Failed to run post-update: %s", err)
			}
		}
	}

	var installablePackages []string
	var installableArguments []string
	var removablePackages []string
	var removableArguments []string

	for _, set := range def.Sets {
		if len(set.Releases) > 0 && !lxd.StringInSlice(release, set.Releases) {
			continue
		}

		if set.Action == "install" {
			installablePackages = append(installablePackages, set.Packages...)
			installableArguments = append(installableArguments, set.Arguments...)
		} else if set.Action == "remove" {
			removablePackages = append(removablePackages, set.Packages...)
			removableArguments = append(removableArguments, set.Arguments...)
		}
	}

	err = manager.Install(installablePackages, installableArguments)
	if err != nil {
		return err
	}

	err = manager.Remove(removablePackages, removableArguments)
	if err != nil {
		return err
	}

	if def.Cleanup {
		err = manager.Clean()
		if err != nil {
			return err
		}
	}

	return nil
}
