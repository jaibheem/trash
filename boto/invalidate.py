import boto
c = boto.connect_cloudfront()
inval_file = raw_input ("Enter file for invalidation\n>")
#print 'Enter file name for invalidation'
#We can hard code the invalidation file names in the code using the below sytax:
#paths = ['/path/to/file1.html', '/path/to/file2.html', ...]
#Warning Each CloudFront invalidation request can only specify up to 1000 paths. If you need to invalidate more than 1000 paths you will need to split up the paths into groups of 1000 or less and create multiple invalidation requests.
paths = [inval_file]
print paths
inval_req = c.create_invalidation_request("cloudfrontID", paths)
print inval_req
print inval_req.id
print inval_req.paths
