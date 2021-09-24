package core

var ocrCharMap = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

func ParseOcrChar(key string) string {
	return ocrCharMap[key]
}
