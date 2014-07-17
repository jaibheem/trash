import boto
connect = boto.connect_cloudfront()
#print connect
print "Enter the invalidation ID:"
inval_id = raw_input (">")
inval_req = connect.invalidation_request_status(u'CFID', inval_id)
print "Status of invalidate request is:", inval_req.status
	
