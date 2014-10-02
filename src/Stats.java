public class Stats {

	public static double max (int [] num) {
		if (num == null || num.length < 1) {
			return 0.0;
		}
		double max = num[0];
		for (int i = 0; i < num.length; i++) {
			if (num[i] > max)
				max = num [i];
		}
		return max;
	}

	public static double max (double [] num) {
		if (num == null || num.length < 1) {
			return 0.0;
		}
		double max = num[0];
		for (int i = 0; i < num.length; i++) {
			if (num[i] > max)
				max = num [i];
		}
		return max;
	}

	public static double min (int [] num) {
		if (num == null || num.length < 1) {
			return 0.0;
		}
		double min = num[0];
		for (int i = 0; i < num.length; i++) {
			if (num[i] < min)
				min = num [i];
		}
		return min;
	}

	public static double min (double [] num) {
		if (num == null || num.length < 1) {
			return 0.0;
		}
		double min = num[0];
		for (int i = 0; i < num.length; i++) {
			if (num[i] < min)
				min = num [i];
		}
		return min;
	}

	public static double avg (int [] num) {
		if (num == null || num.length < 1) {
			return 0.0;
		}
		double sum = 0.0;
		for (int i = 0; i < num.length; i++) {
			sum += num[i];
		}
		return sum/num.length;
	}

	public static double avg (double [] num) {
		if (num == null || num.length < 1) {
			return 0.0;
		}
		double sum = 0.0;
		for (int i = 0; i < num.length; i++) {
			sum += num[i];
		}
		return sum/num.length;
	}
}