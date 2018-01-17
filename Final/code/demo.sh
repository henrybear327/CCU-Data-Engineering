#!/bin/bash

time ./nfa -pat "(-?(0|1|2|3|4|5|6|7|8|9)+)(\.(0|1|2|3|4|5|6|7|8|9)+)?" -str "-0.12465"  2>/dev/null
echo -e "================================================================================\n\n\n"

time ./nfa -pat "(q|w|e|r|t|y|u|i|o|p|a|s|d|f|g|h|j|k|l|z|x|c|v|b|n|m)+://.*\.htm" -str "https://deerchao.net/tutorials/regex/common.htm" 2>/dev/null
echo -e "================================================================================\n\n\n"

time ./nfa -pat "a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" -str "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" 2>/dev/null
echo -e "================================================================================\n\n\n"

time ./nfa -pat "別噢 (沒事)*" -str "別噢 沒事沒事" 2>/dev/null
echo -e "================================================================================\n\n\n"

time ./nfa -pat "哇哈* (A|a)mos loves (to say)?喂*.?" -str "哇哈哈 Amos loves 喂喂喂喂." 2>/dev/null
echo -e "================================================================================\n\n\n"

time ./nfa -pat '\a?\' -str '\\' 2>/dev/null
echo -e "================================================================================\n\n\n"

time ./nfa -pat 'aabb' -str 'aabbbb' 2>/dev/null
echo -e "================================================================================\n\n\n"

time ./nfa -pat '(abcdef)*' -str '' 2>/dev/null
echo -e "================================================================================\n\n\n"

time ./nfa -pat '(abcdef)*' -str 'abcdefabcdefabcdefabcdef' 2>/dev/null
echo -e "================================================================================\n\n\n"

