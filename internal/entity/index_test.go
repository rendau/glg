package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTagJson(t *testing.T) {
	require.Equal(t, "", TagParseJsonName(""))
	require.Equal(t, "", TagParseJsonName("`json:\"-\"`"))
	require.Equal(t, "name", TagParseJsonName("`json:\"name\"`"))
	require.Equal(t, "name", TagParseJsonName("`json:\"name,omitempty\"`"))
}
