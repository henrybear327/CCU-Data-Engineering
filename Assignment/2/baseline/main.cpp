#include <bits/stdc++.h>
#include <sys/time.h>

using namespace std;

#define tmp_files 1000

void run(char** argv)
{
    vector<string> inp;

    {
        struct timeval t1, t2;
        double elapsedTime;

        // start timer
        gettimeofday(&t1, NULL);

        printf("Reading...\n");
        FILE *fd = fopen(argv[1], "r");
        char buffer[40000];
        while (fgets(buffer, 40000, fd) != NULL) {
            inp.push_back(buffer);
        }

        // stop timer
        gettimeofday(&t2, NULL);

        // compute and print the elapsed time in millisec
        elapsedTime = (t2.tv_sec - t1.tv_sec) * 1000.0;    // sec to ms
        elapsedTime += (t2.tv_usec - t1.tv_usec) / 1000.0; // us to ms
        printf("%f ms\n", elapsedTime);

        fclose(fd);
    }

    {
        struct timeval t1, t2;
        double elapsedTime;

        // start timer
        gettimeofday(&t1, NULL);

        printf("Sorting...\n");
        sort(inp.begin(), inp.end());

        // stop timer
        gettimeofday(&t2, NULL);

        // compute and print the elapsed time in millisec
        elapsedTime = (t2.tv_sec - t1.tv_sec) * 1000.0;    // sec to ms
        elapsedTime += (t2.tv_usec - t1.tv_usec) / 1000.0; // us to ms
        printf("%f ms\n", elapsedTime);
    }

    {
        struct timeval t1, t2;
        double elapsedTime;

        // start timer
        gettimeofday(&t1, NULL);

        printf("Writing...\n");
        FILE *ans = fopen(argv[2], "w");
        for (auto i : inp) {
            fprintf(ans, "%s\n", i.c_str());
        }

        // stop timer
        gettimeofday(&t2, NULL);

        // compute and print the elapsed time in millisec
        elapsedTime = (t2.tv_sec - t1.tv_sec) * 1000.0;    // sec to ms
        elapsedTime += (t2.tv_usec - t1.tv_usec) / 1000.0; // us to ms
        printf("%f ms\n", elapsedTime);

        fclose(ans);
    }

    printf("Ends!\n");    
}

int main(int argc, char** argv)
{
    if(argc < 3) {
        printf("Please provide input and out file path\n");
        return 1;
    }

    run(argv);
    run(argv);

    return 0;
}