import java.util.Scanner;
 
public class Test { 
    public static void main(String args[]) {  
        String str = "", regex = "";
        int repeat = 30;

        // for(int i = 0; i < repeat * 2; i++) // exact
            // str += "a";
        
        for(int i = 0; i < repeat; i++)
            str += "a";
        
        for(int i = 0; i < repeat; i++)
            regex += "a?";
        for(int i = 0; i < repeat; i++)
            regex += "a";
        
        System.out.println(str);
        System.out.println(regex);
 
        // 簡單格式驗證 
        if(str.matches(regex)) 
            System.out.println("Match"); 
        else 
            System.out.println("Mismatch");
    } 
} 