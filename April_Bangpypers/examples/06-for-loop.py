#! /usr/bin/env python
# -*- coding: utf-8 -*-

"""Write a program to display the peom The Road Not Taken by Robert Frost.

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
Two roads diverged in a wood, and I—
I took the one less traveled by,
And that has made all the difference"""

# Below is function definition
def main():
    poem = __doc__
    # Step 1
    # Find total number of blank lines in the poem
    # Method 1:
    lines = poem.split("\n")
    
    #print lines
    valid_lines, blank_lines = 0, 0
    for line in lines:
        if line:
            valid_lines += 1
        else:
            blank_lines += 1
    print valid_lines, blank_lines
    #method 2:
    #print lines
    count = 0
    for line in lines:
        if not line:
            count += 1
    print count

    # Step 2
    # use sum function
    # Find total number of words in the poem
    word_count = 0
    for line in lines:
        if line:
            for word in line.split(" "):
                if word:
                    word_count += 1
    print word_count

    # Step 3
    # Display first 5 lines in the poem
    index = 0
    while index < 5:
        print lines[index]
        index += 1

    #Step 3 --> Doing the same above operation using Slicing
    #syntax : name[start:stop:step]
    print lines[0:5]

    # Step 4:
    # Display lines till first line break
    #using while:
    index = 0
    while lines[index] != "":
        print lines[index]
        index += 1
    #using for loop:
    for line in lines:
        if not line:
            break
        print line

# __name__ holds the name of the current module
if __name__ == "__main__":
    main() # call main function
