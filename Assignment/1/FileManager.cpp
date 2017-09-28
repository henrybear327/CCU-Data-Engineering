#include "FileManager.h"
#include "color.h"
#include <cstdio>
#include <cstdlib>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>

FileManager::FileManager(std::string _keywordFile, std::string _textFile,
                         std::string _resultFile)
{
    char currentDirectory[CURRENTDIRECTORYBUFFERSIZE];
    printf(ANSI_COLOR_CYAN "Current directory is %s\n\n",
           getcwd(currentDirectory, CURRENTDIRECTORYBUFFERSIZE));

    // set path
    std::string directoryString(currentDirectory);
    keywordFile = directoryString + "/" + _keywordFile;
    textFile = directoryString + "/" + _textFile;
    resultFile = directoryString + "/" + _resultFile;

    printf("Files to be loaded:\n");
    printf("Keyword: %s\n", keywordFile.c_str());
    printf("Text: %s\n", textFile.c_str());
    printf("Result: %s\n", resultFile.c_str());

    // open fd
    keyword_fd = open(keywordFile.c_str(), O_RDONLY);
    if (keyword_fd == -1) {
        perror(ANSI_COLOR_RED "Error opening keyword file");
        exit(1);
    }

    text_fd = open(textFile.c_str(), O_RDONLY);
    if (text_fd == -1) {
        perror(ANSI_COLOR_RED "Error opening text file");
        exit(1);
    }

    result_fd = open(resultFile.c_str(), O_WRONLY | O_CREAT | O_TRUNC, 0600);
    if (result_fd == -1) {
        perror(ANSI_COLOR_RED "Error creating result file");
        exit(1);
    }

    printf("Done loading files\n\n" ANSI_COLOR_RESET);
}