package codewars.kotlin

fun main() {
    val str = "hElLo"
    val x = if (str.count { it in 'A'..'Z' } <= str.length/2 ) str.lowercase() else str.uppercase()
    println(x)
}