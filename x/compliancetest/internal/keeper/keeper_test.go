package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_TestingResultGetSet(t *testing.T) {
	setup := Setup()

	// check if testing result present
	require.False(t, setup.CompliancetestKeeper.IsTestingResultsPresents(setup.Ctx, test_constants.VID, test_constants.PID))
	require.False(t, setup.CompliancetestKeeper.IsTestingResultPresents(setup.Ctx, test_constants.VID, test_constants.PID, test_constants.Owner))

	// empty testing result before its created
	receivedTestingResult := setup.CompliancetestKeeper.GetTestingResults(setup.Ctx, test_constants.VID, test_constants.PID)
	require.Equal(t, 0, len(receivedTestingResult.Results))

	// create testing results
	testingResult := DefaultTestingResult()
	result := setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, testingResult)
	require.Equal(t, sdk.CodeOK, result.Code)

	// check if testing result present
	require.True(t, setup.CompliancetestKeeper.IsTestingResultsPresents(setup.Ctx, test_constants.VID, test_constants.PID))
	require.True(t, setup.CompliancetestKeeper.IsTestingResultPresents(setup.Ctx, test_constants.VID, test_constants.PID, test_constants.Owner))

	// get testing results
	receivedTestingResult = setup.CompliancetestKeeper.GetTestingResults(setup.Ctx, test_constants.VID, test_constants.PID)
	require.Equal(t, test_constants.VID, receivedTestingResult.VID)
	require.Equal(t, test_constants.PID, receivedTestingResult.PID)
	require.Equal(t, 1, len(receivedTestingResult.Results))
	firstItem := types.NewTestingResultItem(testingResult.TestResult, testingResult.Owner)
	require.Equal(t, receivedTestingResult.Results[0], firstItem)

	// add second testing result
	secondTestingResult := DefaultTestingResult()
	secondTestingResult.Owner = test_constants.Address2
	secondTestingResult.TestResult = utils.RandString()
	result = setup.CompliancetestKeeper.AddTestingResult(setup.Ctx, secondTestingResult)
	require.Equal(t, sdk.CodeOK, result.Code)

	// get testing results
	receivedTestingResult = setup.CompliancetestKeeper.GetTestingResults(setup.Ctx, test_constants.VID, test_constants.PID)
	require.Equal(t, test_constants.VID, receivedTestingResult.VID)
	require.Equal(t, test_constants.PID, receivedTestingResult.PID)
	require.Equal(t, 2, len(receivedTestingResult.Results))
	require.Equal(t, receivedTestingResult.Results[0], firstItem)
	secondItem := types.NewTestingResultItem(secondTestingResult.TestResult, secondTestingResult.Owner)
	require.Equal(t, receivedTestingResult.Results[1], secondItem)
}
