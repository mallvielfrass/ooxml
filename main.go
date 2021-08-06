package main

import (
	"github.com/mallvielfrass/fmc"
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
	nodes := getParentNodes(testingXML)
	for i, item := range nodes {
		fmc.Printfln("#gbt%d) #ybt%s", i+1, item.Name)
	}
}
