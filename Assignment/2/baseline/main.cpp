#include <bits/stdc++.h>
#include <sys/time.h>

using namespace std;

#define tmp_files 1000

void run()
{
    vector<string> inp;

    {
        struct timeval t1, t2;
        double elapsedTime;

        // start timer
        gettimeofday(&t1, NULL);

        printf("Reading...\n");
        FILE *fd = fopen("../testcase/number_500M.in", "r");
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
        FILE *ans = fopen("../testcase/number_500M.baseline.out", "w");
        for (auto i : inp) {
            fprintf(ans, "%s\n", i.c_str());
        }

        printf("Ends!\n");

        // stop timer
        gettimeofday(&t2, NULL);

        // compute and print the elapsed time in millisec
        elapsedTime = (t2.tv_sec - t1.tv_sec) * 1000.0;    // sec to ms
        elapsedTime += (t2.tv_usec - t1.tv_usec) / 1000.0; // us to ms
        printf("%f ms\n", elapsedTime);

        fclose(ans);
    }
}

int main()
{
    run();
    run();

    return 0;
}