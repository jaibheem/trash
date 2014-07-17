import boto
connect = boto.connect_cloudfront()
#print connect
all_invals = connect.get_invalidation_requests(u'CFID')
#print all_invals
print "Status of all invalidation requests that are in progress are:"
for inval in all_invals:
	#if inval.status == "InProgress":
		print "Invalidation ID:%s, Invalidation Status: %s" % (inval.id, inval.status)
	#elif inval.status == "Completed":
	#	exit
