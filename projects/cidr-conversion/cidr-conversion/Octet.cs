
namespace CidrConversion;

public class Octet
{

  /// <summary>
  /// Create an IP Octet
  /// </summary>
  /// <param name="decimalValue">Pass the decimal IP Octet</param> <summary>
  public Octet(byte decimalValue)
  {
    Binary = ToBinary(decimalValue);
  }

  /// <summary>
  /// Create an IP Octet
  /// </summary>
  /// <param name="decimalValue">Pass the decimal IP Octet</param>
  public Octet(int decimalValue)
  {
    if (decimalValue is < 0 or > 255)
    {
      throw new ArgumentOutOfRangeException(nameof(decimalValue));
    }
    var valueInBytes = Convert.ToByte(decimalValue);
    Binary = ToBinary(valueInBytes);
  }

  /// <summary>
  /// Create an IP Octet
  /// </summary>
  /// <param name="value">Pass the decimal IP Octet</param>
  public Octet(string value)
  {
    if (string.IsNullOrWhiteSpace(value))
    {
      throw new ArgumentNullException(nameof(value));
    }

    var isBinary = IsBinary(value);

    if (isBinary)
    {
      Binary = StringToBinary(value);
      return;
    }

    var decimalValue = Convert.ToByte(value);
    Binary = ToBinary(decimalValue);
  }

  /// <summary>
  /// Create an IP Octet
  /// </summary>
  /// <param name="binaryValue">Pass the binary IP Address</param>
  public Octet(int[] binaryValue)
  {
    if (binaryValue.Length == 0)
    {
      throw new ArgumentOutOfRangeException(nameof(binaryValue));
    }
    if (!Array.TrueForAll(binaryValue, x => x is 0 or 1))
    {
      throw new ArgumentOutOfRangeException($"{nameof(binaryValue)}");
    }
    Binary = binaryValue;
  }

  /// <summary>
  /// The IP Octet Value
  /// </summary>
  /// <returns>Octet Decimal value</returns>
  public byte Value => GetDecimalValue();

  /// <summary>
  /// The Binary IP Address Value
  /// </summary>
  /// <value></value>
  public int[] Binary { get; } = [0, 0, 0, 0, 0, 0, 0, 0];

  /// <summary>
  /// The On Bits in the Octet
  /// </summary>
  /// <returns>On bits</returns>
  public int OnBits => Binary.Count(x => x == 1);

  /// <summary>
  /// The Off bits in the octet
  /// </summary>
  public int Offbits => 8 - OnBits;

  /// <summary>
  /// Return Value in the string format
  /// </summary>
  /// <returns>Value of the Octet</returns>
  public override string ToString() => this.Value.ToString();

  /// <summary>
  /// 
  /// </summary>
  /// <returns></returns>
  public string BinaryString => string.Join("", Binary);

  public static implicit operator string(Octet octet) => octet.Value.ToString();
  public static implicit operator byte(Octet octet) => octet.Value;
  public static implicit operator int(Octet octet) => octet.Value;
  public static implicit operator int[](Octet octet) => octet.Binary;

  public static explicit operator Octet(byte value) => new(value);
  public static explicit operator Octet(int value) => new(value);
  public static explicit operator Octet(string value) => new(value);

  private static int[] StringToBinary(string binaryValue)
  {
    if (string.IsNullOrWhiteSpace(binaryValue))
    {
      throw new ArgumentNullException(nameof(binaryValue));
    }
    if (binaryValue.Length != 8)
    {
      throw new ArgumentOutOfRangeException(nameof(binaryValue));
    }

    int[] binary = new int[8];
    binaryValue.ForEach((position, value) =>
    {
      binary[position] = Convert.ToInt32(value);
    });

    return binary;
  }

  private static int[] ToBinary(byte value)
  {
    var newBinary = DecimalToBinary(value);
    var remainingToPopulate = 8 - newBinary.Count;
    var binary = Enumerable.Repeat<int>(0, remainingToPopulate).ToList();
    binary.AddRange(newBinary);
    return binary.ToArray();
  }


  private static List<int> DecimalToBinary(byte value)
  {
    if (value < 2)
    {
      return [value];
    }

    var newBinary = new List<int>();
    while (value > 0)
    {
      var reminder = value % 2;
      value /= 2;
      newBinary.Add(reminder);
    }
    newBinary.Reverse();
    return newBinary.ToList();
  }

  private static int ToDecimal(int position) => Convert.ToInt32(Math.Pow(2, position));

  private byte GetDecimalValue()
  {
    var total = 0;
    for (int count = Binary.Length - 1; count >= 0; count--)
    {
      if (Binary[count] != 1)
      {
        continue;
      }
      total += ToDecimal(7 - count);
    }
    return Convert.ToByte(total);
  }

  private static bool IsBinary(string value)
  {
    var result = value.Length == 8
      && Array.TrueForAll(value.ToCharArray(), x => x is '0' or '1');

    return value.Length == 8 && !result
      ? throw new FormatException("Binary value in not in the correct format")
      : result;
  }

}
