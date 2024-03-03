package com.gig.algorithm.programmers.datastructure.stackqueue;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author : JAKE
 * @date : 2023/10/12
 */
public class CorrectParentheses {

    public static Boolean solution(String str) {
        if (str.length() <= 0) {
            return false;
        }

        String[] charArray = str.split("");

        int listSize = charArray.length;
        if (!(listSize % 2 == 0) || listSize <= 0) {
            return false;
        }

        List<String> array = new ArrayList<>();
        for (String s : charArray) {

            if (!charArray[0].equals("(")) {
                return false;
            }

            if (!charArray[charArray.length - 1].equals(")")) {
                return false;
            }

            if (array.size() == 0) {
                array.add(s);
                continue;
            }

            if (array.get(array.size() - 1).equals(s)) {
                array.add(s);
                continue;
            }

            array.remove(0);
        }

        return array.size() <= 0;
    }

    public static void main(String[] args) {
        System.out.println("answer expected : true, actual : " + solution("()()"));
        System.out.println("answer expected : true, actual : " + solution("(())()"));
        System.out.println("answer expected : false, actual : " + solution(")()("));
        System.out.println("answer expected : false, actual : " + solution("(()("));
    }
}
