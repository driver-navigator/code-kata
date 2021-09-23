package core

var (
	bits0 = [9]string{"", "_", "", "|", "", "|", "|", "_", "|"}
	bits1 = [9]string{"", "", "", "", "", "|", "", "", "|"}
	bits2 = [9]string{"", "_", "", "", "_", "|", "|", "_", ""}
	bits3 = [9]string{"", "_", "", "", "_", "|", "", "_", "|"}
	bits4 = [9]string{"", "", "", "|", "_", "|", "", "", "|"}
	bits5 = [9]string{"", "_", "", "|", "_", "", "", "_", "|"}
	bits6 = [9]string{"", "_", "", "|", "_", "", "|", "_", "|"}
	bits7 = [9]string{"", "_", "", "", "|", "", "", "", "|"}
	bits8 = [9]string{"", "_", "", "|", "_", "|", "|", "_", "|"}
	bits9 = [9]string{"_", "", "", "|", "_", "|", "", "_", "|"}
)

func ParseDigit(input [9]string) int8 {
	switch input {
	case bits1:
		return 1
	case bits2:
		return 2
	case bits3:
		return 3
	case bits4:
		return 4
	case bits5:
		return 5
	case bits6:
		return 6
	case bits7:
		return 7
	case bits8:
		return 8
	case bits9:
		return 9
	default:
		return 0
	}
}
