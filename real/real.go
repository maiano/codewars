package main

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Two Sum
func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}

// Vowel Count
func GetCount(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return
}

// Highest and Lowest

func HighAndLow(in string) string {
	arr := strings.Split(in, " ")
	min, _ := strconv.Atoi(arr[0])
	max := min

	for _, val := range arr {
		value, _ := strconv.Atoi(val)
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}

// Get the Middle Character

func GetMiddle(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	return s[(n-1)/2 : n/2+1]
}

// Shortest Word

func FindShort(s string) int {
	min := len(s)
	for _, v := range strings.Fields(s) {
		if l := len(v); l < min {
			min = l
		}
	}
	return min
}

// Jaden Casing Strings

func ToJadenCase(str string) string {
	s := strings.Fields(str)
	for i, v := range s {
		s[i] = strings.ToUpper(v[:1]) + v[1:]
	}
	return strings.Join(s, " ")
}

func DNAStrand(dna string) string {
	acid := map[rune]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}
	var s bytes.Buffer
	for _, v := range dna {
		s.WriteByte(acid[v])
	}
	return s.String()
}

// func DNAStrand(dna string) string {
// 	return strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G").Replace(dna)
// }

// Beginner Series #3 Sum of Numbers

func GetSum(a, b int) int {
	if a > b {
		a, b = b, a
	}
	return (a + b) * (b - a + 1) / 2
}

// Two to One

func TwoToOne(s1 string, s2 string) string {
	s := strings.Split((s1 + s2), "")
	sort.Strings(s)
	// var i int
	i, prev := 0, ""
	for _, v := range s {
		if v != prev {
			s[i] = v
			prev = v
			i++
		}
	}
	return strings.Join(s[:i], "")
}

// Is this a triangle?

func IsTriangle(a, b, c int) bool {
	arr := []int{a, b, c}
	sort.Ints(arr)
	return arr[2] < arr[0]+arr[1]
}

// func IsTriangle(a, b, c int) bool {
// 	return a+b > c && b+c > a && a+c > b
// }

// Multiples of 3 or 5

func Multiple3And5(number int) int {
	var sum int
	for i := 3; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// func Multiple3And5(number int) int {
// 	n3 := (number - 1) / 3
// 	n5 := (number - 1) / 5
// 	n15 := (number - 1) / 15
// 	return (3*n3*(n3+1) + 5*n5*(n5+1) - 15*n15*(n15+1)) / 2
// }

// Array.diff

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ArrayDiff(a, b []int) (diff []int) {
	for _, v := range a {
		if !contains(b, v) {
			diff = append(diff, v)
		}
	}
	return
}

// func ArrayDiff(a, b []int) (r []int) {
// 	m := make(map[int]bool)
// 	for _, v := range b {
// 		m[v] = true
// 	}
// 	for _, v := range a {
// 		if !m[v] {
// 			r = append(r, v)
// 		}
// 	}
// 	return
// }

// Counting Duplicates

func duplicate_count(s1 string) int {
	s := make(map[rune]int, 0)
	var count int
	for _, v := range strings.ToLower(s1) {
		if s[v]++; s[v] == 2 {
			count++
		}
	}
	return count
}

// function duplicateEncode(word){
// 	return word
// 	  .toLowerCase()
// 	  .split('')
// 	  .map( function (a, i, w) {
// 		return w.indexOf(a) == w.lastIndexOf(a) ? '(' : ')'
// 	  })
// 	  .join('');
//   }

func DuplicateEncode(word string) string {
	w := []byte(strings.ToLower(word))
	var b bytes.Buffer
	for _, v := range w {
		if bytes.Count(w, []byte{v}) > 1 {
			b.WriteByte(')')
		} else {
			b.WriteByte('(')
		}
	}
	return b.String()
}

// Your order, please

func Order(s string) string {
	str := strings.Fields(s)
	re := regexp.MustCompile(`\d`)
	c := make([]string, len(str))
	for _, v := range str {
		n := re.FindString(v)
		num, _ := strconv.Atoi(n)
		c[num-1] = v
	}
	return strings.Join(c, " ")
}

// Tribonacci Sequence

// func Tribonacci(signature [3]float64, n int) []float64 {
// 	t := make([]float64, 0, n)
// 	var a, b, c float64 = signature[0], signature[1], signature[2]
// 	for i := 0; i < n; i++ {
// 		t = append(t, a)
// 		a, b, c = b, c, a+b+c
// 	}
// 	return t
// }

func Tribonacci(signature [3]float64, n int) []float64 {
	t := signature[:]
	var a, b, c float64 = signature[0], signature[1], signature[2]
	for i := 3; i < n; i++ {
		a, b, c = b, c, a+b+c
		t = append(t, c)
	}
	return t[0:n]
}

// Find the unique number

// func FindUniq(arr []float32) float32 {
// 	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
// 	if arr[0] == arr[1] {
// 		return arr[len(arr)-1]
// 	}
// 	return arr[0]
// }

func FindUniq(arr []float32) float32 {
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	for i := 1; i < len(arr); i++ {
		if arr[0] != arr[i] {
			return arr[i]
		}
	}
	return 0
}

func main() {
	fmt.Println("Hello")
}
