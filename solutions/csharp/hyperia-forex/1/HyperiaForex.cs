public struct CurrencyAmount
{
    private decimal amount;
    private string currency;

    public CurrencyAmount(decimal amount, string currency)
    {
        this.amount = amount;
        this.currency = currency;
    }

    public static bool operator ==(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException();
        }

        return left.currency == right.currency;
    }

    public static bool operator !=(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException();
        }

        return left.amount != right.amount || left.currency != right.currency;
    }


    public static bool operator <(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException();
        }

        return left.amount < right.amount;
    }

    public static bool operator >(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException();
        }

        return left.amount > right.amount;
    }

    // TODO: implement arithmetic operators
    public static CurrencyAmount operator +(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException();
        }

        return new CurrencyAmount(left.amount + right.amount, left.currency);
    }

    public static CurrencyAmount operator -(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException();
        }

        return new CurrencyAmount(left.amount - right.amount, left.currency);
    }

    public static CurrencyAmount operator *(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException();
        }

        return new CurrencyAmount(left.amount * right.amount, left.currency);
    }

    public static CurrencyAmount operator /(CurrencyAmount left, CurrencyAmount right)
    {
        if (left.currency != right.currency)
        {
            throw new ArgumentException();
        }

        return new CurrencyAmount(left.amount / right.amount, left.currency);
    }

    public static implicit operator double(CurrencyAmount right) => (double)right.amount;

    public static implicit operator decimal(CurrencyAmount right) => right.amount;
}
