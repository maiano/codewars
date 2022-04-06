package codewars.kotlin

fun solution (number: Int): Int {
    return (0 until number).filter { it % 3 == 0 || it % 5 == 0 }.sum()
}