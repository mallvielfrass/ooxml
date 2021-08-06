package main

// func TestTokenize(t *testing.T) {
// 	d := `<w:p>
// 	        <w:pPr>
// 	           <w:pStyle w:val="Normal"/>
// 	            <w:rPr>
// 	                <w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/>
// 	                <w:color w:val="00000A"/>
// 	                <w:sz w:val="24"/>
// 	                <w:szCs w:val="24"/>
// 	                <w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/>
// 	            </w:rPr>
// 	        </w:pPr>
// 	        <w:r>
// 	            <w:rPr></w:rPr>
// 	            <w:t>This is a</w:t>
// 	        </w:r>
// 	        <w:bookmarkStart w:id="0" w:name="_GoBack"/>
// 	        <w:bookmarkEnd w:id="0"/>
// 	        <w:r>
// 	            <w:rPr></w:rPr>
// 	            <w:t xml:space="preserve"> word document.</w:t>
// 	        </w:r>
// 	    </w:p>`
// 	parsed := getFirstNodes(d)
// 	//fmt.Println(len(parsed))
// 	if len(parsed) <= 0 {
// 		t.Error("Expected nodes")
// 		return
// 	}
// 	wP := parsed[0]
// 	if wP.Name != "w:p" {
// 		t.Error("Expected 'w:p', got ", wP.Name)
// 	}

// }
