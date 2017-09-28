#include "SearchEngine.h"
#include "FileManager.h"
#include "UTF8Helper.h"

#include <cstdio>
#include <cstdlib>
#include <cstring>
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
                printf("Matched at %d: ", i);
                {
                    char location[111];
                    sprintf(location, "\nMatched at %d: ", i);
                    int ret =
                        write(fileManager->resultHelper->fd, location, strlen(location));

                    if (ret == -1) {
                        perror("write() error");
                        exit(1);
                    }
                }

                for (auto k : candidate) {
                    printf("%d\n", fileManager->keywordHelper->dictionary[k].bytes);
                    {
                        int ret = write(fileManager->resultHelper->fd,
                                        fileManager->keywordHelper->dictionary[k].buffer,
                                        fileManager->keywordHelper->dictionary[k].bytes);

                        if (ret == -1) {
                            perror("write() error");
                            exit(1);
                        }
                    }
                }

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

        printf("Keyword added\n");
        match[tmp] = 0;
    }

    printf("Done\n\n");
}

void SearchEngine::loadText()
{
    printf("Loading text...\n");

    for (int code = fileManager->textHelper->extractWord(); code != 0;
         code = fileManager->textHelper->extractWord()) {
        text.push_back(code);
    }

    printf("Done\n\n");
}