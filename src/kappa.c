#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "param.h"
#include "parser.h"

int main (void) {

	_Bool fin = false;
	int unit = UNITIN;

	char * input = (char *) malloc (unit);
	char * ref = (char *) malloc (unit);

	puts (INFO_L1);
	puts (INFO_L2);
	puts (INFO_L3);

	while ( strcmp ("exit", ref) != 0) {
		printf (">>> ");
		getline (&input, &unit, stdin);
		ref = rmnl (input);

		if (strcmp ("exit", ref) != 0) {
			parse (ref);
		}
		else {
			puts ("Bye!");
		}
	}
	
	free (input);

	exit (0);
}