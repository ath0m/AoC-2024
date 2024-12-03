import sys

def parse_input(filename: str) -> list[list[int]]:
    grid = []
    with open(filename) as f:
        for line in f.readlines():
            grid.append([int(v) for v in line.split()])
    return grid

def _check_line(line: list[int]) -> bool:
    for i in range(len(line) - 1):
        if not (1 <= line[i+1] - line[i] <= 3):
            return False
    return True


def puzzle1(grid: list[list[int]]) -> int:
    result = 0
    for line in grid:
        if _check_line(line) or _check_line(line[::-1]):
            result += 1
    return result

def puzzle2(grid: list[list[int]]) -> int:
    result = 0
    for line in grid:
        for i in range(len(line)):
            _line = line[:i] + line[i+1:]
            if _check_line(_line) or _check_line(_line[::-1]):
                result += 1
                break
    return result

if __name__ == '__main__':
    assert len(sys.argv) == 2

    _input = parse_input('input.txt')
    match sys.argv[1]:
        case '1':
            print('puzzle 1:', puzzle1(_input))
        case '2':
            print('puzzle 2:', puzzle2(_input))
