"""Class is something that just contains the structure but Instance is the blueprint of the class 
which defines how the class looks like.. more over like template and manifest"""
#Example 1:
def greet(name):
	return name
print "Hello", greet("Jai")
print "Hello", greet("Anil")

#Example 2:
prices = {"Apple" : 40, "Banana" : 4}
my_purchases = {"Apple" : 1, "Banana" : 10}

total_price = sum(prices[fruit] * my_purchases[fruit]
				  for fruit in my_purchases)
print total_price

#Example 3:
class BankAccount(object):
	def __init__(self, initial_balance=0):
		self.balance = initial_balance
	def deposit(self, amount):
		self.balance += amount
	def withdrawl(self, amount):
		self.balance -= amount
	def overdraw(self):
		return self.balance < 0
my_account = BankAccount(15)
my_account.withdrawl(5)
print my_account.balance
my_account.deposit(20)
print my_account.balance

#Example 4:
class ClassName:
	'This class is used for documeting purpose'

print ClassName.__doc__

#Example 5:
class Employee:
	'This class is used for documeting purpose'
	emp_count = 0
	def __init__(self, name, salary):
		self.name = name
		self.salary = salary
		Employee.emp_count += 1

	def displyCount(self):
		print "Total Number of employees : %d" % Employee.emp_count
	def displayEmployee(self):
		print "Employee Name :", self.name, ", Salary:", self.salary

emp1 = Employee("Jai", 50)
emp2 = Employee("Anil", 70)

emp1.displayEmployee()
emp2.displayEmployee()

print "Total Number of Employees are:", Employee.emp_count

print "Employee.__doc__:", Employee.__doc__
print "Employee.__dict__:", Employee.__dict__
print "Employee.__name__:", Employee.__name__
print "Employee.__module__:", Employee.__module__
print "Employee.__bases__:", Employee.__bases__

#Example 6: Inhertience and Method Overriding
class Parent:
	def myMethod(self):
		print "I am a parent class"
class Child(Parent):
	def childMethod(self):
		print "I am child Class"
c = Child()
c.myMethod()

p = Parent()
p.myMethod()

#Example 7: Hiding
class JustCounter:
	__secretCount = 0
	def count(self):
		self.__secretCount += 1
		print self.__secretCount

counter = JustCounter()
counter.count()
counter.count()
print counter._JustCounter__secretCount

#Example 7 :Pets example for Class
class Pet(object):
	def __init__(self, name, species):
		self.name = name
		self.species = species
	def getName(self):
		return self.name
	def getSpecies(self):
		return self.species
	def __str__(self):
		return "%s is a %s" % (self.name, self.species)
class Dog(Pet):
	def __init__(self, name, chases_cat):
		Pet.__init__(self, name, "Dog")
		self.chases_cat = chases_cat
	def chasesCat(self):
		return self.chases_cat
class Cat(Pet):
	def __init__(self, name, hate_dogs):
		Pet.__init__(self, name, "Cat")
		self.hate_dogs = hate_dogs
	def hateDogs(self):
		return self.hate_dogs