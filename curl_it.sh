#!/bin/bash
clear
curl \
  -sv \
  'http://localhost:8080?resCode=201&resDelay=101' \
  --data \
'line1
line2
'


