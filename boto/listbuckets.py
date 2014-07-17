import boto
from tabulate import tabulate
connection = boto.connect_s3()
print connection
print tabulate([["Bucket Name","Creation Date"]])
for bucket in connection.get_all_buckets():
	#table =  [bucket.name, bucket.creation_date]
	#print table
	print bucket.name, bucket.creation_date
