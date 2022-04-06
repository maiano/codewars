package codewars.kotlin

fun main() {
    println(highAndLow("1 2 -3 4 5"))
}

fun highAndLow(numbers: String): String{
    val list = numbers.split(" ").map { it.toInt() }.sorted()
    return "${list.last()} ${list.first()}"
}