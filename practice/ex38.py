"""
pop will get a item from the list
"""
print __doc__
ten_things = "Apples Oranges Crows Telephone Light Sugar"
print "Let's fix this, there are not 10 items"

stuff = ten_things.split(' ')
more_stuff = ["Day", "Night", "Song", "Frisbee", "Corn", "Banana", "Girl", "Boy"]

while len(stuff) != 10:
	next_item = more_stuff.pop()
	print "Adding %s" % next_item
	stuff.append(next_item)

print stuff
print stuff[-2]