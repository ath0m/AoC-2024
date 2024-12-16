import unittest

from main import puzzle1, puzzle2, parse_input

class Test(unittest.TestCase):
    def test_puzzle1(self):
        _input = parse_input('test.txt')
        self.assertEqual(puzzle1(_input), 2)

    def test_puzzle2(self):
        _input = parse_input('test.txt')
        self.assertEqual(puzzle2(_input), 4)

    def test_parse_input(self):
        expected = [[7, 6, 4, 2, 1], [1, 2, 7, 8, 9], [9, 7, 6, 2, 1], [1,3,2,4,5], [8,6,4,4,1], [1,3,6,7,9]]
        self.assertEqual(parse_input('test.txt'), expected)

if __name__ == '__main__':
    unittest.main()
