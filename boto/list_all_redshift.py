import boto
import boto.redshift

regions = ['us-east-1', 'us-west-2', 'eu-west-1', 'ap-southeast-1', 'ap-northeast-1', 'ap-southeast-2']
for region in regions:
	connection = boto.redshift.connect_to_region(region)
	clusters = connection.describe_clusters()['DescribeClustersResponse']['DescribeClustersResult']['Clusters']
	for cluster in clusters:
		print cluster['ClusterIdentifier']
