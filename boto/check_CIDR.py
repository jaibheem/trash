__author__ = 'Jaibheemasen'
import boto
import boto.vpc
import boto.ec2

regions = ['us-east-1', 'us-west-2', 'eu-west-1', 'ap-southeast-1', 'ap-northeast-1', 'ap-southeast-2']
for region in regions:
    connection = boto.vpc.connect_to_region(region)
    all_sg = connection.get_all_security_groups()
    for each_sg in all_sg:
        print "\n", region, "-->", each_sg.name, "-->:"
        for rule in each_sg.rules:
            print rule.grants