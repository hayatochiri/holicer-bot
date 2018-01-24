package holicerBot

import (
	"fmt"
	"testing"
)

func (actual Tavern) compare(t *testing.T, expect Tavern) {
	output := "Compare\n"
	output += "\tID\n"
	output += fmt.Sprintf("\t\tactual : %d\n", actual.Id)
	output += fmt.Sprintf("\t\texpect : %d\n", expect.Id)
	output += "\tNameJA\n"
	output += fmt.Sprintf("\t\tactual : \"%s\"\n", actual.NameJA)
	output += fmt.Sprintf("\t\texpect : \"%s\"\n", expect.NameJA)
	output += "\tNameEN\n"
	output += fmt.Sprintf("\t\tactual : \"%s\"\n", actual.NameEN)
	output += fmt.Sprintf("\t\texpect : \"%s\"\n", expect.NameEN)
	output += "\tIsRemoved\n"
	output += "\t\tactual : " + TO_BOOL[actual.IsRemoved] + "\n"
	output += "\t\texpect : " + TO_BOOL[expect.IsRemoved] + "\n"

	if actual.Id == expect.Id &&
		actual.NameJA == expect.NameJA &&
		actual.NameEN == expect.NameEN &&
		actual.IsRemoved == expect.IsRemoved {
		t.Logf(output)
	} else {
		t.Errorf(output)
	}
}
