using System.Globalization;

class WeighingMachine
{

    public WeighingMachine(int precision) => Precision = precision;

    public double TareAdjustment { get; set; } = 5;

    public int Precision { get; private set; }

    private double _weight;
    public double Weight
    {
        get => _weight;
        set
        {
            ArgumentOutOfRangeException.ThrowIfNegative(value);
            
            _weight = value;
        }

    }

    public string DisplayWeight
    {
        get
        {
            var format = new NumberFormatInfo()
            {
                NumberDecimalDigits = Precision
            };

            return $"{(Weight - TareAdjustment).ToString("f", format)} kg";
        }
    }
}
