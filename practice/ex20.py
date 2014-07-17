from sys import argv

script, input_file = argv

def print_all(p):
	print p.read()

#A file in Python is kind of like an old tape drive on a mainframe, or maybe a DVD player. 
#It has a "read head," and you can "seek" this read head around the file to positions, then work with it there. 
#Each time you do f.seek(0) you're moving to the start of the file. 
#Each time you do f.readline() you're reading a line from the file, and moving the read head to right after the \n that ends that file.

def rewind(p):
	p.seek(0)

def print_a_line(line_count, f):
	print line_count, f.readline()

current_file = open(input_file)
print "Lets print the whole file\n"
print_all(current_file)

print "Now let's rewind, kind of like a tape"
rewind(current_file)

print "Let's print 3 lines:\n"

current_line = 1
print_a_line(current_line, current_file)

current_line = current_line + 1
print_a_line(current_line, current_file)