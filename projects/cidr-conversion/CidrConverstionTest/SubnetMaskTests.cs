using Xunit;
using FluentAssertions;
using CidrConversion;

namespace CidrConverstionTest;

public class SubnetMaskTests
{
  [Theory]
  [InlineData("255.255.255.255", 32)]
  [InlineData("255.255.255.254", 31)]
  [InlineData("255.255.255.252", 30)]
  [InlineData("255.255.255.248", 29)]
  [InlineData("255.255.255.240", 28)]
  [InlineData("255.255.255.224", 27)]
  [InlineData("255.255.255.192", 26)]
  [InlineData("255.255.255.128", 25)]
  [InlineData("255.255.255.0", 24)]
  [InlineData("255.255.254.0", 23)]
  [InlineData("255.255.252.0", 22)]
  [InlineData("255.255.248.0", 21)]
  [InlineData("255.255.240.0", 20)]
  [InlineData("255.255.224.0", 19)]
  [InlineData("255.255.192.0", 18)]
  [InlineData("255.255.128.0", 17)]
  [InlineData("255.255.0.0", 16)]
  [InlineData("255.254.0.0", 15)]
  [InlineData("255.252.0.0", 14)]
  [InlineData("255.248.0.0", 13)]
  [InlineData("255.240.0.0", 12)]
  [InlineData("255.224.0.0", 11)]
  [InlineData("255.192.0.0", 10)]
  [InlineData("255.128.0.0", 9)]
  [InlineData("255.0.0.0", 8)]
  [InlineData("254.0.0.0", 7)]
  [InlineData("252.0.0.0", 6)]
  [InlineData("248.0.0.0", 5)]
  [InlineData("240.0.0.0", 4)]
  [InlineData("224.0.0.0", 3)]
  [InlineData("192.0.0.0", 2)]
  [InlineData("128.0.0.0", 1)]
  [InlineData("0.0.0.0", 0)]
  public void SubnetMask_AsString_Is_Valid(string subnetMaskIp, int expectedCidr)
  {
    var subnetMask = new SubnetMask(subnetMaskIp);

    subnetMask.Cidr.Should().Be(expectedCidr);
  }

  [Theory]
  [InlineData(255, 255, 255, 255, 32)]
  [InlineData(255, 255, 255, 254, 31)]
  [InlineData(255, 255, 255, 252, 30)]
  [InlineData(255, 255, 255, 248, 29)]
  [InlineData(255, 255, 255, 240, 28)]
  [InlineData(255, 255, 255, 224, 27)]
  [InlineData(255, 255, 255, 192, 26)]
  [InlineData(255, 255, 255, 128, 25)]
  [InlineData(255, 255, 255, 0, 24)]
  [InlineData(255, 255, 254, 0, 23)]
  [InlineData(255, 255, 252, 0, 22)]
  [InlineData(255, 255, 248, 0, 21)]
  [InlineData(255, 255, 240, 0, 20)]
  [InlineData(255, 255, 224, 0, 19)]
  [InlineData(255, 255, 192, 0, 18)]
  [InlineData(255, 255, 128, 0, 17)]
  [InlineData(255, 255, 0, 0, 16)]
  [InlineData(255, 254, 0, 0, 15)]
  [InlineData(255, 252, 0, 0, 14)]
  [InlineData(255, 248, 0, 0, 13)]
  [InlineData(255, 240, 0, 0, 12)]
  [InlineData(255, 224, 0, 0, 11)]
  [InlineData(255, 192, 0, 0, 10)]
  [InlineData(255, 128, 0, 0, 9)]
  [InlineData(255, 0, 0, 0, 8)]
  [InlineData(254, 0, 0, 0, 7)]
  [InlineData(252, 0, 0, 0, 6)]
  [InlineData(248, 0, 0, 0, 5)]
  [InlineData(240, 0, 0, 0, 4)]
  [InlineData(224, 0, 0, 0, 3)]
  [InlineData(192, 0, 0, 0, 2)]
  [InlineData(128, 0, 0, 0, 1)]
  [InlineData(0, 0, 0, 0, 0)]
  public void SubnetMask_AsInt_Is_Valid(int firstOctetValue,
    int secondOctetValue,
    int thirdOctetValue,
    int fourthOctetValue,
    int expectedCidr)
  {
    var subnetMask = new SubnetMask(firstOctetValue,
                                    secondOctetValue,
                                    thirdOctetValue,
                                    fourthOctetValue);

    subnetMask.Cidr.Should().Be(expectedCidr);
  }
}
