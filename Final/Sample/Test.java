import java.util.Scanner;
 
public class Test { 
    public static void main(String args[]) {  
        String str = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"; 
 
        // 簡單格式驗證 
        if(str.matches("a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?a?aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")) 
            System.out.println("Match"); 
        else 
            System.out.println("Mismatch");
    } 
} 