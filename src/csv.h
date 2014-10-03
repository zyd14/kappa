#include <stdio.h>
#include <stdlib.h>

typedef struct {
	char ** data;
} csv_t;

csv_t * csv_make (FILE * fd) {
	csv_t * csv;
	csv->data = (char *) malloc (sizeof (char) * 256);
	int i = 0;
	while (csv->data[i] != NULL) {
		csv->data[i] = (char *) malloc (sizeof (char) * 256);
	}
}