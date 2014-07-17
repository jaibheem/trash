class Person(object):
	def __init__(self, first_name, age, last_name=None):
		self.first_name = first_name
		self.age = age
		self.last_name = last_name
	def is_major(self):
		return self.age >= 18


person = Person(first_name="Jai", last_name="D", age=12)
print "Is the Person Major:", person.is_major(), "And his details are", person.first_name, person.age, person.last_name
