package ooxml

import (
	"fmt"
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
func TestGetParentNodesT(t *testing.T) {
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
	nodes, err := GetParentNodes(testingXML)
	if err != nil {
		t.Error(err)
		return
	}
	if len(nodes) <= 0 {
		t.Error("Expected nodes")
		return
	}
	if len(nodes) != len(tokens) {
		t.Error("wrong array size")
		return
	}
	//fmt.Println(nodes[0])
	for i, item := range nodes {
		assert.Equal(t, item, tokens[i], "they should be equal")
	}

}
func TestGetParentNodesBrokenCases(t *testing.T) {
	testXML := []string{
		"",
		"<div</div>",
		"<zzz></w:rPr><w:t>This is a</w:t>",
		"<w:rPr/></w:rPr><w:t>This is a<w:t>",
		"<div>",
		"div",
		`<elm1 attr="value"> text </elm1> <elm2 attr="value" attr="value">text</root>`,
		"</div><div><div><div><div>",
	}
	for _, item := range testXML {
		//	fmt.Println("range")
		nodes, err := GetParentNodes(item)
		if len(nodes) != 0 {
			t.Errorf("node size expected 0, but %d in test [%s] ", len(nodes), item)
			fmt.Println(nodes)
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
			}
			return
		}
		if err == nil {
			t.Errorf("must be err [%v]", item)
			return
		}
	}

}
func TestParseRPR(t *testing.T) {
	xml := `<w:rPr>
                    <w:b/>
                    <w:b/>
                    <w:bCs/>
                    <w:i/>
                    <w:i/>
                    <w:iCs/>
                    <w:color w:val="F10D0C"/>
                    <w:sz w:val="36"/>
                    <w:szCs w:val="36"/>
                    <w:u w:val="single"/>
                </w:rPr>`
	expected := Font{
		FontSize: 36,
		//FontName: ,
		Bold:      true,
		Italic:    true,
		Color:     "F10D0C",
		Underline: "single",
		Another:   nil,
	}
	nodes, err := GetParentNodes(xml)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Println(nodes[0].Body)
	rpr, err := ParseRPR(nodes[0].Body)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, expected, rpr)
}
