#include "FileManager.h"
#include "SearchEngine.h"
#include <string>

using namespace std;

/*
Configuration
*/

std::string keywordFile = "./testcase/small/keyword.txt";
std::string textFile = "./testcase/small/text.txt";
std::string resultFile = "./testcase/small/result.txt";

int main()
{
    FileManager fileManager(keywordFile, textFile, resultFile);

    // SearchEngine searchEngine(&fileManager);
    // searchEngine.performTextSearch();

    return 0;
}
