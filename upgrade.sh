#!/bin/bash

readonly prefix='github.com/dobyte/due'
readonly len=${#prefix}

while read line
do
  if [[ ${line:0:len} = ${prefix} ]];then
    arr=($line)
    go get "${arr[0]}@latest"
  fi
done < go.mod