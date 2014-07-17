from sys import argv

script, user_name = argv
prompt = '>'

print "Hi %s, I am script %s" % (user_name, script)
print "I'd like to ask few questions"
print "Do you like me %s ?" % user_name
likes = raw_input(prompt)

print "Where do you live %s ?" % user_name
lives = raw_input(prompt)

print "What kind of computer do you have %s?" % user_name
computer = raw_input(prompt)

print """
Alright, so you %s about liking me.
you live in %r. Not sure where it is.
And you have %s computer. Nice""" % (likes, lives, computer)