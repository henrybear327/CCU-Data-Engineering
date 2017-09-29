#include "SearchEngine.h"
#include "FileManager.h"
#include "UTF8Helper.h"

#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <sys/time.h>
#include <unistd.h>

void SearchEngine::performTextSearch()
{
    // TODO: add timer
    printf("Performing keyword searching...\n");

    for (int i = 0; i < (int)text.size(); i++) {
        // printf("%d: %d\n", i, text[i]);

        std::vector<int> candidate;
        for (int j = 0; i + j < (int)text.size() && j < 7; j++) {
            candidate.push_back(text[i + j]);
        }

        for (int j = (int)candidate.size(); j >= 2; j--) {
            // for (auto k : candidate) {
            //     printf("%d ", k);
            //     // getchar();
            // }
            // puts("");

            if (match.find(candidate) != match.end()) {
                // printf("Matched at %d: ", i);
                // {
                //     char location[111];
                //     sprintf(location, "\nMatched at %d: ", i);
                //     int ret =
                //         write(fileManager->resultHelper->fd, location,
                //         strlen(location));

                //     if (ret == -1) {
                //         perror("write() error");
                //         exit(1);
                //     }
                // }

                // for (auto k : candidate) {
                //     printf("%d\n", fileManager->keywordHelper->dictionary[k].bytes);
                //     {
                //         int ret = write(fileManager->resultHelper->fd,
                //                         fileManager->keywordHelper->dictionary[k].buffer,
                //                         fileManager->keywordHelper->dictionary[k].bytes);

                //         if (ret == -1) {
                //             perror("write() error");
                //             exit(1);
                //         }
                //     }
                // }

                match[candidate]++;
                i += (j - 1);
                break;
            }
            candidate.pop_back();
        }
    }

    printf("Done\n");
}

void SearchEngine::loadKeywords()
{
    printf("Loading keywords...\n");

    while (1) {
        bool terminate = false;
        std::vector<int> tmp;
        for (int i = 0, code = fileManager->keywordHelper->extractWord();
             code != 10; i++, code = fileManager->keywordHelper->extractWord()) {
            if (code == 0) {
                terminate = true;
                break;
            }

            tmp.push_back(code);
        }

        if (terminate)
            break;

        // printf("Keyword added\n");
        match[tmp] = 0;
    }

    printf("Done\n\n");
}

void SearchEngine::loadText()
{
    struct timeval starting, ending;
    int elapsedTime;
    gettimeofday(&starting, NULL);

    printf("Loading text...\n");

    for (int code = fileManager->textHelper->extractWord(); code != 0;
         code = fileManager->textHelper->extractWord()) {
        text.push_back(code);
    }

    gettimeofday(&ending, NULL);
    elapsedTime = (ending.tv_sec - starting.tv_sec) * 1000.0;    // sec to ms
    elapsedTime += (ending.tv_usec - starting.tv_usec) / 1000.0; // us to ms
    printf("%d.%03d\n", elapsedTime / 1000, elapsedTime % 1000);

    printf("Done\n\n");
}

void SearchEngine::printFrequencyList()
{
    for (auto res : match) {
        for (auto piece : res.first) {
            int ret = write(fileManager->resultHelper->fd,
                            fileManager->keywordHelper->dictionary[piece].buffer,
                            fileManager->keywordHelper->dictionary[piece].bytes);

            if (ret == -1) {
                perror("write() error");
                exit(1);
            }
        }

        {
            char output[111];
            sprintf(output, ": %d\n", res.second);

            int ret = write(fileManager->resultHelper->fd, output, strlen(output));

            if (ret == -1) {
                perror("write() error");
                exit(1);
            }
        }
    }
}

void SearchEngine::loadFilesToMemory()
{
    loadKeywords();
    loadText();
}