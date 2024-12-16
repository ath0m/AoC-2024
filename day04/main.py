import sys


def parse(filename: str) -> list[str]:
    with open(filename) as f:
        return [line.strip() for line in f.readlines()]


def _check_text(lines: list[str], i: int, j: int, di: int, dj: int, text: str) -> bool:
    if not text:
        return True
    h, w = len(lines), len(lines[0])
    if not (0 <= i < h and 0 <= j < w):
        return False
    if lines[i][j] != text[0]:
        return False
    return _check_text(lines, i + di, j + dj, di, dj, text[1:])


def p1(lines: list[str]) -> int:
    h, w = len(lines), len(lines[0])
    result = 0
    for i in range(h):
        for j in range(w):
            for di, dj in (
                (0, 1),
                (0, -1),
                (1, 0),
                (-1, 0),
                (-1, -1),
                (-1, 1),
                (1, 1),
                (1, -1),
            ):
                if _check_text(lines, i, j, di, dj, "XMAS"):
                    result += 1
    return result


# M.S M.M S.S S.M
# .A. .A. .A. .A.
# M.S S.S M.M S.M
PATTERNS = [
    ["M.S", ".A.", "M.S"],
    ["M.M", ".A.", "S.S"],
    ["S.S", ".A.", "M.M"],
    ["S.M", ".A.", "S.M"],
]


def _check_pattern(lines: list[str], i: int, j: int, pattern: list[str]) -> bool:
    for di in range(3):
        for dj in range(3):
            if pattern[di][dj] != "." and lines[i + di][j + dj] != pattern[di][dj]:
                return False
    return True


def p2(lines: list[str]) -> int:
    h, w = len(lines), len(lines[0])
    result = 0
    for i in range(h - 2):
        for j in range(w - 2):
            for p in PATTERNS:
                if _check_pattern(lines, i, j, p):
                    result += 1
                    break
    return result


if __name__ == "__main__":
    assert len(sys.argv) == 2
    inp = parse("input.txt")
    match sys.argv[1]:
        case "p1":
            print("p1:", p1(inp))
        case "p2":
            print("p2:", p2(inp))
        case _:
            print("parameter needs to be p1 or p2")
