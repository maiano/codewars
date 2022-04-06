package codewars.kotlin

fun main() {
    val n = 25
    val d = 1
    var count = 0
    val x = d.toChar() - 0

    for (el in 1..n) {
        count += (el*el).toString().count { it == d.toString()[0] }
    }
    println(count)
}

