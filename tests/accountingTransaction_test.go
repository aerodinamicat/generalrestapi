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
			{Type: models.TYPE_CREDIT, Amount: 1},
			{Type: models.TYPE_DEBIT, Amount: -2},
			{Type: models.TYPE_CREDIT, Amount: 3},
			{Type: models.TYPE_DEBIT, Amount: -4},
			{Type: models.TYPE_CREDIT, Amount: 5},
		},
		{
			{Type: models.TYPE_CREDIT, Amount: -6.66},
			{Type: models.TYPE_CREDIT, Amount: 7.77},
			{Type: models.TYPE_DEBIT, Amount: -8.88},
			{Type: models.TYPE_DEBIT, Amount: 9.99},
			{Type: models.TYPE_CREDIT, Amount: -10.10},
			{Type: models.TYPE_CREDIT, Amount: 11.11},
			{Type: models.TYPE_DEBIT, Amount: -12.12},
			{Type: models.TYPE_DEBIT, Amount: 13.13},
			{Type: models.TYPE_CREDIT, Amount: -14.14},
			{Type: models.TYPE_CREDIT, Amount: 15.15},
		},
		{
			{Type: models.TYPE_CREDIT, Amount: -16.16},
			{Type: models.TYPE_CREDIT, Amount: 17.17},
			{Type: models.TYPE_CREDIT, Amount: -18.18},
			{Type: models.TYPE_DEBIT, Amount: 19.19},
			{Type: models.TYPE_DEBIT, Amount: -20.20},
			{Type: models.TYPE_DEBIT, Amount: 21.21},
			{Type: models.TYPE_CREDIT, Amount: -22.22},
			{Type: models.TYPE_CREDIT, Amount: 23.23},
			{Type: models.TYPE_CREDIT, Amount: -24.24},
			{Type: models.TYPE_DEBIT, Amount: 25.25},
			{Type: models.TYPE_DEBIT, Amount: -26.26},
			{Type: models.TYPE_DEBIT, Amount: 27.27},
			{Type: models.TYPE_CREDIT, Amount: -28.28},
			{Type: models.TYPE_CREDIT, Amount: 29.29},
			{Type: models.TYPE_CREDIT, Amount: -30.30},
			{Type: models.TYPE_DEBIT, Amount: 31.31},
			{Type: models.TYPE_DEBIT, Amount: -32.32},
			{Type: models.TYPE_DEBIT, Amount: 33.33},
			{Type: models.TYPE_CREDIT, Amount: -34.34},
			{Type: models.TYPE_CREDIT, Amount: 35.35},
		},
		{
			{Type: models.TYPE_CREDIT, Amount: -36.36},
			{Type: models.TYPE_CREDIT, Amount: 37.37},
			{Type: models.TYPE_CREDIT, Amount: -38.38},
			{Type: models.TYPE_CREDIT, Amount: 39.39},
			{Type: models.TYPE_DEBIT, Amount: -40.40},
			{Type: models.TYPE_DEBIT, Amount: 41.41},
			{Type: models.TYPE_DEBIT, Amount: -42.42},
			{Type: models.TYPE_DEBIT, Amount: 43.43},
			{Type: models.TYPE_CREDIT, Amount: -44.44},
			{Type: models.TYPE_CREDIT, Amount: 45.45},
			{Type: models.TYPE_CREDIT, Amount: -46.46},
			{Type: models.TYPE_CREDIT, Amount: 47.47},
			{Type: models.TYPE_DEBIT, Amount: -48.48},
			{Type: models.TYPE_DEBIT, Amount: 50.50},
			{Type: models.TYPE_DEBIT, Amount: -51.51},
			{Type: models.TYPE_DEBIT, Amount: 52.52},
			{Type: models.TYPE_CREDIT, Amount: -53.53},
			{Type: models.TYPE_CREDIT, Amount: 54.54},
			{Type: models.TYPE_CREDIT, Amount: -55.55},
			{Type: models.TYPE_CREDIT, Amount: 56.56},
			{Type: models.TYPE_DEBIT, Amount: -57.57},
			{Type: models.TYPE_DEBIT, Amount: 58.58},
			{Type: models.TYPE_DEBIT, Amount: -59.59},
			{Type: models.TYPE_DEBIT, Amount: 60.60},
			{Type: models.TYPE_CREDIT, Amount: -61.61},
			{Type: models.TYPE_CREDIT, Amount: 62.62},
			{Type: models.TYPE_CREDIT, Amount: -63.63},
			{Type: models.TYPE_CREDIT, Amount: 64.64},
			{Type: models.TYPE_DEBIT, Amount: -65.65},
			{Type: models.TYPE_DEBIT, Amount: 66.66},
			{Type: models.TYPE_DEBIT, Amount: -67.67},
			{Type: models.TYPE_DEBIT, Amount: 68.68},
			{Type: models.TYPE_CREDIT, Amount: -69.69},
			{Type: models.TYPE_CREDIT, Amount: 70.70},
			{Type: models.TYPE_CREDIT, Amount: -71.71},
			{Type: models.TYPE_CREDIT, Amount: 72.72},
			{Type: models.TYPE_DEBIT, Amount: -73.73},
			{Type: models.TYPE_DEBIT, Amount: 74.74},
			{Type: models.TYPE_DEBIT, Amount: -75.75},
			{Type: models.TYPE_DEBIT, Amount: 76.76},
		},
		{
			{Type: models.TYPE_CREDIT, Amount: -77.77},
			{Type: models.TYPE_CREDIT, Amount: 78.78},
			{Type: models.TYPE_CREDIT, Amount: -79.79},
			{Type: models.TYPE_CREDIT, Amount: 80.80},
			{Type: models.TYPE_CREDIT, Amount: -81.81},
			{Type: models.TYPE_DEBIT, Amount: 82.82},
			{Type: models.TYPE_DEBIT, Amount: -83.83},
			{Type: models.TYPE_DEBIT, Amount: 84.84},
			{Type: models.TYPE_DEBIT, Amount: -85.85},
			{Type: models.TYPE_DEBIT, Amount: 86.86},
			{Type: models.TYPE_CREDIT, Amount: -87.87},
			{Type: models.TYPE_CREDIT, Amount: 88.88},
			{Type: models.TYPE_CREDIT, Amount: -89.89},
			{Type: models.TYPE_CREDIT, Amount: 90.90},
			{Type: models.TYPE_CREDIT, Amount: -91.91},
			{Type: models.TYPE_DEBIT, Amount: 92.92},
			{Type: models.TYPE_DEBIT, Amount: -93.93},
			{Type: models.TYPE_DEBIT, Amount: 94.94},
			{Type: models.TYPE_DEBIT, Amount: -95.95},
			{Type: models.TYPE_DEBIT, Amount: 96.96},
			{Type: models.TYPE_CREDIT, Amount: -97.97},
			{Type: models.TYPE_CREDIT, Amount: 98.98},
			{Type: models.TYPE_CREDIT, Amount: -99.99},
			{Type: models.TYPE_CREDIT, Amount: 100.00},
			{Type: models.TYPE_CREDIT, Amount: -101.01},
			{Type: models.TYPE_DEBIT, Amount: 102.02},
			{Type: models.TYPE_DEBIT, Amount: -103.03},
			{Type: models.TYPE_DEBIT, Amount: 104.04},
			{Type: models.TYPE_DEBIT, Amount: -105.05},
			{Type: models.TYPE_DEBIT, Amount: 106.06},
			{Type: models.TYPE_CREDIT, Amount: -107.07},
			{Type: models.TYPE_CREDIT, Amount: 108.08},
			{Type: models.TYPE_CREDIT, Amount: -109.09},
			{Type: models.TYPE_CREDIT, Amount: 110.10},
			{Type: models.TYPE_CREDIT, Amount: -111.11},
			{Type: models.TYPE_DEBIT, Amount: 112.12},
			{Type: models.TYPE_DEBIT, Amount: -113.13},
			{Type: models.TYPE_DEBIT, Amount: 114.14},
			{Type: models.TYPE_DEBIT, Amount: -115.15},
			{Type: models.TYPE_DEBIT, Amount: 116.16},
			{Type: models.TYPE_CREDIT, Amount: -117.17},
			{Type: models.TYPE_CREDIT, Amount: 118.18},
			{Type: models.TYPE_CREDIT, Amount: -119.19},
			{Type: models.TYPE_CREDIT, Amount: 120.20},
			{Type: models.TYPE_CREDIT, Amount: -121.21},
			{Type: models.TYPE_DEBIT, Amount: 122.22},
			{Type: models.TYPE_DEBIT, Amount: -123.23},
			{Type: models.TYPE_DEBIT, Amount: 124.24},
			{Type: models.TYPE_DEBIT, Amount: -125.25},
			{Type: models.TYPE_DEBIT, Amount: 126.26},
			{Type: models.TYPE_CREDIT, Amount: -127.27},
			{Type: models.TYPE_CREDIT, Amount: 128.28},
			{Type: models.TYPE_CREDIT, Amount: -129.29},
			{Type: models.TYPE_CREDIT, Amount: 130.30},
			{Type: models.TYPE_CREDIT, Amount: -131.31},
			{Type: models.TYPE_DEBIT, Amount: 132.32},
			{Type: models.TYPE_DEBIT, Amount: -133.33},
			{Type: models.TYPE_DEBIT, Amount: 134.34},
			{Type: models.TYPE_DEBIT, Amount: -135.35},
			{Type: models.TYPE_DEBIT, Amount: 136.36},
			{Type: models.TYPE_CREDIT, Amount: -137.37},
			{Type: models.TYPE_CREDIT, Amount: 138.38},
			{Type: models.TYPE_CREDIT, Amount: -139.39},
			{Type: models.TYPE_CREDIT, Amount: 140.40},
			{Type: models.TYPE_CREDIT, Amount: -141.41},
			{Type: models.TYPE_DEBIT, Amount: 142.42},
			{Type: models.TYPE_DEBIT, Amount: -143.43},
			{Type: models.TYPE_DEBIT, Amount: 144.44},
			{Type: models.TYPE_DEBIT, Amount: -145.45},
			{Type: models.TYPE_DEBIT, Amount: 146.46},
			{Type: models.TYPE_CREDIT, Amount: -147.47},
			{Type: models.TYPE_CREDIT, Amount: 148.48},
			{Type: models.TYPE_CREDIT, Amount: -149.49},
			{Type: models.TYPE_CREDIT, Amount: 150.50},
			{Type: models.TYPE_CREDIT, Amount: -151.51},
			{Type: models.TYPE_DEBIT, Amount: 152.52},
			{Type: models.TYPE_DEBIT, Amount: -153.53},
			{Type: models.TYPE_DEBIT, Amount: 154.54},
			{Type: models.TYPE_DEBIT, Amount: -155.55},
			{Type: models.TYPE_DEBIT, Amount: 156.56},
		},
		{
			{Type: models.TYPE_DEBIT, Amount: 100},
			{Type: models.TYPE_DEBIT, Amount: 100},
			{Type: models.TYPE_DEBIT, Amount: 100},
			{Type: models.TYPE_DEBIT, Amount: 100},
			{Type: models.TYPE_DEBIT, Amount: 100},
			{Type: models.TYPE_CREDIT, Amount: 100},
			{Type: models.TYPE_CREDIT, Amount: 100},
			{Type: models.TYPE_CREDIT, Amount: 100},
			{Type: models.TYPE_CREDIT, Amount: 100},
			{Type: models.TYPE_CREDIT, Amount: 100},
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
