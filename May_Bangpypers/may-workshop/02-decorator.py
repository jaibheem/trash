# -*- coding: utf-8 -*-
import sys
# Write a function which returns greeting message depending on usernames
def greeting(*names):
	msg = ""
	for name in names:
		msg += "Welcome {}\n".format(name)
	return msg

def test_greeting():
	result = greeting("Jai")
	assert "Jai" in result
	print "Success"
	result = greeting ("Jai", "sri")
	assert "Jai" in result
	assert "sri" in result
	print "success"


# Write a function which returns greeting message depending on usernames
# and number of times greetings needs to be printed

def greeting_times(**kwagrs):
	msg = ""
	for name, times in kwagrs.items():
		for time in range(times):
			msg += "Welcome {}\n".format(name)
	return msg
# Python scopes locals(), globals(), nested functions.

# Write a function called cache which will cache factorial result.
# When factorial is called with same number argument as previous call
# fetch result from cache.
def cache(f):
	cache_results = {}
	def inner(n):
		if n in cache_results:
			print "from Cache"
			return cache_results.get(n, 0)
		else:
			cache_results[n] = f(n)
			return cache_results.get(n)
	return inner

@cache
def factorial(n):
	if n<= 1:
		return 1
	else:
		return n * factorial(n-1)

@cache
def square(n):
	return n * n

# test

# Write a function called login which will take a parameter username.
# Write a decorator called login_required which is on top of login.
# login_decorator will look into listed of predefined username and
# return True or False
def login_required(n):
	usernames = ["Jai", "sri", "test"]
	def inner(username):
		if username in usernames:
			return n(username)
		else:
			return False
	return inner

@login_required
def login(username):
	print "Yet to be implemented"
	return True

def test():
	print "Started test_greeting"
	test_greeting()

def main():
    #print greeting("Jai")
    #print greeting("Jai", "sri")
    #print greeting_times(Jai=5)
    #res = factorial(5)
    #print res
    print factorial(5)
    print square(2)
    res = login("Jai")
    print res


if __name__ == "__main__":
	if "test" in sys.argv:
		test()
		exit(0)
	main()
