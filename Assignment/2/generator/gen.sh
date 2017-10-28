#!/bin/bash

go build -o generator generator.go

for (( i=1; i<=1000000000; i*=10))
do
	echo "generating $i.in"
	time ./generator -n $i -o /tmp/testcases/$i.in
	echo "done!"
done
