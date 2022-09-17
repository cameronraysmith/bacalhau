package bacalhau

import (
	"encoding/json"
	"os"

	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/filecoin-project/bacalhau/pkg/transport/libp2p"
	"github.com/multiformats/go-multiaddr"

	"github.com/spf13/cobra"
)

type IdInfo struct {
	ID string `json:"ID"`
}

var idCmd = &cobra.Command{
	Use:   "id",
	Short: "Show bacalhau node id info",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Cleanup manager ensures that resources are freed before exiting:
		cm := system.NewCleanupManager()
		cm.RegisterCallback(system.CleanupTraceProvider)
		defer cm.Cleanup()
		ctx := cmd.Context()

		transport, err := libp2p.NewTransport(ctx, cm, OS.SwarmPort, []multiaddr.Multiaddr{})
		if err != nil {
			return err
		}

		info := IdInfo{
			ID: transport.HostID(),
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "    ")
		if err := enc.Encode(info); err != nil {
			return err
		}

		return nil
	},
}
