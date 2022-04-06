package codewars.kotlin

fun find(integers: Array<Int>): Int {
    return if (integers.take(3).count { it and 1 == 0 } > 1) integers.first { it and 1 != 0 }
    else integers.first { it and 1 == 0 }
}

fun main() {
    println(find(arrayOf(2, 4, 0, 100, 4, 11, 2602, 36))) // 11
    println(find(arrayOf(160, 3, 1719, 19, 11, 13, -21))) // 160
}