package tests

import (
	"testing"

	"github.com/aerodinamicat/generalrestapi/models"
)

var testTransactions []*models.AccountingTransaction
var testRecords [][]*models.AccountingRecord

func FillTestTransactionVariables() {
	testTransactions = []*models.AccountingTransaction{{}, {}, {}, {}, {}, {}}
	testRecords = [][]*models.AccountingRecord{
		{
			{Nature: models.NATURE_CREDIT, Amount: 1},
			{Nature: models.NATURE_DEBIT, Amount: -2},
			{Nature: models.NATURE_CREDIT, Amount: 3},
			{Nature: models.NATURE_DEBIT, Amount: -4},
			{Nature: models.NATURE_CREDIT, Amount: 5},
		},
		{
			{Nature: models.NATURE_CREDIT, Amount: -6.66},
			{Nature: models.NATURE_CREDIT, Amount: 7.77},
			{Nature: models.NATURE_DEBIT, Amount: -8.88},
			{Nature: models.NATURE_DEBIT, Amount: 9.99},
			{Nature: models.NATURE_CREDIT, Amount: -10.10},
			{Nature: models.NATURE_CREDIT, Amount: 11.11},
			{Nature: models.NATURE_DEBIT, Amount: -12.12},
			{Nature: models.NATURE_DEBIT, Amount: 13.13},
			{Nature: models.NATURE_CREDIT, Amount: -14.14},
			{Nature: models.NATURE_CREDIT, Amount: 15.15},
		},
		{
			{Nature: models.NATURE_CREDIT, Amount: -16.16},
			{Nature: models.NATURE_CREDIT, Amount: 17.17},
			{Nature: models.NATURE_CREDIT, Amount: -18.18},
			{Nature: models.NATURE_DEBIT, Amount: 19.19},
			{Nature: models.NATURE_DEBIT, Amount: -20.20},
			{Nature: models.NATURE_DEBIT, Amount: 21.21},
			{Nature: models.NATURE_CREDIT, Amount: -22.22},
			{Nature: models.NATURE_CREDIT, Amount: 23.23},
			{Nature: models.NATURE_CREDIT, Amount: -24.24},
			{Nature: models.NATURE_DEBIT, Amount: 25.25},
			{Nature: models.NATURE_DEBIT, Amount: -26.26},
			{Nature: models.NATURE_DEBIT, Amount: 27.27},
			{Nature: models.NATURE_CREDIT, Amount: -28.28},
			{Nature: models.NATURE_CREDIT, Amount: 29.29},
			{Nature: models.NATURE_CREDIT, Amount: -30.30},
			{Nature: models.NATURE_DEBIT, Amount: 31.31},
			{Nature: models.NATURE_DEBIT, Amount: -32.32},
			{Nature: models.NATURE_DEBIT, Amount: 33.33},
			{Nature: models.NATURE_CREDIT, Amount: -34.34},
			{Nature: models.NATURE_CREDIT, Amount: 35.35},
		},
		{
			{Nature: models.NATURE_CREDIT, Amount: -36.36},
			{Nature: models.NATURE_CREDIT, Amount: 37.37},
			{Nature: models.NATURE_CREDIT, Amount: -38.38},
			{Nature: models.NATURE_CREDIT, Amount: 39.39},
			{Nature: models.NATURE_DEBIT, Amount: -40.40},
			{Nature: models.NATURE_DEBIT, Amount: 41.41},
			{Nature: models.NATURE_DEBIT, Amount: -42.42},
			{Nature: models.NATURE_DEBIT, Amount: 43.43},
			{Nature: models.NATURE_CREDIT, Amount: -44.44},
			{Nature: models.NATURE_CREDIT, Amount: 45.45},
			{Nature: models.NATURE_CREDIT, Amount: -46.46},
			{Nature: models.NATURE_CREDIT, Amount: 47.47},
			{Nature: models.NATURE_DEBIT, Amount: -48.48},
			{Nature: models.NATURE_DEBIT, Amount: 50.50},
			{Nature: models.NATURE_DEBIT, Amount: -51.51},
			{Nature: models.NATURE_DEBIT, Amount: 52.52},
			{Nature: models.NATURE_CREDIT, Amount: -53.53},
			{Nature: models.NATURE_CREDIT, Amount: 54.54},
			{Nature: models.NATURE_CREDIT, Amount: -55.55},
			{Nature: models.NATURE_CREDIT, Amount: 56.56},
			{Nature: models.NATURE_DEBIT, Amount: -57.57},
			{Nature: models.NATURE_DEBIT, Amount: 58.58},
			{Nature: models.NATURE_DEBIT, Amount: -59.59},
			{Nature: models.NATURE_DEBIT, Amount: 60.60},
			{Nature: models.NATURE_CREDIT, Amount: -61.61},
			{Nature: models.NATURE_CREDIT, Amount: 62.62},
			{Nature: models.NATURE_CREDIT, Amount: -63.63},
			{Nature: models.NATURE_CREDIT, Amount: 64.64},
			{Nature: models.NATURE_DEBIT, Amount: -65.65},
			{Nature: models.NATURE_DEBIT, Amount: 66.66},
			{Nature: models.NATURE_DEBIT, Amount: -67.67},
			{Nature: models.NATURE_DEBIT, Amount: 68.68},
			{Nature: models.NATURE_CREDIT, Amount: -69.69},
			{Nature: models.NATURE_CREDIT, Amount: 70.70},
			{Nature: models.NATURE_CREDIT, Amount: -71.71},
			{Nature: models.NATURE_CREDIT, Amount: 72.72},
			{Nature: models.NATURE_DEBIT, Amount: -73.73},
			{Nature: models.NATURE_DEBIT, Amount: 74.74},
			{Nature: models.NATURE_DEBIT, Amount: -75.75},
			{Nature: models.NATURE_DEBIT, Amount: 76.76},
		},
		{
			{Nature: models.NATURE_CREDIT, Amount: -77.77},
			{Nature: models.NATURE_CREDIT, Amount: 78.78},
			{Nature: models.NATURE_CREDIT, Amount: -79.79},
			{Nature: models.NATURE_CREDIT, Amount: 80.80},
			{Nature: models.NATURE_CREDIT, Amount: -81.81},
			{Nature: models.NATURE_DEBIT, Amount: 82.82},
			{Nature: models.NATURE_DEBIT, Amount: -83.83},
			{Nature: models.NATURE_DEBIT, Amount: 84.84},
			{Nature: models.NATURE_DEBIT, Amount: -85.85},
			{Nature: models.NATURE_DEBIT, Amount: 86.86},
			{Nature: models.NATURE_CREDIT, Amount: -87.87},
			{Nature: models.NATURE_CREDIT, Amount: 88.88},
			{Nature: models.NATURE_CREDIT, Amount: -89.89},
			{Nature: models.NATURE_CREDIT, Amount: 90.90},
			{Nature: models.NATURE_CREDIT, Amount: -91.91},
			{Nature: models.NATURE_DEBIT, Amount: 92.92},
			{Nature: models.NATURE_DEBIT, Amount: -93.93},
			{Nature: models.NATURE_DEBIT, Amount: 94.94},
			{Nature: models.NATURE_DEBIT, Amount: -95.95},
			{Nature: models.NATURE_DEBIT, Amount: 96.96},
			{Nature: models.NATURE_CREDIT, Amount: -97.97},
			{Nature: models.NATURE_CREDIT, Amount: 98.98},
			{Nature: models.NATURE_CREDIT, Amount: -99.99},
			{Nature: models.NATURE_CREDIT, Amount: 100.00},
			{Nature: models.NATURE_CREDIT, Amount: -101.01},
			{Nature: models.NATURE_DEBIT, Amount: 102.02},
			{Nature: models.NATURE_DEBIT, Amount: -103.03},
			{Nature: models.NATURE_DEBIT, Amount: 104.04},
			{Nature: models.NATURE_DEBIT, Amount: -105.05},
			{Nature: models.NATURE_DEBIT, Amount: 106.06},
			{Nature: models.NATURE_CREDIT, Amount: -107.07},
			{Nature: models.NATURE_CREDIT, Amount: 108.08},
			{Nature: models.NATURE_CREDIT, Amount: -109.09},
			{Nature: models.NATURE_CREDIT, Amount: 110.10},
			{Nature: models.NATURE_CREDIT, Amount: -111.11},
			{Nature: models.NATURE_DEBIT, Amount: 112.12},
			{Nature: models.NATURE_DEBIT, Amount: -113.13},
			{Nature: models.NATURE_DEBIT, Amount: 114.14},
			{Nature: models.NATURE_DEBIT, Amount: -115.15},
			{Nature: models.NATURE_DEBIT, Amount: 116.16},
			{Nature: models.NATURE_CREDIT, Amount: -117.17},
			{Nature: models.NATURE_CREDIT, Amount: 118.18},
			{Nature: models.NATURE_CREDIT, Amount: -119.19},
			{Nature: models.NATURE_CREDIT, Amount: 120.20},
			{Nature: models.NATURE_CREDIT, Amount: -121.21},
			{Nature: models.NATURE_DEBIT, Amount: 122.22},
			{Nature: models.NATURE_DEBIT, Amount: -123.23},
			{Nature: models.NATURE_DEBIT, Amount: 124.24},
			{Nature: models.NATURE_DEBIT, Amount: -125.25},
			{Nature: models.NATURE_DEBIT, Amount: 126.26},
			{Nature: models.NATURE_CREDIT, Amount: -127.27},
			{Nature: models.NATURE_CREDIT, Amount: 128.28},
			{Nature: models.NATURE_CREDIT, Amount: -129.29},
			{Nature: models.NATURE_CREDIT, Amount: 130.30},
			{Nature: models.NATURE_CREDIT, Amount: -131.31},
			{Nature: models.NATURE_DEBIT, Amount: 132.32},
			{Nature: models.NATURE_DEBIT, Amount: -133.33},
			{Nature: models.NATURE_DEBIT, Amount: 134.34},
			{Nature: models.NATURE_DEBIT, Amount: -135.35},
			{Nature: models.NATURE_DEBIT, Amount: 136.36},
			{Nature: models.NATURE_CREDIT, Amount: -137.37},
			{Nature: models.NATURE_CREDIT, Amount: 138.38},
			{Nature: models.NATURE_CREDIT, Amount: -139.39},
			{Nature: models.NATURE_CREDIT, Amount: 140.40},
			{Nature: models.NATURE_CREDIT, Amount: -141.41},
			{Nature: models.NATURE_DEBIT, Amount: 142.42},
			{Nature: models.NATURE_DEBIT, Amount: -143.43},
			{Nature: models.NATURE_DEBIT, Amount: 144.44},
			{Nature: models.NATURE_DEBIT, Amount: -145.45},
			{Nature: models.NATURE_DEBIT, Amount: 146.46},
			{Nature: models.NATURE_CREDIT, Amount: -147.47},
			{Nature: models.NATURE_CREDIT, Amount: 148.48},
			{Nature: models.NATURE_CREDIT, Amount: -149.49},
			{Nature: models.NATURE_CREDIT, Amount: 150.50},
			{Nature: models.NATURE_CREDIT, Amount: -151.51},
			{Nature: models.NATURE_DEBIT, Amount: 152.52},
			{Nature: models.NATURE_DEBIT, Amount: -153.53},
			{Nature: models.NATURE_DEBIT, Amount: 154.54},
			{Nature: models.NATURE_DEBIT, Amount: -155.55},
			{Nature: models.NATURE_DEBIT, Amount: 156.56},
		},
		{
			{Nature: models.NATURE_DEBIT, Amount: 100},
			{Nature: models.NATURE_DEBIT, Amount: 100},
			{Nature: models.NATURE_DEBIT, Amount: 100},
			{Nature: models.NATURE_DEBIT, Amount: 100},
			{Nature: models.NATURE_DEBIT, Amount: 100},
			{Nature: models.NATURE_CREDIT, Amount: 100},
			{Nature: models.NATURE_CREDIT, Amount: 100},
			{Nature: models.NATURE_CREDIT, Amount: 100},
			{Nature: models.NATURE_CREDIT, Amount: 100},
			{Nature: models.NATURE_CREDIT, Amount: 100},
		},
	}
}

func TestCalculateZeroSum(t *testing.T) {
	var testResultBoolean bool
	var testResultFloat float64
	var expectedResultBoolean bool
	var expectedResultFloat float64

	expectedResultsFromCalculateZeroSum := []struct {
		nulo     bool
		quantity float64
	}{
		{false, -15},
		{false, -1.01},
		{false, 147.46},
		{false, 1.01},
		{false, 1871.64},
		{true, 0},
	}

	FillTestTransactionVariables()

	for index, item := range testTransactions {
		testResultBoolean, testResultFloat = item.CalculateZeroSum(testRecords[index])
		expectedResultBoolean = expectedResultsFromCalculateZeroSum[index].nulo
		expectedResultFloat = expectedResultsFromCalculateZeroSum[index].quantity

		if expectedResultBoolean != testResultBoolean || expectedResultFloat != testResultFloat {
			t.Errorf("CalculateZeroSum was incorrect,\ngot:\n\t'%t'\t and '%f'\nexpected\n\t'%t'\t and '%f'", testResultBoolean, testResultFloat, expectedResultBoolean, expectedResultFloat)
		}
	}
}
