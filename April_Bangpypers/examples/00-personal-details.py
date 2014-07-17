#! /usr/bin/env python

"""
Write a program to display your details like name, age, address.
"""

# Below is function definition
def main():
    """
    Create three variables name, age, address.
    Hint:
    age = 24
    """
    name, age = "Jai", 24
    #address = "Jai\n Bangalore\n Karnataka"
    address = """
    Jai
    Bangalore
    Karnataka
    """
    print name, age, address, type(age)
 
# __name__ holds the name of the current module
if __name__ == "__main__":
    main() # call main function
