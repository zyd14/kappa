import java.io.*;
import java.util.Scanner;

public class SamParser {

    private int left;
    private int right;
    private int [] bpchain;

    private void addStartEnd (String line) {

        int seq = line.indexOf("seq");
        int end = line.indexOf("M\t");
        if (seq < 0 || end < 0) {
            System.out.println("Inconsistent Data Formant. Line omitted.");
            return;
        }
        String clean = line.substring(seq, end);
        String [] elements = clean.split("\t");

        int start, len;
        try {
            start = Integer.parseInt(elements[1]);
            len = Integer.parseInt(elements[elements.length - 1]);
        } catch (Exception e) {
            System.out.printf("Inconsistent Data Format. Line %s omitted.\n", line);
            return;
        }

        for (int i = start; i < start + len; i++) {
            this.bpchain[i] += 1;
        }
    }

    public void parse (String [] args) {

        if (args.length != 5) {
            System.out.println ("<Program SanParser Arguments>");
            System.out.println ("\targs[0] Input file path/name.sam");
            System.out.println ("\targs[1] Length of contig [Unit: bp]");
            System.out.println ("\targs[2] Starting line number");
            System.out.println ("\targs[3] Ending line number");
            System.out.println ("\targs[4] Output file path/name.csv");
            return;
        }

        String inFilePath = args[0];
        String outFilePath = args[4];
        int n, begin, end;

        try {
            n = Integer.parseInt(args[1]);
            begin = Integer.parseInt(args[2]);
            end = Integer.parseInt(args[3]);
        } catch (Exception e) {
            System.out.println ("<Program SanParser> Bad paramenter. Action aborted.");
            return;
        }

        File inFile = new File (inFilePath);
        File outFile = new File (outFilePath);
        this.bpchain = new int [n];

        try {
            Scanner scnr = new Scanner (inFile);

            // Skip lines
            for (int i = 1; i < begin; i++) {
                scnr.nextLine();
            }

            // Read lines
            for (int i = begin; i <= end; i++) {
                String line = scnr.nextLine();
                this.addStartEnd (line);
            }
            scnr.close();
        } catch (IOException e) {
            System.out.println("<Program SanParser> I/O error. Action aborted.");
        }

        try {
            PrintWriter writer = new PrintWriter(outFile);
            for (int i = 0; i < n; i++) {
                writer.println(i + "," + this.bpchain[i]);
            }
            writer.close();
        } catch (IOException e) {
            System.out.println("<Program SanParser> I/O error. Action aborted");
            for (int i = 0; i < n; i++) {
                System.out.println(i + "," + bpchain[i]);
            }
        }
    }

    public static void main (String [] args) {
        SamParser sam = new SamParser();
        sam.parse(args);
    }
}
