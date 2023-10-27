package main

import (
	"fmt"
	"strconv"
	"strings"
)

var indexTable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxys0123456789+/"
var paddingMap = map[int]string {
	1: "=",
	2: "==",
}
const bitChunk = 6

func main() {
	var toEncode = "alphabetagamma"
	encodedString := encodeString(toEncode)
	fmt.Println(encodedString)
}

func encodeString(toEncode string) string {
	// fetch chunk of 3 characters and encode
	var outputString []string = make([]string, 1)
	
	var width, i = 3, 0
	for ; i<len(toEncode); i+=width {
		if i + width <= len(toEncode) {
			var encoded = encodeStringChunk(toEncode[i:i+width])
			outputString = append(outputString, encoded...)
		} else { // check for padding
				lastChunk := paddingChunk(i, width, toEncode)
				outputString = append(outputString, lastChunk...)
		}
	}
	return string(strings.Join(outputString, ""))
}
func paddingChunk(i int, width int ,toEncode string) []string {
	if (i + width) - len(toEncode) == 1 {
		return processResidueChunk(1, "%08b", toEncode[i:i+width-1])
	} else {
		return processResidueChunk(2, "%016b", toEncode[i:i+width-2])
	}
}
func processResidueChunk(padding int, formatStr string, chunk string) []string {
	var chunkBytes []string = make([]string, 0)
	for i :=0; i< len(chunk); i++ {
		bitString := fmt.Sprintf("%08b",chunk[i])
		for k:=0; k < len(bitString); k++ {
			chunkBytes = append(chunkBytes, string(bitString[k]))
		} 
	}
	zeroBitString := fmt.Sprintf(formatStr, 0)
	for k:=0; k < len(zeroBitString); k++ {
		chunkBytes = append(chunkBytes, string(zeroBitString[k]))
	}

	encodeString := processChunk(chunkBytes)
	encodeString = append(encodeString[:len(encodeString) - padding], paddingMap[padding] )
	return encodeString
}
func encodeStringChunk(chunk string) []string {
	var chunkBytes []string = make([]string, 0)

	for i :=0; i< len(chunk); i++ {
		bitString := fmt.Sprintf("%08b",chunk[i])
		for k:=0; k < len(bitString); k++ {
			chunkBytes = append(chunkBytes, string(bitString[k]))
		} 
	}
	return processChunk(chunkBytes)
}

func processChunk(chunkBytes []string) []string {
	encodeString := make([]string,0)
	for i:=0; i<len(chunkBytes); i+=bitChunk {
		chunkBits := string(strings.Join(chunkBytes[i:i+bitChunk], ""))
		if d, err := strconv.ParseInt(chunkBits, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(d)
			if d>=0 && int(d)< len(indexTable) {
				encodeString = append(encodeString,string(indexTable[d]))
			}
		}
	}
	return encodeString
}