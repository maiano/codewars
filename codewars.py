import functools
from collections import Counter


# Replace With Alphabet Position
def alphabet_position(text: str):
    return " ".join(str(ord(c) - 96) for c in text.lower() if c.isalpha())


# Duplicate Encoder
def duplicate_encode(word):
    word = word.lower()
    counter = Counter(word)
    return "".join(")" if counter[c] > 1 else "(" for c in word)


# Counting Duplicates
def duplicate_count(text: str):
    hashmap = {}
    count = 0
    for c in text.lower():
        hashmap.setdefault(c, 0)
        hashmap[c] += 1
        if hashmap[c] == 2:
            count += 1
    return count


# Array.diff
def array_diff(a, b):
    return [x for x in a if x not in b]


# Who likes it?
def likes(names):
    match names:
        case []:
            return "no one likes this"
        case [a]:
            return f"{a} likes this"
        case [a, b]:
            return f"{a} and {b} like this"
        case [a, b, c]:
            return f"{a}, {b} and {c} like this"
        case [a, b, _, *rest]:
            return f"{a}, {b} and {len(rest)+1} others like this"


# Find the odd int
def find_it(seq):
    return functools.reduce(lambda x, y: x ^ y, seq)


# Stop gninnipS My sdroW!
def spin_words(sentence: str):
    return " ".join(s if len(s) < 5 else s[::-1] for s in sentence.split())


def solution(number):
    return sum(i for i in range(3, number) if i % 3 == 0 or i % 5 == 0)


def find_short(s):
    return min(map(len, s.split()))


# def find_short(s):
#     return min(len(x) for x in s.split())


def high_and_low(numbers):
    return " ".join(i(numbers.split(), key=int) for i in (max, min))


def get_count(sentence):
    return sum(c in "aeiou" for c in sentence)


# return masked string
def maskify(cc):
    return "#" * (len(cc) - 4) + cc[-4:]


# def maskify(cc):
#     return cc[-4:].rjust(len(cc), "#")


def filter_list(l):
    return [i for i in l if isinstance(i, int)]


def accum(st):
    return "-".join((s * i).title() for i, s in enumerate(st, 1))
