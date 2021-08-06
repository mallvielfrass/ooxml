package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokens(t *testing.T) {
	testingXML := `<w:r d='33'><w:rPr></w:rPr><w:t>This is a</w:t></w:r>`
	var tokens []NewToken
	tokens = append(tokens,
		NewToken{
			Name:           "w:r",
			TagStatus:      TagOpenComplimentary,
			Args:           "d='33'",
			TokenSymbOpen:  0,
			TokenSymbClose: 11,
		},
		NewToken{
			Name:           "w:rPr",
			TagStatus:      TagOpenComplimentary,
			Args:           "",
			TokenSymbOpen:  12,
			TokenSymbClose: 18,
		},
		NewToken{
			Name:           "w:rPr",
			TagStatus:      TagClosingComplimentary,
			Args:           "",
			TokenSymbOpen:  19,
			TokenSymbClose: 26,
		},
		NewToken{
			Name:           "w:t",
			TagStatus:      TagOpenComplimentary,
			Args:           "",
			TokenSymbOpen:  27,
			TokenSymbClose: 31,
		},
		NewToken{
			Name:           "w:t",
			TagStatus:      TagClosingComplimentary,
			Args:           "",
			TokenSymbOpen:  41,
			TokenSymbClose: 46,
		},
		NewToken{
			Name:           "w:r",
			TagStatus:      TagClosingComplimentary,
			Args:           "",
			TokenSymbOpen:  47,
			TokenSymbClose: 52,
		})
	parsed := getTokens(testingXML)
	if len(parsed) <= 0 {
		t.Error("Expected nodes")
		return
	}
	if len(parsed) != len(tokens) {
		t.Error("wrong array size")
		return
	}
	for i, item := range parsed {
		assert.Equal(t, item, tokens[i], "they should be equal")
	}
}

// func TestGetFirstNodes(t *testing.T) {
// 	testingXML := `
// 					<div/>
// 					<div></div>
// 					<w:r d='33'>
// 						<w:rPr>
// 						</w:rPr>
// 						<w:t>
// 							This is a
// 						</w:t>
// 					</w:r>`
// 	tokens := getTokens(testingXML)
// 	fmt.Println(getFirstNodes(tokens))

// }
func TestGetParentNodes(t *testing.T) {
	testingXML := `
					<div/>
					<div></div>
					<w:r d='33'>
						<w:rPr>
						</w:rPr>
						<w:t>
							This is a
						</w:t>
					</w:r>`
	bodyWR := `
						<w:rPr>
						</w:rPr>
						<w:t>
							This is a
						</w:t>
					`
	var tokens []EmbeddedToken
	tokens = append(tokens,
		EmbeddedToken{
			Name:      "div",
			TagStatus: TagSelfClosed,
			Body:      "",
			Args:      "",
		},
		EmbeddedToken{
			Name:      "div",
			TagStatus: TagComplimentary,
			Body:      "",
			Args:      "",
		},
		EmbeddedToken{
			Name:      "w:r",
			TagStatus: TagComplimentary,
			Body:      bodyWR,
			Args:      "d='33'",
		},
	)
	nodes := getParentNodes(testingXML)
	for i, item := range nodes {
		assert.Equal(t, item, tokens[i], "they should be equal")
	}

}
