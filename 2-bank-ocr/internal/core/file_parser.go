package core

import (
	"strings"
)

func FileParser(fileText string) [][]string {
	output := make([][]string, 9)
	for i := 0; i < len(output); i++ {
		output[i] = make([]string, 0)
	}

	lines := strings.Split(fileText, "\n")

	for _, line := range lines {
		lineChars := strings.Split(line, "")
		output[0] = append(output[0], (lineChars[0:3])...)
		//output[1] = append(output[1], lineChars[3:6]...)
		//output[2] = append(output[2], lineChars[6:9]...)
		//output[3] = append(output[3], lineChars[9:12]...)
		//output[4] = append(output[4], lineChars[12:15]...)
		//output[5] = append(output[5], lineChars[15:18]...)
		//output[6] = append(output[6], lineChars[18:21]...)
		//output[7] = append(output[7], lineChars[21:24]...)
		//output[8] = append(output[8], lineChars[24:27]...)
	}

	return output
}
