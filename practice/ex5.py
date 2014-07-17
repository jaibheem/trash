my_name = "Jai D"
my_age = 26
my_height = 180
my_weight = 80
my_eyes = 'wheatish'
my_teeth = 'White'
my_hair = 'black'
my_sex = 'male'

if my_sex == 'male':
	print "Let's talk about Mr.%s" %my_name
else:
	print "Let's talk about Ms.%s" %my_name

print "He is %d inches tall" %my_height
print "He is %d kgs heavy" %my_weight
print "He is fat"
print "He has got %s eyes and %s hair" %(my_eyes, my_hair)
print "His teeth are usually %s depending on the coffee" %my_teeth

# this line is tricky, try to get it exactly right
print "If I add %d, %d and %d I get %d" %(my_age, my_height, my_weight, my_age+ my_height+ my_weight)