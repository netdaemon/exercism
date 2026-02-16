public static class CentralBank
{
    public static string DisplayDenomination(long @base, long multiplier)
    {
        try
        {
            var result = checked(@base * multiplier);
            
            return result.ToString();
        }
        catch (OverflowException ex)
        {
            return "*** Too Big ***";
        }
    }

    public static string DisplayGDP(float @base, float multiplier)
    {
        try
        {
            var result = checked(@base * multiplier);

            if (result == float.NegativeInfinity || result == float.PositiveInfinity)
            {
                throw new OverflowException();
            }

            return result.ToString();
        }
        catch (OverflowException ex)
        {
            return "*** Too Big ***";
        }
    }

    public static string DisplayChiefEconomistSalary(decimal salaryBase, decimal multiplier)
    {
        try
        {
            var result = checked(salaryBase * multiplier);
            return result.ToString();
        }
        catch (OverflowException ex)
        {
            return "*** Much Too Big ***";
        }
    }
}
