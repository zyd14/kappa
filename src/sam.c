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

void parse_entry () {
	int i = 0;
	char * entry = (char *) malloc (sizeof (char) * LINELENGTH);
	int status = fgets (entry, LINELENGTH, infd);

	char * RNAME = (char *) malloc (sizeof (char) * LINELENGTH);
	int POS;

	for (i = 0; i < count; i++) {
		while (status != NULL) {
			status = fgets (entry, LINELENGTH, infd);
		}
	}
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

// TODO: Finish implementing this function
void parse_refs () {
	seqs = (Refseq *) malloc (128 * sizeof (Refseq));
	char * line = (char *) malloc (sizeof (char) * LINELENGTH);

	fgets (line, LINELENGTH, infd);
	char * lncp = tabClean (line);


	int i = 0;
	while (charAt (line, 0) == '@') {
		printf ("CLEANLINE=%s\n", lncp);
		if (strcmp (substring (line, 0,3), "@SQ") == 0) {
			int start = indexOfStr (line, "SN:");
			int end = indexOfStr (line, "LN:");
			char * name = substring (line, start+3, end);
			int len = parseInt (substring (line, end, strlen (line)));

			(seqs[i]).seqname = (char *) malloc (sizeof (char) * MAXSEQNUM);
			(seqs[i]).seqname = name;
			(seqs[i]).seqlen = len;
			(seqs[i]).sequence = (int *) malloc (sizeof (int) * len);
			i = i + 1;
			printf ("SEQ: %s LEN: %d\n", name, len);
		}
		fgets (line, LINELENGTH, infd);
		lncp = tabClean (line);
		printf ("CLEANLINE=%s\n", lncp);
	}

	free (line);
}

void parse_terminate () {
	free (seqs);
	fclose (infd);
	fclose (outfd);
}

/**
 * Command line arguments: <INPUT FILE>, <CONTIG LENGTH>, <STARTING LINE>,
 * <ENDING LINE>, <OUTPUT FILE>
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