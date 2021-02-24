#!/bin/bash
clear
resCode=${1:-200} ; shift 1
resDelay=${1:-100} ; shift 1
curl \
  -sv \
  'http://localhost:8080?resCode='$resCode'&resDelay='$resDelay \
  --data \
'line1
line2
'


