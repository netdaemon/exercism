public class Orm
{
    private Database database;

    public Orm(Database database)
    {
        this.database = database;
    }

    public void Write(string data)
    {
        using var db = this.database;
        db.BeginTransaction();
        db.Write(data);
        db.EndTransaction();
    }

    public bool WriteSafely(string data)
    {
        try
        {
            Write(data);
            return true;
        }
        catch (InvalidOperationException)
        {
            return false;
        }
    }
}
