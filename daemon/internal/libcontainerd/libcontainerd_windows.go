package libcontainerd

import (
	"context"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/docker/docker/daemon/internal/libcontainerd/local"
	"github.com/docker/docker/daemon/internal/libcontainerd/remote"
	libcontainerdtypes "github.com/docker/docker/daemon/internal/libcontainerd/types"
	"github.com/docker/docker/pkg/system"
)

// NewClient creates a new libcontainerd client from a containerd client
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend) (libcontainerdtypes.Client, error) {
	if !system.ContainerdRuntimeSupported() {
		return local.NewClient(ctx, b)
	}
	return remote.NewClient(ctx, cli, stateDir, ns, b)
}
