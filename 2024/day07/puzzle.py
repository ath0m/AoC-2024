import sys

from typing import NamedTuple

Equation = NamedTuple("Equation", [("result", int), ("nums", list[int])])


def parse(filename: str) -> list[Equation]:
    equations = []
    with open(filename) as f:
        for line in f.readlines():
            res, rest = line.split(":")
            nums = [int(num) for num in rest.strip().split()]
            equations.append(Equation(int(res), nums))
    return equations


def is_possible(eq: Equation, extra: bool = False) -> bool:
    if not eq.nums:
        return eq.result == 0
    if len(eq.nums) == 1:
        return eq.result == eq.nums[0]
    if max(eq.nums) > eq.result:
        return False
    a, b, *rest = eq.nums
    combs = [a + b, a * b]
    if extra:
        combs.append(int(f"{a}{b}"))
    return any(is_possible(Equation(eq.result, [comb, *rest]), extra) for comb in combs)


def p1(equations: list[Equation]) -> int:
    result = 0
    for eq in equations:
        if is_possible(eq):
            result += eq.result
    return result


def p2(equations: list[Equation]) -> int:
    result = 0
    for eq in equations:
        if is_possible(eq, extra=True):
            result += eq.result
    return result


if __name__ == "__main__":
    assert len(sys.argv) == 2, "p1, p2 argument is expected"

    equations = parse("input.txt")
    match sys.argv[1]:
        case "p1":
            print("p1:", p1(equations))
        case "p2":
            print("p2:", p2(equations))
