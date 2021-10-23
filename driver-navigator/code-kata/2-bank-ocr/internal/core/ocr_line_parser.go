package core

import (
	"strings"
)

func ParseOcrLine(ocrLine string) []string {
	lineParts := strings.Split(ocrLine, "\n")
	return toOcrChars(lineParts)
}

func toOcrChars(lineParts []string) []string {
	output := make([]string, 0)
	cursor := 0
	ocr := ""

	//loop for each ocr char
	for i := 0; i < 9; i++ {
		//loop for each line for each ocr char
		for _, line := range lineParts {
			//append 3 chars to the ocr
			ocr = ocr + line[cursor:cursor+3]
		}
		//append complete ocr to the output array
		output = append(output, ocr)
		//move cursor by 3
		cursor = cursor + 3
		// and reset ocr
		ocr = ""
	}

	return output
}
