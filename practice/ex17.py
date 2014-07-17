from sys import argv
from os.path import exists

script, from_file, to_file = argv
print "Copying from %s to %s" % (from_file, to_file)

in_file = open(from_file)
indata = in_file.read()

print "The input file is %s bytes long" % len(indata)

print "Does the output files exists? %r" %exists(to_file)
print "Hit ENTER to continue and CTRL-C to abort"
raw_input()

out_file = open(to_file, 'w')
out_file.write(indata)

print "All right, all done"

out_file.close()
in_file.close()
