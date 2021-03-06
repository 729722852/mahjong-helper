package main

import (
	"fmt"
	"os"
	"github.com/fatih/color"
	"bufio"
)

func errorExit(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	fmt.Println("按任意键退出...")
	bufio.NewReader(os.Stdin).ReadByte()
	os.Exit(1)
}

//

// 进张数优劣
func getWaitsCountColor(shanten int, waitsCount float64) color.Attribute {
	_getWaitsCountColor := func(fixedWaitsCount float64) color.Attribute {
		switch {
		case fixedWaitsCount < 13: // 4.3*3
			return color.FgHiBlue
		case fixedWaitsCount <= 18: // 6*3
			return color.FgHiYellow
		default: // >6*3
			return color.FgHiRed
		}
	}

	if shanten == 0 {
		return _getWaitsCountColor(waitsCount * 3)
	}
	weight := 1
	for i := 1; i < shanten; i++ {
		weight *= 2
	}
	return _getWaitsCountColor(waitsCount / float64(weight))
}

// 他家中张舍牌提示
func getOtherDiscardAlertColor(index int) color.Attribute {
	if index >= 27 {
		return color.FgWhite
	}
	idx := index%9 + 1
	switch idx {
	case 1, 2, 8, 9:
		return color.FgWhite
	case 3, 7:
		return color.FgHiYellow
	case 4, 5, 6:
		return color.FgHiRed
	default:
		errorExit("[getOtherDiscardAlertColor] 代码有误: idx = ", idx)
		return -1
	}
}

// 铳率高低
func getNumRiskColor(risk float64) color.Attribute {
	switch {
	case risk < 3:
		return color.FgHiBlue
	case risk < 5:
		return color.FgHiCyan
	//case risk < 7.5:
	//	return color.FgYellow
	case risk < 10:
		return color.FgHiYellow
	case risk < 15:
		return color.FgHiRed
	default:
		return color.FgRed
	}
}

//

func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		c += 32
	}
	return c
}

func upper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		c -= 32
	}
	return c
}
