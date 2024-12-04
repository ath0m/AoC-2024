import unittest

from main import parse, p1, p2

class Test(unittest.TestCase):
    def test_parse(self):
        expected = 'xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))'
        self.assertEqual(parse('test1.txt'), expected)

    def test_p1(self):
        inp = parse('test1.txt')
        self.assertEqual(p1(inp), 161)

    def test_p2(self):
        inp = parse('test2.txt')
        self.assertEqual(p2(inp), 48)


if __name__ == '__main__':
    unittest.main()
