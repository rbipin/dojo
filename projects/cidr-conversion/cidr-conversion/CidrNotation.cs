namespace CidrConversion;

public struct CidrNotation
{
  public IpV4Address Prefix { get; set; }
  public int Suffix { get; set; }

  public CidrNotation(IpV4Address ipaddressV4, int suffix)
  {
    Prefix = ipaddressV4;
    Suffix = suffix;
  }

  public override string ToString() => $"{this.Prefix}/{this.Suffix}";

  public static implicit operator string(CidrNotation cidrNotation)
    => cidrNotation.ToString();
}
