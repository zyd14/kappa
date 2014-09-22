/**
 * Implement a math tree
 */
#ifndef MTREE_H
#define MTREE_H


typedef union {
	int i;
	double d;
	float f;
} Val;

typedef struct {
	char opt;
	Val left;
	Val right;
} Node;

#endif