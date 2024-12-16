import unittest

from .puzzle import is_possible, p1, p2, parse, Equation


class TestMain(unittest.TestCase):
    def test_parse(self):
        expected = [
            (190, [10, 19]),
            (3267, [81, 40, 27]),
            (83, [17, 5]),
            (156, [15, 6]),
            (7290, [6, 8, 6, 15]),
            (161011, [16, 10, 13]),
            (192, [17, 8, 14]),
            (21037, [9, 7, 18, 13]),
            (292, [11, 6, 16, 20]),
        ]
        self.assertEqual(parse("test.txt"), expected)

    def test_p1(self):
        equations = parse("test.txt")
        self.assertEqual(p1(equations), 3749)

    def test_p2(self):
        equations = parse("test.txt")
        self.assertEqual(p2(equations), 11387)

    def test_is_possible_extra(self):
        eq = Equation(7290, [6, 8, 6, 15])
        self.assertTrue(is_possible(eq, extra=True))


if __name__ == "__main__":
    unittest.main(verbosity=2)
