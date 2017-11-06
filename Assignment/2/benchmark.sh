#!/bin/bash

go build -o my_sort *.go 

mkdir -p /tmp/testcases/my
mkdir -p /tmp/testcases/my/log

CHUNKS=100
DEPTH=4

for ((i=1; i<=1000000000; i*=10))
do
	echo "benchmarking $i.in"
	# (time ./my_sort -i /tmp/testcases/$i.in -o /tmp/testcases/my/$i.nosort.out -chunks $CHUNKS -depth=$DEPTH -freq=1 -p=0) 2>&1 | tee /tmp/testcases/my/log/$i.nosort.out.output
	# rm /tmp/testcases/my/$i.nosort.out
	# (time ./my_sort -i /tmp/testcases/$i.in -o /tmp/testcases/my/$i.single.out -chunks $CHUNKS -depth=$DEPTH -freq=1 -p=1 && sort -c /tmp/testcases/my/$i.single.out) 2>&1 | tee /tmp/testcases/my/log/$i.single.out.output
	# rm /tmp/testcases/my/$i.single.out
	# (time ./my_sort -i /tmp/testcases/$i.in -o /tmp/testcases/my/$i.parallel.out -chunks $CHUNKS -depth=$DEPTH -freq=1 -p=2 && sort -c /tmp/testcases/my/$i.parallel.out) 2>&1 | tee /tmp/testcases/my/log/$i.parallel.out.output
	# rm /tmp/testcases/my/$i.parallel.out
	(time sort -o /tmp/testcases/my/$i.builtinsort.out /tmp/testcases/$i.in) 2>&1 | tee /tmp/testcases/my/log/$i.builtinsort.out.output
	rm /tmp/testcases/my/$i.builtinsort.out
	echo "done!"
done

