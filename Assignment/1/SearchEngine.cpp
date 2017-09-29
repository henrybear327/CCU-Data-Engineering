#include "SearchEngine.h"
#include "FileManager.h"
#include "UTF8Helper.h"

#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <string>
#include <sys/time.h>
#include <unistd.h>

void SearchEngine::performTextSearch()
{
    struct timeval starting, ending;
    int elapsedTime;
    gettimeofday(&starting, NULL);

    printf("Performing keyword searching...\n");

    for (int i = 0; i < (int)text.size(); i++) {
// printf("%d: %d\n", i, text[i]);
#ifdef VECTORBASED
        std::vector<int> candidate;
#endif

#ifdef STRINGBASED
        std::string candidate;
#endif
        for (int j = 0; i + j < (int)text.size() && j < 7; j++) {
#ifdef VECTORBASED
            candidate.push_back(text[i + j]);
#endif

#ifdef STRINGBASED
            // candidate += std::to_string(text[i + j]);
            candidate += text[i + j];
#endif
        }

        for (int j = (int)candidate.size(); j >= 2; j--) {
            // for (auto k : candidate) {
            //     printf("%d ", k);
            //     // getchar();
            // }
            // puts("");

            std::map<std::vector<int>, int>::iterator it = match.find(candidate);
            if (it != match.end()) {
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

                // match[candidate]++;
                it->second++;
                i += (j - 1);
                break;
            }
            candidate.pop_back();
        }
    }

    printf("Done\n");

    gettimeofday(&ending, NULL);
    elapsedTime = (ending.tv_sec - starting.tv_sec) * 1000.0;    // sec to ms
    elapsedTime += (ending.tv_usec - starting.tv_usec) / 1000.0; // us to ms
    printf("%d.%03d\n", elapsedTime / 1000, elapsedTime % 1000);
}

void SearchEngine::loadKeywords()
{
    printf("Loading keywords...\n");

    while (1) {
        bool terminate = false;

#ifdef VECTORBASED
        std::vector<int> tmp;
#endif

#ifdef STRINGBASED
        std::string tmp;
#endif

        for (int i = 0, code = fileManager->keywordHelper->extractWord();
             code != 10; i++, code = fileManager->keywordHelper->extractWord()) {
            if (code == 0) {
                terminate = true;
                break;
            }

#ifdef VECTORBASED
            tmp.push_back(code);
#endif

#ifdef STRINGBASED
            tmp += std::to_string(code);
#endif
        }

        if (terminate)
            break;

        // printf("Keyword added\n");
        match[tmp] = 0;
    }

    fileManager->keywordHelper->clearOriginalData();

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

#ifdef VECTORBASED
        text.push_back(code);
#endif

#ifdef STRINGBASED
        text += std::to_string(code);
#endif
    }

    fileManager->textHelper->clearOriginalData();

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