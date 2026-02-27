public static class Hamming
{
    public static int Distance(string firstStrand, string secondStrand)
    {
        if (firstStrand.Length != secondStrand.Length)
        {
            throw new ArgumentException();
        }

        int distance = firstStrand.Length;

        for (var i = 0; i < firstStrand.Length; i++)
        {
            if (firstStrand[i] == secondStrand[i])
            {
                distance--;
            }
        }

        return distance;
    }
}