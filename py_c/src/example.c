/* example.c */
#include "example.h"
/* Compute GCD of two positive integers x and y */
int gcd(int x, int y) {
    int g;
    g = y;
    while (x > 0) {
        g = x;
        x = y % x;
        y = g;
    }
    return g;
}
/* Replace a character in a string */
int replace(char *s, char oldch, char newch) {
    int nrep = 0;
    while (s = strchr(s,oldch)) {
        *(s++) = newch;
        nrep++;
    }
    return nrep;
}
/* Distance between two points */
double distance(Point *a, Point *b) {
    double dx,dy;
    dx = a->x - b->x;
    dy = a->y - b->y;
    return sqrt(dx*dx + dy*dy);
}
