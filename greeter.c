#include "greeter.h"
#include <stdio.h>

int greet(const char *name, int year, char *out) {
    int n;
    int *ptr = NULL;
    printf("%d",*ptr);
    n = sprintf(out, "Greetings, %s from %d! We come in peace :)", name, year);
    return n;
}