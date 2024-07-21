public static void Main(string[] args)
{
    string url = args[0];
    int index = url.IndexOf("?");
    string[] parms = url.Remove(0, index+1).Split('&');
    foreach(string parm in parms)
    {
        Console.WriteLine(parm);
    }
}
