def cheese_and_crackers(cheese_count, boxes_of_crackers):
	print "you have %d cheeses" % cheese_count
	print "you have %d boxes of crackers" % boxes_of_crackers
	print "Man that's enough for the party"
	print "get blanket"

print "we can just give the function numbers directly"
cheese_and_crackers(20, 30)

print "OR we can use variables from our script too"
amount_of_cheese = 10
amount_of_crackers = 40

cheese_and_crackers(amount_of_cheese, amount_of_crackers)

print "we can even do math inside it too"
cheese_and_crackers(10+100, 5+6)

print "Also we can combine the two, variables and math:"
cheese_and_crackers(amount_of_cheese+114, amount_of_crackers+156)

print "Enter the Number of cheeses and crackers Manually:\n"
UD_cheeses = int(raw_input("Cheeses>"))
UD_crackers = int(raw_input("Crackers>"))

cheese_and_crackers(UD_cheeses, UD_crackers)
