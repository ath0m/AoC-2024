import re
import sys

def parse(filename: str) -> str:
    with open(filename) as f:
        return f.read().strip()

def p1(inp: str) -> int:
    pattern = r"mul\((\d{1,3},\d{1,3})\)"
    result = 0
    for nums in re.findall(pattern, inp):
        a, b = nums.split(',')
        result += int(a) * int(b)
    return result

def p2(inp: str) -> int:
    pattern = "|".join((
        r"mul\(\d{1,3},\d{1,3}\)",
        r"do\(\)",
        r"don\'t\(\)",
    ))
    result = 0
    enable = True
    for op in re.findall(f"({pattern})", inp):
        if op.startswith("mul"):
            if enable:
                a, b = op[4:-1].split(",")
                result += int(a) * int(b)
        elif op == "do()":
            enable = True
        else:
            enable = False

    return result

if __name__ == '__main__':
    assert len(sys.argv) == 2
    inp = parse('input.txt')
    match sys.argv[1]:
        case 'p1':
            print('p1:', p1(inp))
        case 'p2':
            print('p2:', p2(inp))
        case _:
            print('parameter needs to be p1 or p2')
