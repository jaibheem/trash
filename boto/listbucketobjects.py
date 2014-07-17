import boto
from boto.s3.connection import S3Connection

connection = boto.connect_s3()
bucket_name = raw_input("Enter the bucket name:\n>")
bucket = connection.get_bucket(bucket_name)

for key in bucket.list():
	print key.name.encode('utf-8')
