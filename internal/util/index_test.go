package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCaseSnake2Camel(t *testing.T) {
	require.Equal(t, "Asd", CaseSnake2Camel("asd"))
	require.Equal(t, "AsdDsa", CaseSnake2Camel("asd_dsa"))
}

func TestCaseCamel2Snake(t *testing.T) {
	require.Equal(t, "asd", CaseCamel2Snake("Asd"))
	require.Equal(t, "asd_dsa_qwe", CaseCamel2Snake("AsdDsaQwe"))
	require.Equal(t, "jsonparser", CaseCamel2Snake("JSONParser"))
}
