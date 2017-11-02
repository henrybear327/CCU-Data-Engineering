#!/bin/bash

go build -o generator generator.go

for (( i=1; i<=1000000000; i*=10))
do
	echo "generating $i.in"
	(time ./generator -n $i -o /tmp/testcases/$i.in) 2>&1 | tee /tmp/testcases/log/$i.in.output
	(time sort -d -o /tmp/testcases/$i.out /tmp/testcases/$i.in) 2>&1 | tee /tmp/testcases/log/$i.out.output
	echo "done!"
done

