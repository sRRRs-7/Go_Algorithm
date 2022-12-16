package main

import (
	"fmt"
	"time"

	"github.com/sRRRs-7/Go_Algorithm.git/algorithm"
	"github.com/sRRRs-7/Go_Algorithm.git/atCorder"
	"github.com/sRRRs-7/Go_Algorithm.git/leetcode"
	"github.com/sRRRs-7/Go_Algorithm.git/regex"
	analyzer "github.com/sRRRs-7/Go_Algorithm.git/staticAnalizer"
	"github.com/sRRRs-7/Go_Algorithm.git/utils"
)

const is_io = false
const is_leet = false

func main() {
	// algorithm
	algorithm.BinarySearch()
	algorithm.EuclidMethod()
	algorithm.MergeSort()
	algorithm.Distinct()
	// leetcode
	leetcode.ReverseWords()
	leetcode.SelfCrossing()
	leetcode.AggregateComponent()
	leetcode.TotalAppealString()
	leetcode.GroupOfStrings()
	leetcode.CountIdealArray()
	leetcode.MinimumCostAve()
	leetcode.SubSequence()
	leetcode.MatchStr()
	leetcode.MaxDot()
	leetcode.LongestPrefix()
	//atCorder
	if is_leet {
		atCorder.CutCake()
	}
	atCorder.Parenthesis()
	atCorder.GridMath()
	atCorder.Dp()
	// utils
	utils.BaseNumberConvert()
	utils.ArrayToInteger()
	utils.StringContains()
	utils.SequenceArray()
	utils.MultiThread()
	if is_io {
		utils.Io() // if you execute this function, is_io change true
	}
	utils.Cipher()
	// staticAnalyzer
	analyzer.StaticAnalyzer()
	// regex
	regex.RegexFunc()

	fmt.Println(time.Now)
}
