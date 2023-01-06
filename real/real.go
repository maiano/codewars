package main

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Count the divisors of a number

func Divisors(n int) int {
	count := 1
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

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

// Split Strings

func Solution(str string) []string {
	return regexp.MustCompile(`.{2}`).FindAllString(str+"_", -1)
}

// Find the missing letter

func FindMissingLetter(chars []rune) rune {
	next := chars[0] + 1
	for i := 1; i < len(chars); i++ {
		if chars[i] != next {
			return next
		}
		next = chars[i] + 1
	}
	return 'a'
}

// Highest Scoring Word

func High(s string) string {
	c := strings.Fields(s)
	max := score(c[0])
	indexMax := 0
	for i := 1; i < len(c); i++ {
		if next := score(c[i]); max < next {
			max = next
			indexMax = i
		}
	}
	return c[indexMax]
}

func score(s string) int {
	var sum int
	for _, v := range s {
		sum += int(v) - (int('a') - 1)
	}
	return sum
}

func wave(s string) []string {
	w := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			w = append(w, (s[:i])+strings.ToUpper(string(s[i]))+s[i+1:])
		}
	}
	return w
}

// Roman Numerals Decoder

func Decode(roman string) int {

	romanD := map[rune]int{'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}

	var sum, prev int

	for _, v := range roman {
		curr := romanD[v]
		if curr > prev {
			sum += curr - 2*prev
		} else {
			sum += curr
		}
		prev = curr
	}
	return sum
}

// Moving Zeros To The End

func MoveZeros(arr []int) []int {
	var j int
	temp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			temp[j] = arr[i]
			j++
		}
	}
	return temp
}

// Valid Parentheses

func ValidParentheses(parens string) bool {
	var count int
	for i := 0; i < len(parens); i++ {
		switch parens[i] {
		case '(':
			count++
		case ')':
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

// Human Readable Time

func HumanReadableTime(sec int) string {
	hh := sec / 3600
	mm := sec % 3600 / 60
	ss := sec % 60
	return fmt.Sprintf("%.2d:%.2d:%.2d", hh, mm, ss)
}

// Directions Reduction

func DirReduc(arr []string) []string {
	op := map[string]string{"NORTH": "SOUTH", "EAST": "WEST", "WEST": "EAST", "SOUTH": "NORTH"}
	path := make([]string, 0)

	for _, v := range arr {
		if len(path) > 0 && path[len(path)-1] == op[v] {
			path = path[:len(path)-1]
		} else {
			path = append(path, v)
		}
	}
	return path
}

// First non-repeating character

func FirstNonRepeating(str string) string {

	s := make(map[rune]int, 0)
	for _, v := range strings.ToLower(str) {
		s[v]++
	}
	for i, v := range strings.ToLower(str) {
		if s[v] == 1 {
			return string(str[i])
		}
	}
	return ""
}

// Best travel

func ChooseBestSum(t, k int, ls []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	if t < 0 || k > len(ls) {
		return -1
	}
	if k == 0 {
		return 0
	}
	c0 := ChooseBestSum(t, k, ls[1:])
	c1 := ChooseBestSum(t-ls[0], k-1, ls[1:])
	if c1 < 0 {
		return c0
	}
	return max(c0, ls[0]+c1)
}

// int32 to IPv4

func Int32ToIp(n uint32) string {
	ip := make([]byte, 0)
	for i := 0; i < 4; i++ {
		ip = append(ip, byte(n))
		n = n >> 8
	}
	return fmt.Sprintf("%d.%d.%d.%d", ip[3], ip[2], ip[1], ip[0])
}

// Josephus Permutation

func Josephus(items []interface{}, k int) []interface{} {
	res := []interface{}{}
	var i int
	for len(items) > 0 {
		i = (i + k - 1) % len(items)
		res = append(res, items[i])
		items = append(items[:i], items[i+1:]...)
	}
	return res
}

// Not very secure

func alphanumeric(str string) bool {
	var valid = regexp.MustCompile(`^\w+$`)
	return valid.MatchString(str)
}

// MiniStringFuck

func Interpreter(code string) string {
	var buf bytes.Buffer
	var count byte

	for _, v := range code {
		if v == '+' {
			count++
		}
		if v == '.' {
			buf.WriteByte(count)
		}
	}
	return buf.String()
}

// Beeramid

func Beeramid(bonus int, price float64) int {
	var count int
	maxCan := int(float64(bonus) / price)
	for (count+1)*(count+1) <= maxCan {
		count++
		maxCan -= count * count
	}
	return count
}

// ISBN-10 Validation

func ValidISBN10(isbn string) bool {
	isbn = strings.ToUpper(isbn)
	if ok, _ := regexp.MatchString(`[0-9]{9}[X0-9]{1}`, isbn); !ok {
		return false
	}
	var result int
	for i := 0; i < 9; i++ {
		result += (i + 1) * (int(isbn[i]) - 48)
	}
	if isbn[9] == 'X' {
		result += 100
	} else {
		result += 10 * (int(isbn[9]) - 48)
	}
	return result%11 == 0
}

// ROT13

func Rot13(msg string) string {
	var buf bytes.Buffer
	for _, v := range msg {
		b := byte(v)
		if b >= 'A' && b <= 'Z' {
			b = 'A' + (b-'A'+13)%26
		} else if b >= 'a' && b <= 'z' {
			b = 'a' + (b-'a'+13)%26
		}
		buf.WriteByte(b)
	}
	return buf.String()
}

// WeIrD StRiNg CaSe

func toWeirdCase(str string) string {
	str = strings.ToLower(str)
	var result []string
	for _, v := range strings.Fields(str) {
		s := []rune(v)
		for i := 0; i < len(s); i += 2 {
			s[i] = unicode.ToUpper(s[i])
		}
		result = append(result, string(s))
	}
	return strings.Join(result, " ")
}

// IP Validation

func Is_valid_ip(ip string) bool {
	var valid = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	return valid.MatchString(ip)
}

// The Supermarket Queue

func QueueTime(customers []int, n int) int {

	t := make([]int, n)
	for i := 0; i < n && len(customers) != 0; i++ {
		t[i] = customers[0]
		customers = customers[1:]
	}
	for len(customers) > 0 {
		i := minIndex(t)
		t[i] += customers[0]
		customers = customers[1:]
	}
	return maxOf(t)
}

func minIndex(vars []int) int {
	min := vars[0]
	index := 0
	for i, v := range vars {
		if min > v {
			min = v
			index = i
		}
	}
	return index
}

func maxOf(vars []int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func main() {
	fmt.Println("Hello")
}
