#ifndef PARSER_H
#define PARSER_H

#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

#include "param.h"
#include "token.h"

_Bool isint (char * str);
_Bool isfloat (char * str);
_Bool mathexp (char * str);
int getint (char * str);
char * lower (char * str);
char * upper (char * str);
char * rmnl (char * str);
int match (char * str);
void parse (char * str);

/**
 * Tell if string can be identified as int
 */
_Bool isint (char * str) {
	int len = strlen (str);
	int i;
	for (i = 0; i < len; i++) {
		if (str[i] > '9' || str[i] < '0') {
			return false;
		}
	}
	return true;
}

/**
 * Convert string to int
 */
int getint (char * str) {
	if (!isint (str)) return -1;
	else {
		char * toint = (char *) malloc (UNITIN);
		int i = 0;
		for (i = 0; str[i] != '\0'; i++) {
			toint[i] = str[i];
		}
		int get = atoi (toint);
		free (toint);
		return get;
	}
}

_Bool mathexp (char * str) {
	char MATH [] = {ADD, SUB, MUL, DIV, MOD, LPA, RPA, FAC, POW, DEC};
	int len = strlen (str);
	int i = 0;
	_Bool isexp = true;
	_Bool isopt = false;
	for (i = 0; i < len; i++) {
		if ( str[i] < '0' || str[i] > '9') {
			int j = 0;
			for (j = 0; j < SIZE_H; j++) {
				if (str[i] == MATH[j]) {
					isopt = true;
					break;
				}
			}
			if (isopt == false)
				return false;
		}
	}
	return isexp && isopt;
}

char * rmnl (char * str) {
	int len = strlen (str);
	char * nstr = (char *) malloc (len);
	int i = 0;
	for (i = 0; str[i] != '\n'; i++) {
		nstr[i] = str[i];
	}
	return nstr;
}

char * lower (char * str) {
	char * nstr = (char *) malloc (strlen (str));
	int i;
	for (i = 0; str[i] != EOS; i++) {
		nstr[i] = str[i];
		if (nstr[i] >= 'A' && nstr[i] <= 'Z') {
			nstr[i] = nstr[i] + 32 ;
		}
	}
	return nstr;
}

char * upper (char * str) {
	char * nstr = (char *) malloc (strlen (str));
	int i;
	for (i = 0; str[i] != EOS; i++) {
		nstr[i] = str[i];
		if (nstr[i] >= 'a' && nstr[i] <= 'z') {
			nstr[i] = nstr[i] - 32 ;
		}
	}
	return nstr;
}

void parse (char * str) {
	printf ("upper (): %s\n", upper (str));
	printf ("lower (): %s\n", lower (str));
	printf ("isint (): %d\n", isint (str));
	printf ("mathexp (): %d\n", mathexp (str));
	if (isint (str)) {
		printf ("int val = %d\n", getint(str));
	}
	if (mathexp (str)) {
		printf ("\'%s\' is a valid mathematical expression.\n", str);
	}
}

int match (char * str) {

}

#endif