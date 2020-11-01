package entity

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseTagJson(t *testing.T) {
	require.Equal(t, "", ParseTagJson(""))
	require.Equal(t, "", ParseTagJson("`json:\"-\"`"))
	require.Equal(t, "name", ParseTagJson("`json:\"name\"`"))
	require.Equal(t, "name", ParseTagJson("`json:\"name,omitempty\"`"))
}
