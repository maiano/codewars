package codewars.kotlin

fun main() {
    println(getMiddle("a"))
}
fun getMiddle(word: String): String {
    return if (word.length % 2 != 0) word[word.length / 2].toString() else word.substring(word.length /2 - 1..word.length /2)
}