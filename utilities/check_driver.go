package utilities

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// CheckMainPuzzleSolution can be used to test that a main program generates
// the correct puzzle output.
func CheckMainPuzzleSolution(t *testing.T, expect1, expect2 int) {
	// Run the puzzle driver.
	cmd := exec.Command("go", "run", "main.go")

	// Fetch the output written to the console.
	raw, err := cmd.CombinedOutput()
	require.NoError(t, err)
	output := string(raw)

	// Write puzzle output to the console via test object. Will only be shown
	// if the test fails. Leading newline for nice formatting.
	t.Log(string(linefeed) + output)

	// The test output should only have two lines, if nothing fails.
	output = strings.TrimRight(output, string(linefeed)) // remove trailing \n
	lines := strings.Split(output, string(linefeed))
	require.Len(t, lines, 2, "puzzle driver ought to only write two lines")

	// Check each puzzle solution to make sure it is correct.
	var ok int
	var actual1, actual2 int

	ok, err = fmt.Sscanf(lines[0], "Part one answer is: %d", &actual1)
	require.NoError(t, err, "failed to scan part one")
	require.Equal(t, 1, ok, "no answer value for part one could be scanned")
	assert.Equal(t, expect1, actual1, "part one answer is incorrect")

	ok, err = fmt.Sscanf(lines[1], "Part two answer is: %d", &actual2)
	require.NoError(t, err, "failed to scan part two")
	require.Equal(t, 1, ok, "no answer value for part two could be scanned")
	assert.Equal(t, expect2, actual2, "part two answer is incorrect")
}
