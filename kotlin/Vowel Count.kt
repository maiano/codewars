package codewars.kotlin

fun main() {
    val str = "abracadabra"
    val range = arrayOf('a', 'e', 'i', 'o', 'u')
    var count = 0
    for (el in str) {
        if (range.contains(el)) count++
    }
    println(count)
}

