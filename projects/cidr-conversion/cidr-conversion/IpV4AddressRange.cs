using CidrConversion;

public class IpV4AddressRange
{
  public IpV4Address StartingIp { get; set; }

  public IpV4Address EndingIp { get; set; }

  public SubnetMask SubnetMask { get; set; }

  public IpV4Address NetworkIdentification { get; set; }

  public IpV4Address BroadcastAddress { get; set; }

  public int Increment { get; set; }

}