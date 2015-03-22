import boto
connection = boto.connect_s3()
bucket = connection.get_bucket('bucketname')
bucket_entries = bucket.list(prefix='backup/')
for entry in bucket_entries:
	print entry
