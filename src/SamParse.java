import java.io.*;
import java.util.*;

public class SamParse {
    private LinkedList<String> seqnames = new LinkedList<String>();
    private LinkedList<Integer> seqsizes = new LinkedList<Integer>();
    private LinkedList<SamData> refseqs = new LinkedList<SamData> ();

    public SamParse () {
    }

    private static boolean parseArgs (String [] args) {
        return false;
    }

    public static void main (String [] args) {
        boolean goodArgs = parseArgs (args);
        if (!goodArgs) {
            return;
        }
    }
}
