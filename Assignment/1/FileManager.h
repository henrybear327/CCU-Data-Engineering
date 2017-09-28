#ifndef FILE_MANAGER_H
#define FILE_MANAGER_H

#include "UTF8Helper.h"
#include <string>

#define CURRENTDIRECTORYBUFFERSIZE 10000

struct FileManager {
public:
    FileManager(std::string _keywordFile, std::string _textFile,
                std::string _resultFile);

    UTF8Helper *keywordHelper;
    UTF8Helper *textHelper;
    UTF8Helper *resultHelper;

private:
    std::string keywordFile;
    std::string textFile;
    std::string resultFile;
};
#endif
