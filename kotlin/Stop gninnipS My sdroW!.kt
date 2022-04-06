package codewars.kotlin

fun spinWords(sentence: String): String {
    return sentence.split(' ').joinToString(" ") { if (it.length >= 5) it.reversed() else it }
}
fun main() {
    println(spinWords( "Hey fellow warriors" ))
}
