#ifndef FILE_MANAGER_H
#define FILE_MANAGER_H

#include <string>

#define CURRENTDIRECTORYBUFFERSIZE 10000

struct FileManager {
public:
    FileManager(std::string _keywordFile, std::string _textFile,
                std::string _resultFile);

private:
    std::string keywordFile;
    std::string textFile;
    std::string resultFile;

    int keyword_fd;
    int text_fd;
    int result_fd;
};
#endif
