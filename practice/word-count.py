#! /usr/bin/env python

"""
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

def word_count(poem):
	#First you need to get the lines, non-empty lines
	lines = poem.split("\n")
	#Now display all the lines, non-emplty line
	count = 0
	for line in lines:
		if line: #This will eliminate empty lines
			#print line
			#split each line into words
			for word in line.split(" "): #split each line in to words
				if word: #this will eliminate empty words
					count += 1
				#print word
	return count

def word_frequency(poem):
	lines = poem.split("\n") # split lines
	empty_dict = {} #Take a empty dict, later we keep adding the key and values here
	for line in lines: 
		if line: #for every non-emplty line
			for word in line.split(" "):
				if word: #split into non-empty words
					if word in empty_dict: #see if word is in empty_dict, (see else block)
						empty_dict[word] += 1 #then increments the value of the key by 1 for every next occurance
					else:
						empty_dict[word] = 1 #first time it goes to else and assign key as word and values as 1.(if)
	return empty_dict



def main():
	#Excercise : use word_count function to count the number of words.
	#Hint : First split poem into lines, lines into words, and then iterate through them and count
	poem = __doc__
	#print poem
	print "Total Number of words in the poem are: {}".format(word_count(poem))
	#word_count = word_count(poem)
	print "Word Frequency:\n", word_frequency(poem)
	#print "List of words:", word_frequency(poem).keys()
	#print word_frequency(poem).items()

if __name__ == "__main__":
	main()
