// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaGenerateCommand(t *testing.T) {
	stdout, stderr, err := e2e.UDSPK("schema", "generate", "-b", "src/test/values_schemas/03_generate_schema")
	require.NoError(t, err, stdout, stderr)

	require.Contains(t, stderr, "Schema generated at")

	stdout, stderr, err = e2e.UDSPK("schema", "validate", "-b", "src/test/values_schemas/03_generate_schema")
	require.NoError(t, err, stdout, stderr)

	require.Contains(t, stderr, "All schemas match.")
}
