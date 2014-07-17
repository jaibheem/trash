from sys import argv
script, filename = argv

txt = open(filename)
print "Here is the content of file %r:" % filename
print txt.read()

print "Would you like to read again?"
if raw_input(">") == "yes":
	print "Here you go:"
	txt = open(filename)
	print txt.read()
else:
	print "Bye Bye"