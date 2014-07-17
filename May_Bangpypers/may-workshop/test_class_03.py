import unittest
from class_03 import Person

class TestPerson(unittest.TestCase):
	def setUp(self):
		print "Calling setup"
		self.person = Person(first_name="Jai", last_name="D", age=24)

	def test_get_age(self):
		self.assertEqual(self.person.get_age(), 24)

	def test_is_major(self):
		self.assertTrue(self.person.is_major())

if __name__ == "__main__":
	import sys
	print sys.path
	unittest.main()