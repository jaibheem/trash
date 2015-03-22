__author__ = 'Jaibheemasen'
import boto
import boto.ec2
import os
import subprocess
from subprocess import Popen

regions = ['us-east-1', 'us-west-2', 'eu-west-1', 'ap-southeast-1', 'ap-northeast-1', 'ap-southeast-2']
for region in regions:
    connection = boto.ec2.connect_to_region(region)
    all_instances = connection.get_all_instance_status()
    for instance in all_instances:
        if instance.events:
	    print instance.id
