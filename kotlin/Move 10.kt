package codewars.kotlin

fun main() {
    val word = "helloworld"
    val str = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
    val x = mutableListOf<Char>()
    for (i in word) {
        x.add(str[(str.indexOf(i) + 10) % str.length])
    }
    println(x.joinToString("")) // rovvygybvn
}


fun move(word: String): String = word
    .map { c ->
        ('a'..'z')
            .toList()
            .run { this[(this.indexOf(c) + 10) % this.size] }
    }.joinToString("")

