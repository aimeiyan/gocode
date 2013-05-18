#include <stdio.h>
#include <stdlib.h>

int lis(int arr[], int n) {
    int d[n];
    d[0] = 1;
    for(int i = 1; i < n; i ++) {
        int max = 0;
        for (int j = 0; j < i; j++) {
            int t = d[j];
            if (arr[j] < arr[i]) {
                t += 1;
            }
            if (t > max) {
                max = t;
            }
        }
        d[i] = max;
    }

    for (int i = 0; i < n; ++ i) {
        printf("a[%d] = %d\n", i, d[i]);
    }

    return d[n-1];
}

int lcs(char *str1, char *str2) {
    int str1n = strlen(str1);
    int str2n = strlen(str2);
    int d[str1n + 1][str2n + 1];

    // printf("------%d, %d\n", sizeof(d), sizeof(d[0]));
    for (int i = 0; i <= str1n; ++ i) {
        d[i][0] = 0;
    }

    for (int i = 0; i <= str2n; ++ i) {
        d[0][i] = 0;
    }

    for (int i = 1; i <= str1n; ++ i) {
        for (int j = 1; j <= str2n; ++ j) {
            if(str1[i - 1] == str2[j -1]) {
                d[i][j] = d[i-1][j-1] + 1;
            } else {
                d[i][j] = d[i-1][j] > d[i][j-1] ? d[i-1][j] : d[i][j-1];
            }
        }
    }
    printf("%-16c", ' ');
    for (int i = 0; i < str2n; ++ i) {
        printf("%-8c", str2[i]);
    }

    printf("\n");

    for (int i = 0; i <= str1n; ++ i) {
        if (i == 0) {
            printf("%-8c", ' ');

        } else {
            printf("%-8c", str1[i-1]);
        }

        for (int j = 0; j <= str2n; ++ j) {
            printf("%-7d ", d[i][j]);
        }
        printf("\n");
    }
    // printf("---------\n");

    return d[str1n][str2n];
}

int lcs2(char *str1, int str1n, char *str2, int str2n) {
    // printf("%s, %d, %s, %d\n", str1, str1n, str2, str2n);
    if (str1n == 0 || str2n == 0) {
        return 0;
    }

    if(str1[str1n - 1] == str2[str2n - 1]) {
        return lcs2(str1, str1n - 1, str2, str2n - 1) + 1;
    } else {
        int m1 = lcs2(str1, str1n, str2, str2n - 1);
        int m2 = lcs2(str1, str1n -1 , str2, str2n);

        if (m1 > m2) {
            return m1;
        } else {
            return m2;
        }
    }
}

int main(int argc, char** argv) {
    int arr[] = {1, 7, 3, 5, 9, 8};
    printf("%d\n", lis(arr, sizeof(arr)/ sizeof(*arr)));

    char *str1 = "abcfbc";
    char *str2 = "abfcabc";

    int n = lcs2(str1, strlen(str1), str2, strlen(str2));

    printf("%d, %d\n",n , lcs(str1, str2));

    return 0;
}
