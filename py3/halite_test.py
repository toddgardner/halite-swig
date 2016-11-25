import unittest
import halite

class TestHaliteBindings(unittest.TestCase):

  def test_random_map(self):
    expected = '1 0 1 2 1 1 1 0 205 194 194 205 '
    actual = halite.randomMap(2, 2, 2, 3704032075)
    self.assertEqual(actual, expected)

if __name__ == '__main__':
  unittest.main()
