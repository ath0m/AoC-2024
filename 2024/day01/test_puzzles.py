import unittest

from puzzles import puzzle1, puzzle2, read_input


class TestPuzzles(unittest.TestCase):
    def test_puzzle1(self):
        l1, l2 = read_input("test.txt")
        self.assertEqual(puzzle1(l1, l2), 11)

    def test_read_input(self):
        self.assertEqual(
            read_input("test.txt"), ([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3])
        )

    def test_puzzle2(self):
        l1, l2 = read_input("test.txt")
        self.assertEqual(puzzle2(l1, l2), 31)


if __name__ == "__main__":
    unittest.main()
