#include <stdlib.h>
#include <string.h>

double average (int * a);
char   charAt (char * str, int index);
_Bool  containStr (char * str, char * tar);
int    indexAfter (char * str, char c, int rep);
int    indexOf (char * str, char c);
int    indexOfStr (char * str, char * cmp);
int    max (int * a);
int    min (int * a);
int    parseInt (char * str);
double parseDouble (char * str);
char * spaceClean (char * str);
int *  subarray (int * a, int low, int high);
char * substring (char * str, int start, int end);


double average (int * a) {
	int sum = 0;
	int i = 0;
	while (&a[i] != NULL) {
		sum = sum + a[i];
		i = i + 1;
	}
	return sum / i;
}

char charAt (char * str, int index) {
	if (index > strlen (str)) return '\0';
	return str[index];
}

_Bool containStr (char * str, char * tar) {
	int slen = strlen (str), tlen = strlen (tar);
	int i = 0, j = 0;
	for (i = 0; i < slen; i++) {
		_Bool has = true;
		for (j = 0; j < tlen; j++) {
			if (str[i + j] != tar[j]) {
				has = false;
				break;
			}
		}
		if (has == true) return true;
	}
	return false;
}

int indexAfter (char * str, char c, int rep) {
	int i = 0;
	int pos = 0;
	while (i < rep) {
		pos = pos + indexOf (str, c);
		str = substring (str, pos, strlen (str));
		i = i + 1;
	}
	return pos;
}

int indexOf (char * str, char c) {
	int i = 0, len = strlen(str);
	for (i = 0; i < len; i++) {
		if (str[i] == c) {
			return i;
		}
	}
	return -1;
}

int indexOfStr (char * str, char * cmp) {
	int slen = strlen (str);
	int clen = strlen (cmp);

	int i = 0, j = 0, pos;
	for (i = 0; i < slen; i++) {
		_Bool match = true;
		for (j = 0; j < clen; j++) {
			if (str[i + j] != cmp[j]) {
				match = false;
				break;
			}
		}
		if (match == true) return i;
	}
	return -1;
}

int max (int * a) {
	int max = a[0];
	int i = 0;
	while (&a[i] != NULL) {
		if (a[i] > max) {
			max = a[i];
		}
		i = i + 1;
	}
	return max;
}

int min (int * a) {
	int min = a[0];
	int i = 0;
	while (&a[i] != NULL) {
		if (a[i] < min) {
			min = a[i];
		}
		i = i + 1;
	}
	return min;
}

char * spaceClean (char * str) {
	int len = strlen (str);
	char * cp = (char *) malloc (sizeof (char) * len);
	int i = 0, j = 0;
	for (i = 0; i < len; i++) {
		if (str[i] != '\t' && str[i] !=' ') {
			cp[j] = str[i];
			j = j + 1;
		} 
	}
	return cp;
}

int * subarray (int * a, int low, int high) {
	if (high <= low) return NULL;
	int * seq = (int *) malloc (sizeof (int) * (high - low));
	int i = 0;
	for (i = 0; i < high - low; i++) {
		seq[i] = a[i + low];
	}
	return seq;
}

char * substring (char * str, int start, int end) {
	if (end < start) return NULL;
	char * cp = (char *) malloc (sizeof (int) * (end - start));
	int i = 0;
	for (i = 0; i < end - start; i++) {
		cp[i] = str[start + i];
	}
	return cp;
}

int parseInt (char * str) {
	int start, end, len, i = 0;
	while (str[i] > '9' || str [i] < '0') {
		i++;
	}
	start = i;
	while (str[i] <= '9' && str[i] >= '0') {
		i++;
	}
	end = i;
	len = end - start;
	char * cp = (char *) malloc (sizeof (char) * len);
	for (i = 0; i < len; i++) {
		cp[i] = str[start + i];
	}
	int val = atoi (cp);
	free (cp);
	return val;
}

double parseDouble (char * str) {
	int start, end, len, count = 0, i = 0;
	while (str[i] > '9' || str [i] < '0') {
		i++;
	}
	start = i;
	while ((str[i] <= '9' && str[i] >= '0') || str[i] == '.') {
		i++;
		if (count > 1) break;
		if (str[i] == '.') {
			count = count + 1;
		}
	}
	end = i;
	len = end - start;
	char * cp = (char *) malloc (sizeof (char) * len);
	for (i = 0; i < len; i++) {
		cp[i] = str[start + i];
	}
	double val = atof (cp);
	free (cp);
	return val;
}