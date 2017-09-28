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
    searchEngine.performTextSearch();

    for (auto res : searchEngine.match) {
        for (auto piece : res.first) {
            int ret = write(
                          searchEngine.fileManager->resultHelper->fd,
                          searchEngine.fileManager->keywordHelper->dictionary[piece].buffer,
                          searchEngine.fileManager->keywordHelper->dictionary[piece].bytes);

            if (ret == -1) {
                perror("write() error");
                exit(1);
            }
        }

        {
            char output[111];
            sprintf(output, ": %d\n", res.second);

            int ret = write(searchEngine.fileManager->resultHelper->fd, output,
                            strlen(output));

            if (ret == -1) {
                perror("write() error");
                exit(1);
            }
        }
    }

    return 0;
}
