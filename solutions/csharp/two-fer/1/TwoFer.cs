public static class TwoFer
{
    public static string Speak() => Speak("you");

    public static string Speak(string name) => $"One for {name}, one for me.";
}
