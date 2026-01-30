using System.Text;

public static class Identifier
{
    private static bool IsGreekLowerCase(char c) => (c >= 'α' && c <= 'ω');

    public static string Clean(string identifier)
    {
        var sb = new StringBuilder();

        var isAfterDash = false;

        foreach (var c in identifier)
        {
            sb.Append(c switch
            {
                _ when IsGreekLowerCase(c) => default,
                _ when isAfterDash => char.ToUpperInvariant(c),
                _ when char.IsLetter(c) => c,
                _ when char.IsWhiteSpace(c) => '_',
                _ when char.IsControl(c) => "CTRL",
                _ => default
            });
            
            isAfterDash = c.Equals('-');
        }

        return sb.ToString();
    }
}
