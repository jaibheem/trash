import sys
from sys import argv

if len(sys.argv) < 4:
	print "Total Number of arguments given are:", len(sys.argv), "please enter 4 arguments"
elif len(sys.argv) > 4:
	print "Total Number of arguments gives are:", len(sys.argv), "please enter 4 arguments"
else:
	script, first, second, third = argv
	print "your script name is:" , script
	print "First Variable is :", first
	print "Second Variable is:", second
	print "Third variable is:", third