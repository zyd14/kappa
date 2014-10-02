/**
 * DNA sequence file parser: .same file parser for sorted sam file produced by 
 * Bowtie.
 * Author: Yubing Hou
 * Year: 2014
 * Course: COMP 488-Metagenomics
 * School: Loyola University Chicago
 */
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "statistics.h"

#define MAXSEQNUM 128
#define LINELENGTH 2048

typedef struct {
	char * seqname;
	int * sequence;
	int seqlen;
} Refseq;

int count;
Refseq * seqs;
FILE *infd;
FILE *outfd;
//============================== Function Signatures =========================//
double getavg (int low, int up);
int    getmax (int low, int up);
int    getmin (int low, int up);
void   parse_args (int argc, char * argv []);
void   parse_entry ();
_Bool  parse_file (char * argv []);
void   parse_refs ();
void   parse_terminate ();
void   refInfo (Refseq * r);
//============================================================================//


double getavg (int low, int up) {
	int * seq = subarray (seqs->sequence, low, up);
	double avg = average (seq);
	free (seq);
	return avg;
}

int getmax (int low, int up) {
	int * seq = subarray (seqs->sequence, low, up);
	int val = max (seq);
	free (seq);
	return val;
}

int getmin (int low, int up) {
	int * seq = subarray (seqs->sequence, low, up);
	int val = min (seq);
	free (seq);
	return val;
}

void parse_args (int argc, char * argv []) {
	if (argc != 3) {
		puts ("Must give 2 valid arguments in following order:");
		puts ("./sam [input file] [output file]");
		exit (1);
	}
}

// FINISH implementing this method!
void parse_entry () {
	char * entry = (char *) malloc (sizeof (char) * LINELENGTH);
	int line = 0;
	int c;

	char refid [80];
	while ((c = fgetc (infd)) != EOF) {
		if (c == '\n') line = line + 1;
	}
	printf ("LINE = %d\n", line);
	//int start = parseInt ()
	/*int i = 0;
	
	int status = fgets (entry, LINELENGTH, infd);

	char * RNAME = (char *) malloc (sizeof (char) * LINELENGTH);
	int POS;

	for (i = 0; i < count; i++) {
		while (status != NULL) {
			status = ;
		}
	}*/
	free (entry);
}

_Bool parse_file (char * argv []) {
	char *infile = argv[1];
	char *outfile = argv[2];
	infd  = fopen (infile, "r");
	outfd = fopen (outfile, "w");

	if (infd == NULL || outfd == NULL ) {
		if (infd == NULL)
			printf ("Unable to access file \'%s\'\n", infile);
		if (outfd == NULL)
			printf  ("Unable to access file \'%s\'\n", outfile);
		puts ("Action aborted");
		exit (1);
	}
	return true;
}

// TODO 01: Console currently print weird characters
void parse_refs () {
	seqs = (Refseq *) malloc (128 * sizeof (Refseq));
	char * line = (char *) malloc (sizeof (char) * LINELENGTH);
	char * lncp = (char *) malloc (sizeof (char) * LINELENGTH);
	
	_Bool isref = true;


	int i = 0;
	while (isref) {
		fgets (line, LINELENGTH, infd);
		lncp = spaceClean (line);
		if (containStr (lncp,"@SQ") == true) {
			int start = indexOfStr (lncp, "SN:");
			int end = indexOfStr (lncp, "LN:");
			char * name = substring (lncp, start+3, end);
			int len = parseInt (substring (lncp, end, strlen (lncp)));

			(seqs[i]).seqname = (char *) malloc (sizeof (char) * MAXSEQNUM);
			(seqs[i]).seqname = name;
			(seqs[i]).seqlen = len;
			(seqs[i]).sequence = (int *) malloc (sizeof (int) * len);
			refInfo (&seqs[i]);
			i = i + 1;
		}
		if (charAt(lncp, 0) != '@') isref = false;
	}

	free (line);
	free (lncp);
}

void parse_terminate () {
	free (seqs);
	fclose (infd);
	fclose (outfd);
}

void refInfo (Refseq * r) {
	printf ("\nref=<%s>\n", r->seqname);
	printf ("len=<%d>\n", r->seqlen);
}

/**
 * Command line arguments: <INPUT FILE>, <OUTPUT FILE>
 */
int main (int argc, char * argv []) {

	parse_args (argc, argv);
	if (parse_file (argv) == true) {
		parse_refs ();
		parse_entry ();
		parse_terminate ();
	}
	exit (0);
}