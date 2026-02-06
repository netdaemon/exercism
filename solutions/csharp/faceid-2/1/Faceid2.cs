public class FacialFeatures
{
    public string EyeColor { get; }
    public decimal PhiltrumWidth { get; }

    public FacialFeatures(string eyeColor, decimal philtrumWidth)
    {
        EyeColor = eyeColor;
        PhiltrumWidth = philtrumWidth;
    }

    public override bool Equals(object? obj)
    {
        var facialFeatures = obj as FacialFeatures;

        return facialFeatures?.EyeColor ==  EyeColor && facialFeatures?.PhiltrumWidth == PhiltrumWidth;
    }

    public override int GetHashCode() => HashCode.Combine(EyeColor, PhiltrumWidth);
}

public class Identity
{
    public string Email { get; }
    public FacialFeatures FacialFeatures { get; }

    public Identity(string email, FacialFeatures facialFeatures)
    {
        Email = email;
        FacialFeatures = facialFeatures;
    }

    public override bool Equals(object? obj)
    {
        var identity = obj as Identity;

        return identity?.Email == Email && identity.FacialFeatures.Equals(FacialFeatures);
    }

    public override int GetHashCode() => HashCode.Combine(Email, FacialFeatures);
}

public class Authenticator
{
    private static readonly Identity _admin = new Identity("admin@exerc.ism", new FacialFeatures("green", 0.9m));
    private HashSet<Identity> _registeredIdentities = new();

    public static bool AreSameFace(FacialFeatures faceA, FacialFeatures faceB) => faceA.Equals(faceB);

    public bool IsAdmin(Identity identity) => identity.Equals(_admin);

    public bool Register(Identity identity)
    {
        if (IsRegistered(identity))
        {
            return false;
        }

        _registeredIdentities.Add(identity);

        return true;
    }

    public bool IsRegistered(Identity identity) => _registeredIdentities.Contains(identity);

    public static bool AreSameObject(Identity identityA, Identity identityB) => ReferenceEquals(identityA, identityB);
}
