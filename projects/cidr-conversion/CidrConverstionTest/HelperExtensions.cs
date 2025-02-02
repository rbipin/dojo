namespace CidrConverstionTest;
public static class HelperExtensions
{
  public static int[] ToIntegerArray(this string value)
  {
    var integerArray = new int[value.Length];
    for (int count = 0; count < integerArray.Length; count++)
    {
      integerArray[count] = Convert.ToInt32(value[count].ToString());
    }
    return integerArray;
  }
}