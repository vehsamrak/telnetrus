package telnetrus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dataset = []struct {
	input  []byte
	output string
}{
	{[]byte{48, 49}, "01"},
	{[]byte{50, 51, 52}, "234"},
}

func TestTelnetrus(test *testing.T) {
	suite.Run(test, new(telnetrusTest))
}

type telnetrusTest struct {
	suite.Suite
}

func (suite *telnetrusTest) Test_Process_bytesWithYaLetter_correctString() {
	for id, dataset := range dataset {
		result := Process(dataset.input)

		assert.Equal(suite.T(), dataset.output, result, fmt.Sprintf("Dataset #%d", id), dataset.input, dataset.output)
	}
}
