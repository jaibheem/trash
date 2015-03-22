__author__ = 'Jaibheemasen'
import boto
import boto.ec2.elb
import boto.vpc

regions = ['us-east-1', 'us-west-2', 'eu-west-1', 'ap-southeast-1', 'ap-northeast-1', 'ap-southeast-2']
for region in regions:
    connection = boto.ec2.elb.connect_to_region(region)
    all_lb = connection.get_all_load_balancers()
    for lb in all_lb:
        print lb.dns_name, "\t", connection.get_all_lb_attributes(lb.name).connecting_settings.idle_timeout
