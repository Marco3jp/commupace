#!/bin/bash

if [ $# -ne 1 ]; then
  echo "sshfsのマウントポイントを指定してください (./mount.sh ~/Desktop/mount_point)" 1>&2
  exit 1
fi

sshfs vagrant@10.1.2.100:/vagrant $1  -o reconnect
