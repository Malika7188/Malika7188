# Documentation

For the instance of (cap) which in this case is word[i], the program checks the previous word before (cap); word[i-1], then capitalize the first index of word[i-1] i.e (“this is a house (cap)”) and the outcome is (“this is a House”) then trim (cap) by joining words before and after cap

For (hex), the program checks the the word before (hex) and replace it with with the decimal version of it (i.e “i am 42(hex) years”) the expected outcome would be (“i am 62 years”). the (hex) will therefore be removed by appending the word before and after (hex) to form the updated result

For the instance of apostrophe, the program uses the odd and even format to move the apostrophes. i.e "this is an (’ elephant ’ ) the firs apostrophe is treated as an even number then its move forward at word[i+1]. The second apostrophe will be treated as an odd number then moved back at word[i-1].

For the instance of (cap, ) i.e (cap, 3)  the program checks the first 3 previous word before (cap); word[i-1], then capitalize the first index of every 3 previous words before word[i] i.e ("this is a house (cap, 3)") and the outcome is ("this Is A House") then trim (cap, 3) by joining words before and after cap to form the updated words

For the instance of (low, ) i.e (low, 3) the program checks the first 3 previous word before (low); word[i], then converts the first index of every 3 previous words before word[i] to lower case. i.e (“this IS A HOUSE (low, 3)”) and the outcome is (“this is a house”) then trim (cap, 3) by joining words before and after cap to form the updated words

For vowels; the program checks if the word[i] is a vowel, if the next word[i+1][0] starts with a vowel i.e "This is a apostrophe" if it finds a vowels at word[i+1][0] it converts word[i] to an i.e the result is: "This is an apostrophe

In the main package, the program calls the argument from the command line and checks the len(os.Args) of arguments to be 3. If the len(os.Args) is less than or more than 3 then the program exits… the program therefore read the content of argument and return file containing the contents and the errors encountered… being that its my main function, all the functions are called here