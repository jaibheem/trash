#! /usr/bin/env python
import os
import sys
"""
Write a program to display the peom The Road Not Taken by Robert Frost.

Two roads diverged in a yellow wood,
And sorry I could not travel both
And be one traveler, long I stood
And looked down one as far as I could
To where it bent in the undergrowth;

Then took the other, as just as fair,
And having perhaps the better claim,
Because it was grassy and wanted wear;
Though as for that the passing there
Had worn them really about the same,

And both that morning equally lay
In leaves no step had trodden black.
Oh, I kept the first for another day!
Yet knowing how way leads on to way,
I doubted if I should ever come back.

I shall be telling this with a sigh
Somewhere ages and ages hence:
Two roads diverged in a wood, and I--
I took the one less traveled by,
And that has made all the difference
"""

def main():
    # Step 1
    # Check whether 'road-not-taken.txt' file is present in current directory
    print os.path.exists("road-not-taken")
    #print sys.argv

    nos = sys.argv[1:] # we don't want the program name which will be sys.argv[0]
    #print nos
    print sum([int(no) for no in nos])
    # Step 2
    # Get list of words as command line arguments and check its presence


# __name__ holds the name of the current module
if __name__ == "__main__":
    main() # call main function
