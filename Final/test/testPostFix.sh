#!/bin/bash

Red='\033[0;31m'          # Red
Green='\033[0;32m'        # Green
Yellow='\033[0;33m'       # Yellow
Blue='\033[0;34m'         # Blue
NC='\033[0m'              # No Color

# make my program
echo "Making my program..."
cd ../code && time make

# build answer
echo "Building answer nfa"
cd ../Sample && time gcc -O2 nfa.c -o nfa

echo ""

echo "Running testcases"
TESTCASES=( "a?" "(a?b)|c" "(a?b)|c|d" "abcdef" "(ab)+(cd)?" "((ab)+|(cd)?)a" "abc(a?b*c+)|(ccc+)|(c+a*a?)cba")
PATTERN="asdfghjkkk"
for i in "${TESTCASES[@]}"
do
    # echo "../code/nfa -pat $i -str $PATTERN"
    MY_OUTPUT="$(../code/nfa -pat $i -str $PATTERN)"
    echo $MY_OUTPUT

    # echo "../Sample/nfa $i $PATTERN"
    ANS_OUTPUT="$(../Sample/nfa $i $PATTERN)"
    echo $ANS_OUTPUT

    if [ "$MY_OUTPUT" == "$ANS_OUTPUT" ]; then
        printf "${Green}ACCEPTED${NC}\n"
    else
        printf "${Red}Wrong answer${NC}\n"
    fi

    echo ""
done