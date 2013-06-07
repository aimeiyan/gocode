/* file : example.h */
#include <stdio.h>
#include <string.h>
#include <math.h>
typedef struct Point {
    double x;
    double y;
} Point;
/* Compute the GCD of two integers x and y */
extern int
gcd(int x, int y);
/* Replace och with nch in s and return the number of replacements */
extern int
replace(char *s, char och, char nch);
/* Compute the distance between two points */
extern double distance(Point *a, Point *b);
/* A preprocessor constant */
#define MAGIC 0x31337
