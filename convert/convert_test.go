package convert_test

import (
	"testing"

	. "github.com/alexdreptu/sysinfo/convert"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	value    float64
	expected float64
}

func TestBytesToKibibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1024, expected: 1},
		{value: 1331.2, expected: 1.3},
		{value: 1740.8, expected: 1.7},
		{value: 24268.8, expected: 23.7},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Byte).ToKibibytes())
	}
}

func TestBytesToMebibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1048576, expected: 1},
		{value: 1363148.8, expected: 1.3},
		{value: 1782579.2, expected: 1.7},
		{value: 24851251.2, expected: 23.7},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Byte).ToMebibytes())
	}
}

func TestBytesToGibibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1073741824, expected: 1},
		{value: 1395864371.2, expected: 1.3},
		{value: 1825361100.8, expected: 1.7},
		{value: 25447681228.8, expected: 23.7},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Byte).ToGibibytes())
	}
}

func TestKibibytesToBytes(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1024},
		{value: 1.3, expected: 1331.2},
		{value: 1.7, expected: 1740.8},
		{value: 23.7, expected: 24268.8},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Kibibyte).ToBytes())
	}
}

func TestKibibytesToMebibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1024, expected: 1},
		{value: 1331.2, expected: 1.3},
		{value: 1740.8, expected: 1.7},
		{value: 24268.8, expected: 23.7},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Kibibyte).ToMebibytes())
	}
}

func TestKibibytesToGibibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1048576, expected: 1},
		{value: 1363148.8, expected: 1.3},
		{value: 1782579.2, expected: 1.7},
		{value: 24851251.2, expected: 23.7},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Kibibyte).ToGibibytes())
	}
}

func TestMebibytesToBytes(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1048576},
		{value: 1.3, expected: 1363148.8},
		{value: 1.7, expected: 1782579.2},
		{value: 23.7, expected: 24851251.2},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Mebibyte).ToBytes())
	}
}

func TestMebibytesToKibibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1024},
		{value: 1.3, expected: 1331.2},
		{value: 1.7, expected: 1740.8},
		{value: 23.7, expected: 24268.8},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Mebibyte).ToKibibytes())
	}
}

func TestMebibytesToGibibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1024, expected: 1},
		{value: 1331.2, expected: 1.3},
		{value: 1740.8, expected: 1.7},
		{value: 24268.8, expected: 23.7},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Mebibyte).ToGibibytes())
	}
}

func TestGibibytesToBytes(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1073741824},
		{value: 1.3, expected: 1395864371.2},
		{value: 1.7, expected: 1825361100.8},
		{value: 23.7, expected: 25447681228.8},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Gibibyte).ToBytes())
	}
}

func TestGibibytesToKibibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1048576},
		{value: 1.3, expected: 1363148.8},
		{value: 1.7, expected: 1782579.2},
		{value: 23.7, expected: 24851251.2},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Gibibyte).ToKibibytes())
	}
}

func TestGibibytesToMebibytes(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1024},
		{value: 1.3, expected: 1331.2},
		{value: 1.7, expected: 1740.8},
		{value: 23.7, expected: 24268.8},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Size(tc.value) * Gibibyte).ToMebibytes())
	}
}

func TestMegahertzToGigahertz(t *testing.T) {
	testCases := []testCase{
		{value: 1000, expected: 1},
		{value: 1300, expected: 1.3},
		{value: 1700, expected: 1.7},
		{value: 23700, expected: 23.7},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Frequency(tc.value) * Megahertz).ToGigahertz())
	}
}

func TestGigahertzToKHertz(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1000000000},
		{value: 1.3, expected: 1300000000},
		{value: 1.7, expected: 1700000000},
		{value: 23.7, expected: 23700000000},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Frequency(tc.value) * Gigahertz).ToHertz())
	}
}

func TestGigahertzToKilohertz(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1000000},
		{value: 1.3, expected: 1300000},
		{value: 1.7, expected: 1700000},
		{value: 23.7, expected: 23700000},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Frequency(tc.value) * Gigahertz).ToKilohertz())
	}
}

func TestGigahertzToMegahertz(t *testing.T) {
	testCases := []testCase{
		{value: 1, expected: 1000},
		{value: 1.3, expected: 1300},
		{value: 1.7, expected: 1700},
		{value: 23.7, expected: 23700},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, (Frequency(tc.value) * Gigahertz).ToMegahertz())
	}
}
