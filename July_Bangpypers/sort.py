topsites = []
#sort the list by 1st index and disply, if you want to sort with 0th index use operator.itemgetter(0)
[x[0] for x in sorted(topsites.items(), key=operator.itemgetter(1))]

