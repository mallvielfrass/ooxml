package main

import (
	"fmt"

	"github.com/mallvielfrass/fmc"
	"github.com/mallvielfrass/ooxml"
)

func main() {
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

	nodes, err := ooxml.GetParentNodes(testingXML)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, item := range nodes {
		fmc.Printfln("#gbt%d) #ybt%s", i+1, item.Name)
	}
}
