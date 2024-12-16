import unittest

from main import parse, p1, p2


class Test(unittest.TestCase):
    def test_parse(self):
        expected = [
            "MMMSXXMASM",
            "MSAMXMSMSA",
            "AMXSXMAAMM",
            "MSAMASMSMX",
            "XMASAMXAMM",
            "XXAMMXXAMA",
            "SMSMSASXSS",
            "SAXAMASAAA",
            "MAMMMXMMMM",
            "MXMXAXMASX",
        ]
        self.assertEqual(parse("test1.txt"), expected)

    def test_p1(self):
        inp = parse("test1.txt")
        self.assertEqual(p1(inp), 18)

    def test_p2(self):
        inp = parse("test1.txt")
        self.assertEqual(p2(inp), 9)


if __name__ == "__main__":
    unittest.main()
