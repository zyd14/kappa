#include <iostream>
#include <string>

using namespace std;

int strmat (string a, string b) {
	int len_a = a.size ();
	int len_b = b.size ();

	int i, j, p, match;
	for (i = 0; i < len_a; i++) {
		bool match = true;
		bool count = 0;
		for (j = 0; (j < len_a - i) && (j < len_b); j++) {
			match = (a[i + j] == b[j]);
			
			if (match == true) {
				count++;
			}
			else break;
		}
		if (match == true) {
			p = i;
			return p;
		} else continue;
	}
	return -1;
}

void printm (string a, string b) {
	int match = strmat(a, b);
	cout << a << "\n";
	int i = 0;
	for (i = 0; i < match; i++) {
		cout << " ";
	}
	cout << b << "\n";
}

int main (void) {
	string a = "ATCGGTCA";
	string b =  "TCGGTCAT";
	string c =   "CGGTCATC";
	string d =    "GGTCATCG";

	cout << "string a = " << a << "\n";
	cout << "string b = " << b << "\n";
	cout << "string c = " << c << "\n";
	cout << "string d = " << d << "\n";

	string contigs [] = {a, b, c, d};

	int i = 0, j = 0;
	for (i = 0; i < 4; i++) {
		for (j = 0; j < 4; j++) {
			//cout << contigs[i] << " matches " << contigs[j] << " from " <<
			 //strmat (contigs[j], contigs[i]) << "\n";
			printm (contigs [j], contigs [i]);
		}
	}
/*
	cout << "b matches a from " << strmat (a, b) << "\n";
	cout << "c matches a from " << strmat (a, c) << "\n";
	cout << "d matches a from " << strmat (a, d) << "\n";
	cout << "a matches b from " << strmat (b, a) << "\n";
	cout << "a matches c from " << strmat (c, a) << "\n";
	cout << "a matches d from " << strmat (d, a) << "\n";
*/
	return 0;
}