from sys import argv

script, filename = argv

print "we are going to erase the file %s:" % filename
print "If you don't want that press Ctrl+C "
print "If you do want that hit RETURN"

raw_input("?")

print "Opening a file"
target = open(filename, 'w+')

print "Truncating the file Good bye"
target.truncate()

print "Now I am going to as three lines"

line1 = raw_input("Line 1:")
line2 = raw_input("Line 2:")
line3 = raw_input("Line 3:")

print "I am going to write these 3 lines to file"

target.write(line1)
target.write("\n")
target.write(line2)
target.write("\n")
target.write(line3)
target.write("\n")

target = open(filename)
print "contents of the File are:"
print target.read()

print "And finally we close it"
target.close()