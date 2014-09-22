all: bin/kappa

bin/kappa: src/kappa.c
	gcc src/kappa.c -o bin/kappa