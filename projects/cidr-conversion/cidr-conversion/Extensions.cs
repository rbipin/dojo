namespace CidrConversion;
public static class Extensions
{

  public static void ForEach(this string source, Action<int, char> action)
  {
    if (string.IsNullOrWhiteSpace(source))
    {
      throw new ArgumentNullException(nameof(source));
    }

    for (int i = 0; i < source.Length; i++)
    {
      action(i, source[i]);
    }
  }
}