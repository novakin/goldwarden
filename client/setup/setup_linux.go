//go:build linux

package setup

import (
	"fmt"

	"github.com/quexten/goldwarden/agent/config"
	"github.com/quexten/goldwarden/cmd"
)

func VerifySetup(runtimeConfig config.RuntimeConfig) bool {
	if !cmd.IsPolkitSetup() && !runtimeConfig.DisableAuth {
		fmt.Println("Polkit is not setup. Run 'goldwarden setup polkit' to set it up.")
		return false
	}

	return true
}
