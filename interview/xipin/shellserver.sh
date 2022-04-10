#!/bin/sh
echo "udp server start at 127.0.0.1:16523"
while true;do
  nc -lu 16523 -w 1 | (read line;
  echo $line)
done
