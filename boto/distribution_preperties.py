import boto
c = boto.connect_cloudfront()
all_distributions = c.get_all_distributions()
print all_distributions
for distribution in all_distributions:
	print distribution.domain_name, "-->", distribution.comment, "-->", distribution.origin, "-->", distribution.status
#print distribution_id.domain_name
