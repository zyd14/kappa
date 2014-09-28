#ifndef STACK_H
#define STACK_H

#include <stdbool.h>

#define SIZE (10)
#define BOT (-1)
#define EMPTY (0)
#define ZERO (0)

#define BOOLEAN 0
#define INTEGER 1
#define FLOAT 2
#define DOUBLE 3
#define STRING 4

typedef struct {
	int t;

	union SVal {
		_Bool bval;
		int ival;
		float fval;
		double dval;
		char * strval;
	} val;
} SItem;

/**
 * Define a stack data structure. Top is the "cursor" of a digit. If there is 
 * no item in stack, cursor will be 0. If the stack is full, cursor will be
 * the number of items in stack
 */

typedef struct {
	int t;
	int cursor;
	SItem items [SIZE];
} Stack;


void stackinit (Stack * stack, int type);
void siteminit (SItem * item, int type, void value);
_Bool empty (Stack *s);
_Bool full (Stack *s);
void push (Stack *s, SItem item);
SItem * pop (Stack *s);
SItem * top (Stack *s);
void print_stack (Stack * s);

#endif