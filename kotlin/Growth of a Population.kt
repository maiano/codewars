package codewars.kotlin

fun main() {
    println(nb_year(1500000, 2.5, 10000,2000000))
}
fun nb_year(p0: Int, percent: Double, aug: Int, p: Int): Int {
    var year = 0
    var p00 = p0
    while (p00 < p) {
        p00 += ((p00 * percent / 100) + aug).toInt()
        year++
    }
    return year
}

