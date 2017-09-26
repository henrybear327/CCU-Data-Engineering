CC=/usr/local/Cellar/gcc/7.2.0/bin/g++-7
CXXFLAGS=-Wall -Wextra -Wshadow -std=c++11

all: main

main: main.cpp
	$(CC) $(CXXFLAGS) main.cpp -o main

format: *.cpp
    clang-format -style=LLVM -i *.cpp
    astyle --style=linux *.cpp

clean: main
	rm main
