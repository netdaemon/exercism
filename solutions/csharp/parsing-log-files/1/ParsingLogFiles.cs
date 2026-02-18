using System.Text.RegularExpressions;

public class LogParser
{
    private const string IsValidLinePattern = @"^\[(TRC|DBG|INF|WRN|ERR|FTL)\]";
    private const string LogLinePattern = @"<.*?>";
    private const string QuotedPasswordsPattern = @""".*password.*""";
    private const string EndOfLinePattern = @"end-of-line\d*";
    private const string PasswordsPattern = @"password\w+";

    public bool IsValidLine(string text) => new Regex(IsValidLinePattern).IsMatch(text);

    public string[] SplitLogLine(string text) => new Regex(LogLinePattern).Split(text);

    public int CountQuotedPasswords(string lines) => new Regex(QuotedPasswordsPattern, RegexOptions.IgnoreCase).Matches(lines).Count;

    public string RemoveEndOfLineText(string line) => new Regex(EndOfLinePattern).Replace(line, "");

    public string[] ListLinesWithPasswords(string[] lines)
    {
        var newLines = new List<string>();

        foreach (var line in lines)
        {
            var match = new Regex(PasswordsPattern, RegexOptions.IgnoreCase).Match(line);

            if (match.Success)
            {
                newLines.Add($"{match.Value}: {line}" );
            }
            else {
                newLines.Add($"--------: {line}");
            }
        }

        return [.. newLines];
    } 

}
