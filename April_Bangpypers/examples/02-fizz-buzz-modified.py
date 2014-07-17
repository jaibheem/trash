#! /usr/bin/env python
# -*- coding: utf-8 -*-

"""
Write a program to print fizz, buzz or fizzbuzz if one of the conditions is meet.

If the number is a multiple of 3, print fizz; a multiple of 5, print buzz;
and if a multiple of both 3 and 5, print fizzbuzz.
"""

# Below is function definition
def main():
    # Create variable named num.
    # HINT:
    # if cond:
    #     pass
    # elif cond:
    #     pass
    # else:
    #     pass
    number = 15
    if number % 3 == 0 and number % 5 == 0:
        print "Fizzbuzz"
    elif number % 5 == 0:
        print "buzz"
    elif number % 3 == 0:
        print "fizz"
    else:
        print "{} didn't satify any condition".format(number)

# __name__ holds the name of the current module
if __name__ == "__main__":
    main() # call main function
