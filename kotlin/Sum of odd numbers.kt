package codewars.kotlin

fun rowSumOddNumbers(n: Int): Int {
    var temp = n * (n - 1) + 1
    var sum = temp
    repeat(n - 1) {
        temp += 2
        sum += temp
    }
    return sum
}
fun main() {
    println(rowSumOddNumbers(3))
}