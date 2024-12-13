// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaValidateCommand(t *testing.T) {
	stdout, stderr, err := e2e.UDSPK("schema", "validate", "-b", "src/test/values_schemas/04_validate_schema")
	require.NoError(t, err, stdout, stderr)

	require.Contains(t, stderr, "All schemas match.")
}

func TestSchemaValidateFailCommand(t *testing.T) {
	stdout, stderr, err := e2e.UDSPK("schema", "validate", "-b", "src/test/values_schemas/04_validate_schema_fail")
	require.NoError(t, err, stdout, stderr)
	require.Contains(t, stderr, "Schema differences found.")
}
