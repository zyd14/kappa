/**
 * Implement Stack Using C
 * @author Yubing Hou
 * @date August 7, 2014
 */
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "stack.h"


int main (void) {
	Stack stack;
	stackinit (&stack, STRING);
	SItem ()
	printf ("Is stack full? %d\n", full(&stack));
	printf ("Is stack empty? %d\n", empty (&stack));
	print_stack(&stack);
	push(&stack)
	exit (0);
}


/**
 * Initialize a stack
 */
void stackinit (Stack * stack, int type) {
	stack->cursor = EMPTY;
	stack->t = type;
	int i;
	for (i = EMPTY; i < SIZE; i++) {
		SItem item;
		item.t = type;
		switch (type) {
			case (BOOLEAN): {
				item.t = BOOLEAN;
				item.val.bval = false;
			}
			case (INTEGER): {
				item.t = INTEGER;
				item.val.ival = 0;
			}
			case (FLOAT): {
				item.t = FLOAT;
				item.val.fval = 0.0f;
			}
			case (DOUBLE): {
				item.t = DOUBLE;
				item.val.dval = 0.0;
			}
			case (STRING): {
				item.t = STRING;
				item.val.strval = "";
			}
		}
		stack->items[i] = item;;
	}
}

/**
 * Tell whether stack is empty
 */
_Bool empty (Stack * s) {
	return s->cursor == EMPTY;
}

/**
 * Tell whether stack is full
 */
_Bool full (Stack * s) {
	return s->cursor == SIZE;
}

/**
 * Push an item to stack. Executation will halt if stack overflows
 */
void push (Stack *s, SItem *item) {
	if (full(s) == false && (item->t == s->t)) {
		s->items[s->cursor] = item;
		s->cursor += 1;
	} else if (item->t != s->t) {
		puts ("<Error: Type Mismatch>");
	}
	else {
		puts("<Error: Stack Overflow>");
		exit(1);
	}
}

/**
 * Pop out an item from stack. Executation will halt if stack underflows
 *
SItem * pop (Stack *s) {
	if (isempty(s) == true) {
		puts("<Error: Stack Underflow>");
		return NULL;
	}
	else {
		int tmp = (*s).items[(*s).cursor];
		(*s).cursor -= 1;
		return tmp;
	}
}*/

/**
 * Get the value of the top element of this stack
 *
SItem * top (Stack *s) {
	if (isempty(s) == true) {
		puts("<Empty Stack: No Top Element>");
		return -1;
	}
	else {
		int tmp = s->items[s->cursor];
		printf("\t(stack top: %d)\n", tmp);
		return tmp;
	}
}
*/
/**
 * Display information 
 *
void toString(Stack *s) {
	puts("Type: struct Stack");
	printf("Top: %d\n", s->cursor);
	printf("Full Status: %d\n", isfull(s));
	printf("Empty Status: %d\n", isempty(s));
}*/

/**
 * Print out stack
 */
void print_stack (Stack *s) {
	int i = SIZE;
	for (i = SIZE - 1; i >= EMPTY; i--) {
		switch (s->t) {
			case (BOOLEAN): printf("\t%d |-%d-|\n", i, s->items[i].val.bval);
			case (INTEGER): printf("\t%d |-%d-|\n", i, s->items[i].val.ival);
			case (FLOAT):   printf("\t%d |-%f-|\n", i, s->items[i].val.fval);
			case (DOUBLE):  printf("\t%d |-%lf-|\n", i, s->items[i].val.dval);
			case (STRING):  printf("\t%d |-%s-|\n", i, s->items[i].val.strval);
		}
	}
}