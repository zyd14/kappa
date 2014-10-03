#include <stdio.h>
#include <string.h>

typedef struct {
	char name [256];
	char * data;
} tsv_t;

tsv_t * tsv_make (FILE * fd) {
}