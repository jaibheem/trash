people = 30
cars = 40
buses = 15

if cars > people:
	print "We should take cars"
elif cars < people:
	print "We shouldn't take cars"
else:
	print "We can't decide"

if buses > cars:
	print "That's too many buses"
elif buses < cars:
	print "May be we could take some buses"
else:
	print "we can't decide"

if people > buses:
	print "Let's just take the buses"
else:
	print "Fine, let's stay home"
