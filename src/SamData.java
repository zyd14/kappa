public class SamData {

    private int [] seq;
    private String name;

    public SamData (int size, String name) {
        this.seq = new int [size];
        this.name = name;
    }

    public boolean increment (int low, int up) {
        if (low >= up) {
            return false;
        } else {
            for (int i = low; i < up; i++) {
                this.seq[i] = this.seq[i] + 1;
            }
            return true;
        }
    }
}
