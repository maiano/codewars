package codewars.kotlin

fun newAvg(a: DoubleArray, navg: Double): Long {
    val avg = navg * (a.size + 1) - a.sum()
    val result = (kotlin.math.ceil(avg)).toLong()
    return if (result > 0) result else throw IllegalArgumentException()
}
fun main() {
    var a = doubleArrayOf(14.0, 30.0, 5.0, 7.0, 9.0, 11.0, 15.0)
    println(newAvg(a, 92.0))
}