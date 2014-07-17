the_count = [1, 2, 3, 4, 5]
fruits = ['apples', 'oranges', 'pears', 'apricots']
changes = [1, 'pennies', 2, 'dimes', 3, 'quarters']

# this first kind of for-loop goes through a list

for number in the_count:
	print "This is count %d" % number

# same as above
for fruit in fruits:
	print "A fruit of Type : %s" % fruit

# also we can go through mixed lists too
# notice we have to use %r since we don't know what's in it
for i in changes:
	print "I got %r" % i

# we can also build lists, first start with an empty one
elements = []

# then use the range function to do 0 to 5 counts
for i in range(0, 6):
	print "Adding %i to the list" % i
	# append is a function that lists understand
	elements.append(i)
	print elements

print elements
print elements[0:3]
print elements * 2
print elements[1]
# now we can print them out too
for i in elements:
	print "Element was: %d" % i