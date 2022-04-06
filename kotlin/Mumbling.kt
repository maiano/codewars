package codewars.kotlin

fun accum(s: String): String =
    s.mapIndexed { index, c -> "${c.uppercase()}${c.toString().lowercase().repeat(index)}" }
        .joinToString("-")

fun main() {
    println(accum("abcd"))  // A-Bb-Ccc-Dddd
    println(accum("ssup"))  // S-Ss-Uuu-Pppp
}