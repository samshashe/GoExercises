package main

func CountCamelCaseWords(input string) int {
	if input[0] < 97 && input[0] > 122 {
		return -1
	}
	count := 1
	for _, str := range input {
		if str >= 65 && str <= 90 {
			count++
		}
	}

	return count
}
func EncryptString(input string, shift byte) string {
	if shift > 25 {
		shift = shift % 26
	}
	var result string
	var shiftedChar byte
	for i, str := range input {
		if str >= 65 && str <= 90 {
			shiftedChar = input[i] + shift
			if shiftedChar > 'Z' {
				shiftedChar = (shiftedChar - 'Z' - 1) + 'A'
			}
			result = result + string(shiftedChar)
			//input = strings.Replace(input, string(input[i]), string(input[i]+shift), 1)
		} else if str >= 97 && str <= 122 {
			shiftedChar = input[i] + shift
			if shiftedChar > 'z' {
				shiftedChar = (shiftedChar - 'z' - 1) + 'a'
			}
			result = result + string(shiftedChar)

		}
	}

	return result
}
