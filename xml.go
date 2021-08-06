package main

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/mallvielfrass/fmc"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"p"`
	WpPr    string   `xml:"pPr"`
	Wr      []string `xml:"rPr"`
	//Description string   `xml:",innerxml"`
}

func concate(s []string) string {
	var body string
	for i := 0; i < len(s); i++ {
		body += s[i]
	}
	return body
}
func getNode(s string) (string, []string) {
	f := strings.Split(s, ">")
	firstTagSplit := f[0]
	firstTag := firstTagSplit //+ ">"
	prbl := strings.Split(strings.Split(firstTag, "<")[1], " ")[0]
	tagWithParams := firstTag + ">"
	closeFirst := "</" + prbl + ">"

	arr := strings.Split(s, tagWithParams)
	bodyPrepare := concate(arr[1:])
	bodyAP := strings.Split(bodyPrepare, closeFirst)

	body := bodyAP[:len(bodyAP)-1]
	fmt.Printf("body: %s\n", body)
	fmt.Println(closeFirst)
	return prbl, []string{}
}
func tokenize(d string) []Token {

	token := ""
	tokenT := false
	argT := false
	lastSymbol := ""
	arg := " "
	//mainToken := ""
	tokens := []Token{}
	typeT := ""
	fS := 0
	LS := 0
	tokN := 0
	for i, l := range d {
		let := string(l)
		//fmt.Printf("%s", )
		switch let {
		case "<":
			if !checkScreen(lastSymbol) {
				if !tokenT {
					tokenT = true
					fS = i
					typeT = "opened"
				} else {
					fmt.Printf("error elem [%s]\n", let)
				}
			} else {
				arg += let
			}

		case ">":
			if !checkScreen(lastSymbol) {
				if tokenT {
					LS = i
					//				fmt.Printf("token: %s, params: %s type: %s\n", token, arg, typeT)
					tokens = append(tokens, Token{
						Name:        token,
						Arg:         arg,
						TypeT:       typeT,
						OpenSymbol:  fS,
						CloseSymbol: LS,
						TokenNumber: tokN,
					})
					tokN += 1
					tokenT = false
					token = ""
					arg = " "
					argT = false
				} else {
					fmt.Printf("error elem [%s]\n", let)
				}
			} else {
				arg += let
			}
		case " ":
			if tokenT {
				if !argT {
					argT = true
				}
				arg += let
			}
		case "/":
			if !checkScreen(lastSymbol) {
				if lastSymbol != "<" {
					typeT = "selfclosed"
				} else {
					typeT = "closed"
				}

			} else {
				if argT {
					arg += let
				} else {
					token += let
				}
			}
		default:
			if tokenT {
				if argT {
					arg += let
				} else {
					token += let
				}

				//fmt.Printf("tokenDef: %s\n", token)
			}
		}
		lastSymbol = let
	}
	return tokens
}
func getParent(d []Token) (Token, Token) {
	//var nodes []Token
	searchNode := d[0]
	searchNodeName := searchNode.Name
	fmt.Printf("searched node: %s\n", searchNodeName)
	iter := 0
	for _, token := range d {
		fmc.Printfln("ranged node: #gbt%s", token.Name)
		if token.Name == searchNodeName {

			if token.TypeT == "opened" {
				iter += 1
				fmc.Printfln("iter node: #ybt%s", token.Name)
			}
			if token.TypeT == "closed" {
				//
				if iter < 2 {
					fmc.Printfln("node finded: #ybt%s", token.Name)
					return searchNode, token
				} else {
					iter -= 1
				}

			}
		}
		//	if
	}
	return Token{}, Token{}
}

type TokenNode struct {
	Name string
	Body string

	Arg   string
	TypeT string
}

// func getParentNodes(tokens []Token, d string) []Token {

// 	lenTok := len(tokens)
// 	var toks []Token
// 	// toks = append(toks, Token{
// 	// 	Name: s1.Name,
// 	// 	Arg:  s1.Arg,
// 	// 	Body: d[s1.CloseSymbol:s2.OpenSymbol],
// 	// })
// 	//i := 1
// 	for {
// 		s1, s2 := getParent(tokens)
// 		fmc.Printfln("s1: %v | s2: %v ", s1, s2)
// 		fmt.Printf("[s1.CloseSymbol+1 %d: s2.OpenSymbol %d\n", s1.CloseSymbol+1, s2.OpenSymbol)
// 		toks = append(toks, Token{
// 			Name: s1.Name,
// 			Arg:  s1.Arg,
// 			Body: d[s1.CloseSymbol+1 : s2.OpenSymbol],
// 		})
// 		fmt.Printf("s1 pos: %d s2 pos: %d lenTok: %d\n", s1.TokenNumber, s2.TokenNumber, lenTok)
// 		if s2.TokenNumber == lenTok-1 {
// 			return toks
// 		}
// 		tokens = tokens[s2.TokenNumber:]
// 	}

// 	//	return toks
// }

func main() {
	// d := `<w:p>
	//         <w:pPr>
	//            <w:pStyle w:val="Normal"/>
	//             <w:rPr>
	//                 <w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/>
	//                 <w:color w:val="00000A"/>
	//                 <w:sz w:val="24"/>
	//                 <w:szCs w:val="24"/>
	//                 <w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/>
	//             </w:rPr>
	//         </w:pPr>
	//         <w:r>
	//             <w:rPr></w:rPr>
	//             <w:t>This is a</w:t>
	//         </w:r>
	//         <w:bookmarkStart w:id="0" w:name="_GoBack"/>
	//         <w:bookmarkEnd w:id="0"/>
	//         <w:r>
	//             <w:rPr></w:rPr>
	//             <w:t xml:space="preserve"> word document.</w:t>
	//         </w:r>
	//     </w:p>`

	// //f := []rune(d)
	// tokens := tokenize(d)
	// t := getParentNodes(tokens, d)
	// fmt.Println(t[0].Body)
	// fmt.Println(len(t[0].Body))
	// r := `<w:pPr>
	//             <w:pStyle w:val="Normal"/>
	//             <w:rPr>
	//                 <w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/>
	//                 <w:color w:val="00000A"/>
	//                 <w:sz w:val="24"/>
	//                 <w:szCs w:val="24"/>
	//                 <w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/>
	//             </w:rPr>
	//         </w:pPr>`
	// tR := tokenize(t[0].Body)
	// fmt.Println(tR)
	// for _, item := range tR {
	// 	fmt.Printf("item: %+v \n", item)
	// }
	// fmc.Println("#gbtLine")
	// getParentNodes(tR, t[0].Body)
	// dRune := []rune(d)
	// tokenALen := len(tokens)
	// for v := 0; v < tokenALen; v++ {
	// 	tk := tokens[v]
	// 	fmt.Println(string(dRune[tk.OpenSymbol : tk.CloseSymbol+1]))
	// }

}
