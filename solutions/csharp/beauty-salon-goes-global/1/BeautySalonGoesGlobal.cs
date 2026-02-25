using System.Globalization;
using System.Runtime.InteropServices;

public enum Location
{
    NewYork,
    London,
    Paris
}

public enum AlertLevel
{
    Early,
    Standard,
    Late
}

public static class Appointment
{
    private static bool _isWindows = RuntimeInformation.IsOSPlatform(OSPlatform.Windows);

    public static DateTime ShowLocalTime(DateTime dtUtc) => TimeZoneInfo.ConvertTimeFromUtc(dtUtc, TimeZoneInfo.Local);

    public static DateTime Schedule(string appointmentDateDescription, Location location)
    {
        var dt = DateTime.Parse(appointmentDateDescription);
        var tzInfo = GetTimeZoneInfo(location);
        return TimeZoneInfo.ConvertTimeToUtc(dt, tzInfo);
    }

    public static DateTime GetAlertTime(DateTime appointment, AlertLevel alertLevel) => alertLevel switch
    {
        AlertLevel.Early => appointment.AddDays(-1),
        AlertLevel.Standard => appointment.AddMinutes(-105),
        AlertLevel.Late => appointment.AddMinutes(-30),
        _ => throw new ArgumentOutOfRangeException()
    };

    public static bool HasDaylightSavingChanged(DateTime dt, Location location)
    {
        var tzInfo = GetTimeZoneInfo(location);
        return tzInfo.IsDaylightSavingTime(dt.AddDays(-7)) || tzInfo.IsDaylightSavingTime(dt.AddDays(7));
    }

    public static DateTime NormalizeDateTime(string dtStr, Location location)
    {
        var cultureInfo = GetCultureInfo(location);

        DateTime.TryParse(dtStr, cultureInfo, DateTimeStyles.None, out var dt);

        return dt;
    }

    public static string GetTimeZoneString(Location location) => location switch
    {
        Location.NewYork => _isWindows ? "Eastern Standard Time" : "America/New_York",
        Location.London => _isWindows ? "GMT Standard Time" : "Europe/London",
        Location.Paris => _isWindows ? "W. Europe Standard Time" : "Europe/Paris",
        _ => throw new ArgumentOutOfRangeException()
    };

    public static TimeZoneInfo GetTimeZoneInfo(Location location)
    {
        var tz = GetTimeZoneString(location);
        return TimeZoneInfo.FindSystemTimeZoneById(tz);
    }

    public static CultureInfo GetCultureInfo(Location location)
    {
        var culture = location switch
        {
            Location.NewYork => "en-US",
            Location.London => "en-GB",
            Location.Paris => "fr-FR",
            _ => throw new ArgumentOutOfRangeException()
        };

        return CultureInfo.GetCultureInfo(culture);
    }
}
