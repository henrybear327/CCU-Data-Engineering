/*
UTF8 reader

Support high level abstraction for reading in the utf8 word, hashing utf8 word

Additionally support checking for valid utf8 text when compiled with
CHECKINTEGRITY

The longest utf8 encoding after stripping uses only 21 bits, so it's more than
enough to fit in an int

Procedure
Check for validility
Determine type of utf8
Load it into an int
Perform whatever you want
*/

#ifndef UTF8_READER_H
#define UTF8_READER_H

#endif