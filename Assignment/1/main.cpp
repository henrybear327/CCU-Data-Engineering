#include "FileManager.h"
#include "SearchEngine.h"
#include <cstring>
#include <string>

using namespace std;

/*
Configuration
*/

// std::string keywordFile = "./testcase/small/keyword.txt";
// std::string textFile = "./testcase/small/text.txt";
// std::string resultFile = "./testcase/small/result.txt";

std::string keywordFile = "./testcase/large/term.txt";
std::string textFile = "./testcase/large/doc.txt";
std::string resultFile = "./testcase/large/myResult.txt";

int main()
{
    FileManager fileManager(keywordFile, textFile, resultFile);

    SearchEngine searchEngine(&fileManager);
    searchEngine.loadFilesToMemory();
    searchEngine.performTextSearch();
    searchEngine.printFrequencyList();

    return 0;
}
