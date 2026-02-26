public static class PhoneNumber
{
    private const string NewYorkAreaCode = "212";
    
    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber) => (phoneNumber[0..3] == "212", phoneNumber[4..7] == "555", phoneNumber[8..12]);

    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) phoneNumberInfo) => phoneNumberInfo.IsFake;
}
