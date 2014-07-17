#! /usr/bin/env python

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
    # Write the poem to a file 'road-not-taken.txt'
    poem = __doc__
    #Method 1:
    f = open("road-not-taken", "w")
    f.write(poem)
    f.close()
    #method 2:
    with open('road-not-taken.txt', "w") as f:
        f.write(poem)

    # Step 2
    # Read back the poem from the file and print it.
    #Method 1:
    with open('road-not-taken') as f:
        print f.read()
    #Method 2:
    with open('road-not-taken') as f:
        for line in f:
            print line, # we use , for continual printing


# __name__ holds the name of the current module
if __name__ == "__main__":
    main() # call main function
