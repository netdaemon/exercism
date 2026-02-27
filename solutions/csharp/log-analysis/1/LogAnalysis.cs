public static class LogAnalysis 
{
    public static string SubstringAfter(this string str, string delimeter) => str.Split(delimeter)[1];

    public static string SubstringBetween(this string str, string delimeter1, string delimeter2) => str.Split(delimeter1).Last().Split(delimeter2).First();
    
    public static string Message(this string str) => str.SubstringAfter(": ");

    public static string LogLevel(this string str) => str.SubstringBetween("[", "]");
}