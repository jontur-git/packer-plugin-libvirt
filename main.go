package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/plugin"

	"github.com/jontur-git/packer-plugin-libvirt/builder/libvirt"
	"github.com/jontur-git/packer-plugin-libvirt/builder/libvirt"
	"github.com/jontur-git/packer-plugin-libvirt/version"
	"github.com/jontur-git/packer-plugin-libvirt/version"
)

func main() {
	pps := plugin.NewSet()
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(libvirt.Builder))
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(libvirt.Builder))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
