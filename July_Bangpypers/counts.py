fruit_bag = ['apple' , 'orange', 'apple']
fruit_group = {}
for fruit in fruit_bag:
	fruit_group[fruit] = fruit_group.get(fruit, 0) + 1
print fruit_group

#Method2: try using defaultdict




