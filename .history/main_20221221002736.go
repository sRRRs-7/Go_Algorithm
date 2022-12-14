package main

import (
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
	utils.Contains()
	utils.Cipher()
	utils.ConvertBase64()
	if is_io {
		utils.Io() // if you execute this function, is_io change true
	}
	// staticAnalyzer
	analyzer.StaticAnalyzer()
	// regex
	regex.RegexFunc()
	// language processing
	hebon

}
