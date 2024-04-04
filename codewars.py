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