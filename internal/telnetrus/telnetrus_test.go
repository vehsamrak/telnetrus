package telnetrus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var encodedDataset = []struct {
	encoding string
	input    []byte
	output   string
}{
	{Windows1251, []byte{0xff}, "я"},
	{Windows1251, []byte{0xff, 0xff}, "яя"},
	{Windows1251, []byte{0xff, 0xff, 0xff}, "яяя"},
	{KOI8, []byte{0xd1}, "я"},
	{KOI8, []byte{0xd1, 0xd1}, "яя"},
	{KOI8, []byte{0xd1, 0xd1, 0xd1}, "яяя"},
}

var utfDataset = []struct {
	encoding string
	input    string
	output   []byte
}{
	{Windows1251, "я", []byte{0xff, 0xff}},
	{Windows1251, "яя", []byte{0xff, 0xff, 0xff, 0xff}},
	{Windows1251, "яяя", []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	{KOI8, "я", []byte{0xd1}},
	{KOI8, "яя", []byte{0xd1, 0xd1}},
	{KOI8, "яяя", []byte{0xd1, 0xd1, 0xd1}},
}

func TestTelnetrus(test *testing.T) {
	suite.Run(test, new(telnetrusTest))
}

type telnetrusTest struct {
	suite.Suite
}

func (suite *telnetrusTest) Test_ToUTF8_EncodingAndEncodedBytes_correctEncodedString() {
	for id, dataset := range encodedDataset {
		result, _ := ToUTF8(dataset.encoding, dataset.input)

		assert.Equal(suite.T(), dataset.output, result, fmt.Sprintf("Dataset #%d", id), dataset.input, dataset.output)
	}
}

func (suite *telnetrusTest) Test_FromUTF8_EncodingAndUTF8String_correctUTF8String() {
	for id, dataset := range utfDataset {
		result, _ := FromUTF8(dataset.encoding, dataset.input)

		assert.Equal(suite.T(), dataset.output, result, fmt.Sprintf("Dataset #%d", id), dataset.input, dataset.output)
	}
}
