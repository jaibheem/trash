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

def is_sonnet(poem):
    """Return the poem is sonnet or not
    """
    lines = poem.split("\n")
    #Method 1:
    #words = []
    #for line in lines:
    #    if line: #Non emplty
    #        words.append(line)
    #Method 2:
    cleaned_lines = [line for line in lines if line]
    #print words
    return len(cleaned_lines) == 14


def total_words(poem):
    """Function returns total number of words in the poem
    """
    lines = poem.split("\n")
    word_count = 0
    for line in lines:
        if line:
            for word in line.split(" "):
                if word:
                    word_count += 1
    return word_count


def word_count(poem):
    """Returns dictionary containing words as key and count as value
    """
    lines = poem.split("\n")
    cleaned_lines = [line for line in lines if line] # To get cleaned lines (list comprehention)
    word_frequency = {}
    for line in cleaned_lines:
        words = [word for word in line.split(" ") if word] # to get cleaned words
    #print words
        for word in words:
            if word in word_frequency:
                word_frequency[word] += 1
            else:
                word_frequency[word] = 1
    return word_frequency

def unique_words(poem):
    """Return unique words in the poem as a set
    """
    words_count = word_count(poem)
    return words_count.keys()

def find_total_occurrences(poem, *words):
    """Find total occurences of the words in the poem.
    """
    
    word_frequency = word_count(poem)
    #Method 1:
    #counts = {}
    #for key in words:
    #    counts[key] = word_frequency.get(key, 0)
    #return counts

    #Method 2: Dictionary comprehention
    return {key: word_frequency.get(key, 0) for key in words}

def main():
    poem = __doc__

    # Step 1
    # Edit the function total_words to return total words in the poem
    total_no_of_words = total_words(poem)
    print total_no_of_words

    # Step 2
    # Find the poem is sonnet or not ? check if it has 14 lines or not
    print "Is poem is Sonnet: {}".format(is_sonnet(poem))

    # Step 3
    # Find the word frequency in the peom using word_count
    frequency = word_count(poem)
    print frequency

    #To print all the words in poem.
    print unique_words(poem)


    # Step 4
    # Find unique words in the poem using unique_words
    print find_total_occurrences(poem, "road", "Jai", "I")
    print find_total_occurrences(poem, "forst")


    # Step 5
    # Find total occurences of few words

# __name__ holds the name of the current module
if __name__ == "__main__":
    main() # call main function
