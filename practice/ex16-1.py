from sys import argv
script, filename = argv

txt = open(filename)
print "contents of the file %s are:" %filename
print txt.read()