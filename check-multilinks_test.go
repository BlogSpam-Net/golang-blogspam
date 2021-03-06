//
// Test for our link-method plugin.
//

package main

import (
	"testing"
)

//
// Test that we can accept random linking-methods.
//
func TestLinkTypesOK(t *testing.T) {

	//
	//  Each type alone is OK
	//
	input := []string{"<a href=\"https://example.com\">test</a>",
		"[url=https://exmapl.eocm]title[/url]",
		"[link=https://moi.com/]steve[/link]:",
		"[ \t]https?:/"}

	//
	// Try them all
	//
	for _, i := range input {
		//
		// Test a simple comment.
		//
		result, detail := checkLinkingTypes(Submission{Comment: i})
		if result != Undecided {
			t.Errorf("Unexpected response: '%v'", result)
		}
		if len(detail) != 0 {
			t.Errorf("Unexpected response: '%v'", detail)
		}
	}
}

//
// Test that all methods are too much.
//
func TestLinkTypesSpam(t *testing.T) {

	//
	//  Try all three
	//
	input := "<a href=\"https://steve.fi/\">Steve Kemp</a>, [url=http://moi.kisssa]Finnihs[/url]  http://bare.link.com/"

	result, detail := checkLinkingTypes(Submission{Comment: input})
	if result != Spam {
		t.Errorf("Unexpected response: '%v'", result)
	}
	if len(detail) == 0 {
		t.Errorf("Unexpected response: '%v'", detail)
	}
}
