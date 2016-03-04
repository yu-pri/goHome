#!/bin/bash
echo "save file $1"
ls -l $1
nohup su -c "rsync -avzhe 'ssh -p 22095' $1 al@office.webdevs.com:/home/al/motion" - al &
