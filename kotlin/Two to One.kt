package codewars.kotlin

fun longest(s1:String, s2:String):String {
    return (s1 + s2).toHashSet().joinToString("")
//    return (s1 + s2).toSortedSet().joinToString("")
}