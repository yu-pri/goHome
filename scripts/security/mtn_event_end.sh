#!/bin/bash
echo "motion ended $1"

find /tmp/motion/* -mtime +1 -exec rm {} \;
nohup su -c "rsync -avzhe 'ssh -p 22095' /tmp/motion  al@office.webdevs.com:/home/al/" - al &
