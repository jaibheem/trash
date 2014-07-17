#! /usr/bin/env python
# -*- coding: utf-8 -*-

"""
Write a program to solve (a + b) * ( a + b).
"""

# Below is function definition
def main():
    # Create variables a and b.
    # Use the formula a * a + 2 * a * b + b * b
    # print the result back
    a, b = 2, 3
    result = a * a + 2 * a * b + b * b
    print "Result of (%d + %d) ^ 2 = %d" %(a, b, result)
    print "Result of ({} + {}) ^ 2 = {}".format(a, b, result)

# __name__ holds the name of the current module
if __name__ == "__main__":
    main() # call main function
