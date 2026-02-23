public struct Coord
{
    public Coord(ushort x, ushort y)
    {
        X = x;
        Y = y;
    }

    public ushort X { get; }
    public ushort Y { get; }
}

public struct Plot
{
    public Coord Coord1 { get; }
    public Coord Coord2 { get; }
    public Coord Coord3 { get; }
    public Coord Coord4 { get; }
    public double LongestSide { get; }

    public Plot(Coord coord1, Coord coord2, Coord coord3, Coord coord4)
    {
        Coord1 = coord1;
        Coord2 = coord2;
        Coord3 = coord3;
        Coord4 = coord4;

        var sideLengths = new List<double>()
        {
            GetDistance(coord1, coord2),
            GetDistance(coord2, coord3),
            GetDistance(coord3, coord4),
            GetDistance(coord4, coord1)
        };

        LongestSide = sideLengths.Max();
    }

    private double GetDistance(Coord a, Coord b) => Math.Sqrt(Math.Pow(b.X - a.X, 2) + Math.Pow(b.Y - a.Y, 2));
}


public class ClaimsHandler
{
    private List<Plot> _plots = new List<Plot>();

    public void StakeClaim(Plot plot)
    {
        if (!IsClaimStaked(plot))
        {
            _plots.Add(plot);
        }
    }

    public bool IsClaimStaked(Plot plot) => _plots.Contains(plot);

    public bool IsLastClaim(Plot plot) => _plots.Last().Equals(plot);

    public Plot GetClaimWithLongestSide() => _plots.OrderByDescending(x => x.LongestSide).First();
}
