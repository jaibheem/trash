# this one is like your scripts with argv
def print_two(*args):
	arg1, arg2, arg3, arg4, arg5, arg6 = args
	print "arg1: %s arg2: %s %s %s %s %s" % (arg1, arg2, arg3, arg4, arg5, arg6)

# ok, that *args is actually pointless, we can just do this
def print_two_again(arg1, arg2):
	print "arg1:%s arg2: %s" % (arg1, arg2)

# this just takes one argument
def print_one(arg1):
	print "arg1:%s" % arg1

# this one takes no arguments
def print_nothing():
	print "I got nothing"

print_two("Hello", "HI", 3, 4, 5, 6)
print_two_again("Hello Again", "HI Again")
print_one("Print one")
print_nothing()


