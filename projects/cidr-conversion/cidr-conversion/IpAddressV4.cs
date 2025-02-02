namespace CidrConversion;

public class IpV4Address
{
  public Octet FirstOctet { get; }
  public Octet SecondOctet { get; }
  public Octet ThirdOctet { get; }
  public Octet FourthOctet { get; }

  public IpV4Address(string ipv4Address)
  {
    var split = ipv4Address.Split('.');
    if (split.Length != 4)
    {
      throw new ArgumentOutOfRangeException(nameof(ipv4Address));
    }
    FirstOctet = new Octet(Convert.ToByte(split[0]));
    SecondOctet = new Octet(Convert.ToByte(split[1]));
    ThirdOctet = new Octet(Convert.ToByte(split[2]));
    FourthOctet = new Octet(Convert.ToByte(split[3]));
  }
  public IpV4Address(string firstOctet,
    string secondOctet,
    string thirdOctet,
    string fourthOctet)
  {
    FirstOctet = new Octet(firstOctet);
    SecondOctet = new Octet(secondOctet);
    ThirdOctet = new Octet(thirdOctet);
    FourthOctet = new Octet(fourthOctet);
  }

  public IpV4Address(int[] firstOctet,
    int[] secondOctet,
    int[] thirdOctet,
    int[] fourthOctet)
  {
    FirstOctet = new Octet(firstOctet);
    SecondOctet = new Octet(secondOctet);
    ThirdOctet = new Octet(thirdOctet);
    FourthOctet = new Octet(fourthOctet);
  }

  public IpV4Address(int firstOctet, int secondOctet, int thirdOctet, int fourthOctet)
  {
    FirstOctet = new Octet(firstOctet);
    SecondOctet = new Octet(secondOctet);
    ThirdOctet = new Octet(thirdOctet);
    FourthOctet = new Octet(fourthOctet);
  }

  public IpV4Address(byte firstOctet, byte secondOctet, byte thirdOctet, byte fourthOctet)
  {
    FirstOctet = new Octet(firstOctet);
    SecondOctet = new Octet(secondOctet);
    ThirdOctet = new Octet(thirdOctet);
    FourthOctet = new Octet(fourthOctet);
  }

  public IpV4Address(Octet firstOctet, Octet secondOctet, Octet thirdOctet, Octet fourthOctet)
  {
    FirstOctet = firstOctet;
    SecondOctet = secondOctet;
    ThirdOctet = thirdOctet;
    FourthOctet = fourthOctet;
  }

  public static implicit operator string(IpV4Address ipaddressV4) => ipaddressV4.ToString();

  public static explicit operator IpV4Address(string ipaddressV4) => new(ipaddressV4);

  public override string ToString()
    => $"{FirstOctet}.{SecondOctet}.{ThirdOctet}.{FourthOctet}";

  public string Binary
   => $"{FirstOctet.BinaryString}.{SecondOctet.BinaryString}.{ThirdOctet.BinaryString}.{FourthOctet.BinaryString}";

}
