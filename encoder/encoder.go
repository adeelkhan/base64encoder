package encoder

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var indexTable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var paddingMap = map[int]string{
	1: "=",
	2: "==",
}

const bitChunk = 6
const decodeBitChunk = 8

// func main() {
// 	var toEncode = "should is should is should is"
// 	encodedString := EncodeString(toEncode)
// 	fmt.Println(toEncode, encodedString)

//		var toDecode, _ = DecodeString("c2hvdWxkIGlzIHNob3VsZCBpcyBzaG91bGQgaXM=")
//		fmt.Println(toDecode)
//	}
func DecodeString(toDecode string) (string, error) {
	var decodeString = make([]string, 0)
	var width = 4
	var val []string
	var err error
	var end int

	for i := 0; i < len(toDecode); i += width {
		if string(toDecode[i+width-1]) != "=" {
			end = i + width
			val, err = decodeStringChunk(toDecode[i:end])
			if err != nil {
				return "", err
			}
			decodeString = append(decodeString, val...)
		} else {
			if string(toDecode[i+width-1]) == "=" &&
				string(toDecode[i+width-2]) == "=" {
				end = i + 2
				val, err = decodeStringChunk(toDecode[i:end])
				if err != nil {
					return "", err
				}
				// append first one bytes
				decodeString = append(decodeString, val[:1]...)
			} else {
				end = i + width - 1
				val, err = decodeStringChunk(toDecode[i:end])
				if err != nil {
					return "", err
				}
				// append first two bytes
				decodeString = append(decodeString, val[:2]...)
			}
		}
	}
	return string(strings.Join(decodeString, "")), nil
}
func decodeStringChunk(chunk string) ([]string, error) {
	chunkBytes := make([]string, 0)
	for i := 0; i < len(chunk); i++ {
		var m int
		for m = 0; m < len(indexTable); m++ {
			if string(chunk[i]) == string(indexTable[m]) {
				break
			}
		}
		if m == len(indexTable) {
			return []string{""}, errors.New("incorrect string to decode")
		}
		bitString := fmt.Sprintf("%06b", m)
		for k := 0; k < len(bitString); k++ {
			chunkBytes = append(chunkBytes, string(bitString[k]))
		}
	}
	decodeString := processDecode(chunkBytes)
	return decodeString, nil
}

func processDecode(chunkBytes []string) []string {
	decodeString := make([]string, 0)
	for i := 0; i < len(chunkBytes); i += decodeBitChunk {
		chunkBits := string(strings.Join(chunkBytes[i:i+decodeBitChunk], ""))
		if d, err := strconv.ParseInt(chunkBits, 2, 8); err != nil {
			fmt.Println(err)
		} else {
			decodeString = append(decodeString, string(byte(d)))
		}
	}
	return decodeString
}

func EncodeString(toEncode string) string {
	// fetch chunk of 3 characters and encode
	var outputString []string = make([]string, 1)

	var width, i = 3, 0
	for ; i < len(toEncode); i += width {
		if i+width <= len(toEncode) {
			var encoded = encodeStringChunk(toEncode[i : i+width])
			outputString = append(outputString, encoded...)
		} else { // check for padding
			lastChunk := paddingChunk(i, width, toEncode)
			outputString = append(outputString, lastChunk...)
		}
	}
	return string(strings.Join(outputString, ""))
}

func paddingChunk(i int, width int, toEncode string) []string {
	if (i+width)-len(toEncode) == 1 {
		return processResidueChunk(1, "%08b", toEncode[i:i+width-1])
	} else {
		return processResidueChunk(2, "%016b", toEncode[i:i+width-2])
	}
}

func processResidueChunk(padding int, formatStr string, chunk string) []string {
	var chunkBytes []string = make([]string, 0)
	for i := 0; i < len(chunk); i++ {
		bitString := fmt.Sprintf("%08b", chunk[i])
		for k := 0; k < len(bitString); k++ {
			chunkBytes = append(chunkBytes, string(bitString[k]))
		}
	}
	zeroBitString := fmt.Sprintf(formatStr, 0)
	for k := 0; k < len(zeroBitString); k++ {
		chunkBytes = append(chunkBytes, string(zeroBitString[k]))
	}

	encodeString := processChunk(chunkBytes)
	encodeString = append(encodeString[:len(encodeString)-padding], paddingMap[padding])
	return encodeString
}

func encodeStringChunk(chunk string) []string {
	var chunkBytes []string = make([]string, 0)

	for i := 0; i < len(chunk); i++ {
		bitString := fmt.Sprintf("%08b", chunk[i])
		for k := 0; k < len(bitString); k++ {
			chunkBytes = append(chunkBytes, string(bitString[k]))
		}
	}
	return processChunk(chunkBytes)
}

func processChunk(chunkBytes []string) []string {
	encodeString := make([]string, 0)
	for i := 0; i < len(chunkBytes); i += bitChunk {
		chunkBits := string(strings.Join(chunkBytes[i:i+bitChunk], ""))
		if d, err := strconv.ParseInt(chunkBits, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			if d >= 0 && int(d) < len(indexTable) {
				encodeString = append(encodeString, string(indexTable[d]))
			}
		}
	}
	return encodeString
}
