#!/bin/bash
CURRENT=$(df / | grep / | awk '{ print $5}' | sed 's/%//g')
THRESHOLD=75

if [ "$CURRENT" -gt "$THRESHOLD" ] ; then
    echo "Subject: `hostname` is currently using : $CURRENT%" > mail.txt
    sudo /usr/sbin/sendmail -F admin@gmail.com -t jaibheemsen@gmail.com < mail.txt
fi
