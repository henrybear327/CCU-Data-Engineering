build: external_merge_sort.go winner_tree.go
	go build -o my_sort external_merge_sort.go winner_tree.go

test_1000000: build
	time ./my_sort -i ./testcase/number_1000000.in -tmp=./tmp -chunks=100

test_10000000: build
	time ./my_sort -i ./testcase/number_10000000.in -tmp=./tmp -chunks=10