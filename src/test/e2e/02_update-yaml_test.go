package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	uds "github.com/defenseunicorns/uds-cli/src/types"
	zarf "github.com/zarf-dev/zarf/src/api/v1alpha1"
)

func TestUpdateYamlCommand(t *testing.T) {
	e2e.CreateSandboxDir(t, "bundle")
	defer e2e.CleanupSandboxDir(t)

	// Create a dummy zarf yaml with devel as version
	e2e.CreateZarfYaml(t, "src/test/sandbox")
	// Create a dummy uds-bundle yaml with devel as version
	e2e.CreateUDSBundleYaml(t, "src/test/sandbox/bundle")

	stdout, stderr, err := e2e.UDSReleaserDir("src/test/sandbox", "update-yaml", "base", "-d", "../")
	require.NoError(t, err, stdout, stderr)

	// Check that the zarf.yaml was updated
	var zarfPackage zarf.ZarfPackage
	err = e2e.LoadYaml("src/test/sandbox/zarf.yaml", &zarfPackage)
	require.NoError(t, err)

	require.Equal(t, "1.0.0-uds.0", zarfPackage.Metadata.Version)

	// Check that the uds-bundle.yaml was updated
	var bundle uds.UDSBundle
	err = e2e.LoadYaml("src/test/sandbox/bundle/uds-bundle.yaml", &bundle)
	require.NoError(t, err)

	require.Equal(t, "1.0.0-uds.0", bundle.Metadata.Version)
	require.Equal(t, "1.0.0-uds.0", bundle.Packages[0].Ref)
}
