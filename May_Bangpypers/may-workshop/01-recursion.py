# -*- coding: utf-8 -*-
import sys
# Write a function flatten which will flatten the nested list
def flatten(nos, result=None):
	#print "result:{}".format(result)
	if result is None:
		result = []
	#print nos, result

	for no in nos:
		#print no
		if isinstance(no, list):
			#print no
			flatten(no, result)
		else:
			result.append(no)

	return result

# Test
def test_flaten():
	assert flatten([1, 2]) == [1, 2]
	print "Success"
	assert flatten([]) == []
	print "Success"

# write a function which will compute factorial of n, using recursion
def factorial(no):
	#print no
	if no <= 1:
		return 1
	else:
		return no * factorial(no - 1)
#Test Case:
def test_factorial():
	assert factorial(1) == 1
	print "Success"
	assert factorial(5) == 120
	print "Success"
	assert factorial(-1) == 1
	print "Success"
	assert factorial(0) == 1
	print "Success"
 
# Test
def test():
	print "started test_flaten"
	test_flaten()
	print "started test_factorial"
	test_factorial()

def main():
    print flatten([1, 2, [3, 4]])
    print factorial(5)


if __name__ == "__main__":
	if 'test' in sys.argv:
		test()
		exit(0)
	main()
