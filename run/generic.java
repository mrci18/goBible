import java.util.*;

public class Charles {
    public static void main (String[] args) {
        Integer[] iray = {1,2,3,4};
        Character[] cray = {'c', 'h', 'a'};

        printMe(iray);
        printMe(cray);
    }
//Overload methods
    // public static void printMe(Integer[] i) {
    //     for (Integer x : i)
    //         System.out.printf("%s", x);

    //     System.out.println();
    // }

    // public static void printMe(Character[] i) {
    //     for (Character x : i)
    //         System.out.printf("%s", x);

    //     System.out.println();
    // }

//Generic method 
    public static <T> void printMe(T[] x) {
        for (T b : x)
            System.out.printf("%s", b);

        System.out.println();
    }

}