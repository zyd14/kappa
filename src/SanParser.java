import java.io.*;
import java.util.Scanner;

public class SanParser {

    private static int left;
    private static int right;
    private static int bpchain [];

    public static void addStartEnd (String line, int startColumn, int endColumn) {
        String [] elements = line.split(" ");

        int start, end;
        try {
            start = Integer.parseInt(elements[startColumn]);
            end = Integer.parseInt(elements[endColumn]);
        } catch (NumberFormatException e) {
            System.out.println("<Program SanParser> Number Format Error");
            System.out.println ("Action Aborted.");
            return;
        }

        for (int i = start; i < end; i++) {
            bpchain[i] = bpchain[i] + 1;
        }
    }

    public static void main (String [] args) {

        if (args.length != 5) {
            System.out.println ("<Program SanParser> args[0] INFILE, args[1] N, args[2] COL_START, args[3] COL_END, args[4] OUTFILE");
            System.out.println ("Action Aborted.");
            return;
        }

        Scanner scnr = new Scanner (System.in);
        String inFilePath = args[0];
        String outFilePath = args[4];
        int n, startColumn, endColumn;

        try {
            n = Integer.parseInt(args[1]);
        } catch (Exception e) {
            System.out.println ("<Program SanParser> Bad param N");
            System.out.println ("Action Aborted.");
            return;
        }

        try {
            startColumn = Integer.parseInt(args[2]);
        } catch (Exception e) {
            System.out.println ("<Program SanParser> Bad param COL_START");
            System.out.println ("Action Aborted.");
            return;
        }

        try {
            endColumn = Integer.parseInt(args[3]);
        } catch (Exception e) {
            System.out.println ("<Program SanParser> Bad param COL_END");
            System.out.println ("Action Aborted.");
            return;
        }

        if (startColumn >= endColumn) {
            System.out.println ("<Program SanParser> Bad param: COL_START >= COL_END");
            System.out.println ("Action Aborted.");
            return;
        }

        File inFile = new File (inFilePath);
        File outFile = new File (outFilePath);
        bpchain = new int [n];

        try {
            Scanner fileScnr = new Scanner (inFile);
            while (fileScnr.hasNextLine ()) {
                String line = fileScnr.nextLine();
                addStartEnd (line, startColumn, endColumn);
            }
            fileScnr.close();
        } catch (FileNotFoundException e1) {
            System.out.println("<Program SanParser> Invalid file information.");
        }

        try {
            PrintWriter writer = new PrintWriter(outFile);
            for (int i = 0; i < n; i++) {
                writer.println(i + "," + bpchain[i]);
            }
            writer.close();
        } catch (FileNotFoundException e1) {
            System.out.println("<Program SanParser> Invalid file information.");
            for (int i = 0; i < n; i++) {
                System.out.println(i + "," + bpchain[i]);
            }
        }


        scnr.close ();
    }
}
