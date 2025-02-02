
namespace CidrConversion;

public class Cidr
{

  /// <summary>
  /// 
  /// </summary>
  /// <param name="startingIp"></param>
  /// <param name="endingIp"></param>
  /// <returns></returns> <summary>
  public static CidrNotation ToCidr(string startingIp, string endingIp)
  {
    var startIp = new IpV4Address(startingIp);
    var endIp = new IpV4Address(endingIp);

    return ToCidr(startIp, endIp);
  }

  /// <summary>
  /// 
  /// </summary>
  /// <param name="startingIp"></param>
  /// <param name="endingIp"></param>
  /// <returns></returns> <summary>
  public static CidrNotation ToCidr(IpV4Address startingIp, IpV4Address endingIp)
  {
    // To Do: Need to calculate
    return new CidrNotation(startingIp, 0);
  }

  /// <summary>
  /// 
  /// </summary>
  /// <param name="ipAddressV4"></param>
  /// <param name="subnetMask"></param>
  /// <returns></returns>
  public static CidrNotation ToCidr(IpV4Address ipAddressV4, SubnetMask subnetMask)
  {
    return new CidrNotation(ipAddressV4, subnetMask.Cidr);
  }

  /// <summary>
  /// 
  /// </summary>
  /// <param name="ipaddress"></param>
  /// <param name="cidr"></param>
  /// <returns></returns>
  public static IEnumerable<IpV4AddressRange> ToIpAddressV4Range(IpV4Address ipaddress, int cidr)
  {
    return new List<IpV4AddressRange>();
  }

}