package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCaseSnake2Camel(t *testing.T) {
	require.Equal(t, "Asd", Case2Camel("asd", false))
	require.Equal(t, "AsdDsa", Case2Camel("asd_dsa", false))
	require.Equal(t, "AsdDsa", Case2Camel("AsdDsa", false))
	require.Equal(t, "asdDsa", Case2Camel("AsdDsa", true))
	require.Equal(t, "asdDsa", Case2Camel("asd_dsa", true))
	require.Equal(t, "asd", Case2Camel("asd", true))
}

func TestCaseCamel2Snake(t *testing.T) {
	require.Equal(t, "asd", Case2Snake("Asd"))
	require.Equal(t, "asd_dsa_qwe", Case2Snake("AsdDsaQwe"))
	require.Equal(t, "jsonparser", Case2Snake("JSONParser"))
	require.Equal(t, "asd_dsa_qwe", Case2Snake("asd_dsa_qwe"))
}
