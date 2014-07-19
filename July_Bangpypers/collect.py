import collections
d = collections.defaultdict(int)
print d[1]


#Method 3:
fruit_bag = ['apple' , 'orange', 'apple']

fruit_collections = collections.Counter(fruit_bag)
print fruit_collections
dir(fruit_collections)