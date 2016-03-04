#!/bin/bash
echo "camera lost $1"
nohup /opt/bin/send_motion &
