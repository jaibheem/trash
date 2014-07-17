def add(a, b):
	print "Adding %d + %d" % (a, b)
	return a + b
def substarct(a, b):
	print "substarcting %d - %d" % (a, b)
	return a - b
def multiply(a, b):
	print "multiplying %d * %d" % (a, b)
	return a * b
def divide(a,b):
	print "Dividing %d / %d" % (a, b)
	return a / b

print "Let's do some math with just functions"

age = add(50, 30)
height = substarct(11, 2)
weight = multiply(15, 7)
iq = divide (100, 5)

print "Age: %d, height: %d, weight: %d, iq: %d" % (age, height, weight, iq)

#A puzzle for extra credit, type it in anyway
print "Here is the puzzle"

what = add(age, substarct(height, multiply(weight, divide(iq, 2))))
print "That becomes: ", what, "Can you do it by hand?"