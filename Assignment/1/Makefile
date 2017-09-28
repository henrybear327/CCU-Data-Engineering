CC=/usr/local/Cellar/gcc/7.2.0/bin/g++-7
CXXFLAGS=-Wall -Wextra -Wshadow -std=c++11

all: searchEngine fileManager main
	$(CC) $(CXXFLAGS) *.o -o main

main: main.cpp
	$(CC) $(CXXFLAGS) -c main.cpp

searchEngine: SearchEngine.cpp
	$(CC) $(CXXFLAGS) -c SearchEngine.cpp

fileManager: FileManager.h
	$(CC) $(CXXFLAGS) -c FileManager.h

utf8reader: UTF8Reader.h UTF8Reader.cpp
	
format: 
    clang-format -style=LLVM -i *.cpp *.h
    astyle --style=linux *.cpp *.h

clean: 
	rm -f *.o
	rm -f main
