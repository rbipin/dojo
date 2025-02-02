using System;

namespace CidrConversion;

public class SubnetMask : IpV4Address
{
  public SubnetMask(string ipv4Address) : base(ipv4Address)
  {
  }

  public SubnetMask(string firstOctet, string secondOctet, string thirdOctet, string fourthOctet) : base(firstOctet, secondOctet, thirdOctet, fourthOctet)
  {
  }

  public SubnetMask(int[] firstOctet, int[] secondOctet, int[] thirdOctet, int[] fourthOctet) : base(firstOctet, secondOctet, thirdOctet, fourthOctet)
  {
  }

  public SubnetMask(int firstOctet, int secondOctet, int thirdOctet, int fourthOctet) : base(firstOctet, secondOctet, thirdOctet, fourthOctet)
  {
  }

  public SubnetMask(byte firstOctet, byte secondOctet, byte thirdOctet, byte fourthOctet) : base(firstOctet, secondOctet, thirdOctet, fourthOctet)
  {
  }

  public SubnetMask(Octet firstOctet, Octet secondOctet, Octet thirdOctet, Octet fourthOctet) : base(firstOctet, secondOctet, thirdOctet, fourthOctet)
  {
  }

  public int Cidr
   => FirstOctet.OnBits + SecondOctet.OnBits + ThirdOctet.OnBits + FourthOctet.OnBits;

}
