import sys
from collections import Counter


def puzzle1(l1: list[int], l2: list[int]) -> int:
    l1.sort()
    l2.sort()
    return sum(abs(v1 - v2) for v1, v2 in zip(l1, l2))


def puzzle2(l1: list[int], l2: list[int]) -> int:
    cnt = Counter(l2)
    return sum(v * cnt[v] for v in l1)


def read_input(filename: str) -> tuple[list[int], list[int]]:
    l1, l2 = [], []
    with open(filename) as f:
        for line in f.readlines():
            v1, v2 = line.split()
            l1.append(int(v1))
            l2.append(int(v2))
    return l1, l2


if __name__ == "__main__":
    assert len(sys.argv) == 2, "Needs to provide argument (puzzle1 or puzzle2)"

    puzzle = sys.argv[1]
    assert puzzle in {"puzzle1", "puzzle2"}

    l1, l2 = read_input("input.txt")

    match puzzle:
        case "puzzle1":
            print("Puzzle 1:", puzzle1(l1, l2))
        case "puzzle2":
            print("Puzzle 2:", puzzle2(l1, l2))
