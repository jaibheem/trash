# -*- coding: utf-8 -*-
import sys

# Write a function to square a given number

def square(number):
    return number * number

# test
def test_square():
    assert square(4) == 16
    print "Success 1"
    assert square(2) != 23, "Square of 2 is not 23"
    print "Success 2"

# Write a function to cube a number
def cube(square_fun, number):
    return square_fun(number) * number

# test
def test_cube():
    assert cube(square, 2) == 8
    print "Cube success 1"
    assert cube(square, -1) == -1
    print "Negative Success"

# Given a list of numbers return list of squares.
def list_of_squares(numbers):
    #with Map
    return map(square, numbers)
    #without Map:
    results = []
    for no in numbers:
        results.append(square(no))
    return results
    #Using list comprohension:
    #results = [square(no) for no in numbers]

# test
def test_list_of_squares():
    assert list_of_squares([1, 2, 4]) == [1, 4, 16]
    print "list_success"
    assert list_of_squares([-1]) == [1]
    print "list_success"
# Given a list of numbers filter even numbers
def is_even(no):
    return no % 2 == 0

def get_list_of_evens(nos):
    return filter(is_even, nos)

# test
def test_get_list_of_evens():
    assert get_list_of_evens([1, 2]) == [2]
    print "success"
    assert get_list_of_evens([1, 3]) == []
    print "success"
    assert get_list_of_evens([1, 2, 4]) == [2, 4]
    print "success"

# Combine map, filter, sum.
# Given a list of numbers, find all even numbers and find sum of squares
def square_of_evens(nos):
    evens = filter(is_even_number, nos)
    squares = map(square, evens)
    return sum(squares)

def is_even_number(no):
    return no % 2 == 0

#Test Case:
def test_square_of_evens():
    assert square_of_evens([1, 2, 3, 4]) == 20
    print "Success"
    assert square_of_evens([1, 3, 5]) == 0
    print "0 Success"
    assert square_of_evens([]) == 0
    print "Empty success"

# Test
def test():
    print "Started test_square"
    test_square()
    print "Started test_cube"
    test_cube()
    print "Started test_list_of_squares"
    test_list_of_squares()
    print "Started test_get_list_of_evens"
    test_get_list_of_evens()
    print "Started test_square_of_evens"
    test_square_of_evens()

def main():
    result = square(23) 
    print "Square of the given number is", result
    #print result
    print(list_of_squares([1, 2, 3, 4]))
    print get_list_of_evens([1, 2, 3, 4, 5, 6, 7])
    print "Sum of even square"
    print square_of_evens([1, 2, 3, 4])

if __name__ == "__main__":
    #=> Check whether test is present in args
    #print globals()
    if 'test' in sys.argv:
        test()
        exit(0)
    main()
