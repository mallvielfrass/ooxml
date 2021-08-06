package main

import "fmt"

type NewToken struct {
	Name           string
	TagStatus      int
	TokenSymbOpen  int
	TokenSymbClose int
	Args           string
}

const TagComplimentary = 4
const TagOpenComplimentary = 1
const TagClosingComplimentary = 2
const TagSelfClosed = 3

const Creating = true
const Empty = false

func getTokens(d string) []NewToken {
	var Tokens []NewToken
	//token flags
	name := ""
	tagStatus := TagSelfClosed //check Tag open/closed/selfclosed
	tokenStatus := Empty       // token opening( true) or not(false)
	ArgStatus := Empty
	arg := ""
	lastChar := ""
	TokenSymbOpen := 0
	TokenSymbClose := 0
	for i, l := range d {
		let := string(l)
		switch let {
		case "<":
			switch tokenStatus {
			case Creating:
				//  if lastChar!="\\"{

				//  }
			case Empty:
				tokenStatus = Creating
				tagStatus = TagOpenComplimentary
				TokenSymbOpen = i
			}
		case ">":
			switch tokenStatus {
			case Creating:
				if lastChar == "/" {
					tagStatus = TagSelfClosed
				}
				TokenSymbClose = i
				Tokens = append(Tokens, NewToken{
					Name:           name,
					TagStatus:      tagStatus,
					TokenSymbOpen:  TokenSymbOpen,
					TokenSymbClose: TokenSymbClose,
					Args:           arg,
				})
				name = ""
				tagStatus = TagSelfClosed //check Tag open/closed/selfclosed
				ArgStatus = Empty
				arg = ""
				tokenStatus = Empty
			case Empty:

			}
		case "/":
			if lastChar == "<" {
				tagStatus = TagClosingComplimentary
			}
		case "\\":
		case " ":
			if tokenStatus {
				ArgStatus = Creating
			}
		case "\t":
		case "\n":
		default:
			if tokenStatus {
				if ArgStatus {
					arg += let
				} else {
					name += let
				}
			}
		}
		lastChar = let
	}
	return Tokens
}

func getFirstNodes(tokens []NewToken) []NewToken {
	//	fmt.Printf("debug run %v\n", tokens)
	if len(tokens) == 0 {
		return []NewToken{}
	}
	token := tokens[0]
	if len(tokens) == 1 {
		return []NewToken{token}
	}
	if token.TagStatus == TagSelfClosed {
		var f []NewToken
		f = append(f, token)
		another := getFirstNodes(tokens[1:])
		f = append(f, another...)
		return f
	}
	iter := 0
	var tk []NewToken
	for i := 1; i < len(tokens); i++ {
		let := tokens[i]
		//	fmt.Printf("let : %v |token %s\n", let, token.Name)
		if let.Name == token.Name {
			if let.TagStatus == TagOpenComplimentary {
				//			fmt.Printf("iter += 1\n")
				iter += 1
			}
			if let.TagStatus == TagClosingComplimentary {
				if iter != 0 {
					//				fmt.Printf("iter -= 1\n")
					iter -= 1
				} else {
					//				fmt.Printf("tag %s closed\n", token.Name)
					tk = append(tk, token, let)
					if len(tokens) <= i {
						return tk
					}
					//				fmt.Printf("another\n")
					another := getFirstNodes(tokens[i+1:])
					tk = append(tk, another...)
					return tk
				}

			}
		}
	}
	return []NewToken{}
}

type EmbeddedToken struct {
	Name      string
	TagStatus int
	Body      string
	Args      string
}

func getParentNodes(d string) []EmbeddedToken {
	tokens := getTokens(d)
	items := getFirstNodes(tokens)
	var env NewToken
	envStatus := Empty
	var eTokens []EmbeddedToken
	for _, item := range items {
		if item.TagStatus == TagSelfClosed {
			eTokens = append(eTokens, EmbeddedToken{
				Name:      item.Name,
				TagStatus: TagSelfClosed,
				Args:      item.Args,
				Body:      "",
			})
		}
		if item.TagStatus == TagOpenComplimentary {
			if envStatus == Empty {
				envStatus = Creating
				env = item
			} else {
				fmt.Printf("error adding element %v\n", item)
				break
			}
		}
		if item.TagStatus == TagClosingComplimentary {
			if envStatus == Creating {
				eTokens = append(eTokens, EmbeddedToken{
					Name:      item.Name,
					TagStatus: TagComplimentary,
					Args:      env.Args,
					Body:      d[env.TokenSymbClose+1 : item.TokenSymbOpen],
				})

				envStatus = Empty
				env = NewToken{}
			} else {
				fmt.Printf("error adding element %v\n", item)
				break
			}
		}
	}
	return eTokens
}
