#!/bin/bash

go build -o my_sort *.go 

mkdir -p /tmp/testcases/my
mkdir -p /tmp/testcases/my/log

for (( i=1; i<=1000000000; i*=10))
do
	echo "generating $i.in"
	(time ./my_sort -i /tmp/testcases/$i.in -o /tmp/testcases/my/$i.nosort.out -chunks 100 -depth=6 -freq=1 -p=0) 2>&1 | tee /tmp/testcases/my/log/$i.nosort.out.output
	(time ./my_sort -i /tmp/testcases/$i.in -o /tmp/testcases/my/$i.single.out -chunks 100 -depth=6 -freq=1 -p=1 && sort -c /tmp/testcases/my/$i.single.out) 2>&1 | tee /tmp/testcases/my/log/$i.single.out.output
	(time ./my_sort -i /tmp/testcases/$i.in -o /tmp/testcases/my/$i.parallel.out -chunks 100 -depth=6 -freq=1 -p=2 && sort -c /tmp/testcases/my/$i.parallel.out) 2>&1 | tee /tmp/testcases/my/log/$i.parallel.out.output
	echo "done!"
done

