package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCaseSnake2Camel(t *testing.T) {
	require.Equal(t, "Asd", Case2Camel("asd"))
	require.Equal(t, "AsdDsa", Case2Camel("asd_dsa"))
	require.Equal(t, "AsdDsa", Case2Camel("AsdDsa"))
}

func TestCaseCamel2Snake(t *testing.T) {
	require.Equal(t, "asd", Case2Snake("Asd"))
	require.Equal(t, "asd_dsa_qwe", Case2Snake("AsdDsaQwe"))
	require.Equal(t, "jsonparser", Case2Snake("JSONParser"))
	require.Equal(t, "asd_dsa_qwe", Case2Snake("asd_dsa_qwe"))
}
