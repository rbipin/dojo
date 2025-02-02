using CidrConversion;
using FluentAssertions;

namespace CidrConverstionTest;

public class IpAddressV4Test
{
    public static TheoryData<string, byte, byte, byte, byte, string> ipAddressFullString
        => new()
        {
            {
                "192.168.101.10",
                192, 168, 101, 10,
                "11000000.10101000.01100101.00001010"
            },
            {
                "10.10.10.150",
                10,10,10,150,
                "00001010.00001010.00001010.10010110"
            },
            {
                "0.0.0.0",
                0,0,0,0,
                "00000000.00000000.00000000.00000000"
            },
            {
                "255.255.255.255",
                255,255,255,255,
                "11111111.11111111.11111111.11111111"
            }
        };

    [Theory]
    [MemberData(nameof(ipAddressFullString))]
    public void IpAddressV4_WithIPString_Should_Succeed(string ipAddressString,
        byte firstOctet, byte secondOctet,
        byte thirdOctet, byte fourthOctet,
        string binaryIpAddress)
    {
        // Act
        var ipAddress = new IpV4Address(ipAddressString);

        // Assert
        ipAddress.FirstOctet.Value.Should().Be(firstOctet);
        ipAddress.SecondOctet.Value.Should().Be(secondOctet);
        ipAddress.ThirdOctet.Value.Should().Be(thirdOctet);
        ipAddress.FourthOctet.Value.Should().Be(fourthOctet);
        ipAddress.Binary.Should().Be(binaryIpAddress);
        ipAddress.ToString().Should().Be(ipAddressString);
    }

    public static TheoryData<byte, byte, byte, byte, string> ipAddressWithByte
        => new()
        {
            {
                192, 168, 101, 10,
                "11000000.10101000.01100101.00001010"
            },
            {
                10, 10, 10, 150,
                "00001010.00001010.00001010.10010110"
            },
            {
                0, 0, 0, 0,
                "00000000.00000000.00000000.00000000"
            },
            {
                255, 255, 255, 255,
                "11111111.11111111.11111111.11111111"
            }
        };

    [Theory]
    [MemberData(nameof(ipAddressWithByte))]
    public void IpAddressV4_WithByte_Should_Succeed(byte firstOctet,
        byte secondOctet,
        byte thirdOctet, byte fourthOctet,
    string binaryIpAddress)
    {
        // Setup
        var ipAddressString = $"{firstOctet}.{secondOctet}.{thirdOctet}.{fourthOctet}";

        // Act
        var ipAddress = new IpV4Address(firstOctet, secondOctet, thirdOctet, fourthOctet);

        // Assert
        ipAddress.FirstOctet.Value.Should().Be(firstOctet);
        ipAddress.SecondOctet.Value.Should().Be(secondOctet);
        ipAddress.ThirdOctet.Value.Should().Be(thirdOctet);
        ipAddress.FourthOctet.Value.Should().Be(fourthOctet);
        ipAddress.Binary.Should().Be(binaryIpAddress);
        ipAddress.ToString().Should().Be(ipAddressString);
    }


    public static TheoryData<int, int, int, int, string> ipAddressWithInt
        => new()
        {
            {
                192, 168, 101, 10,
                "11000000.10101000.01100101.00001010"
            },
            {
                10, 10, 10, 150,
                "00001010.00001010.00001010.10010110"
            },
            {
                0, 0, 0, 0,
                "00000000.00000000.00000000.00000000"
            },
            {
                255, 255, 255, 255,
                "11111111.11111111.11111111.11111111"
            }
        };

    [Theory]
    [MemberData(nameof(ipAddressWithInt))]
    public void IpAddressV4_WithInteger_Should_Succeed(byte firstOctet,
        byte secondOctet,
        byte thirdOctet, byte fourthOctet,
    string binaryIpAddress)
    {
        // Setup
        var ipAddressString = $"{firstOctet}.{secondOctet}.{thirdOctet}.{fourthOctet}";

        // Act
        var ipAddress = new IpV4Address(firstOctet, secondOctet, thirdOctet, fourthOctet);

        // Assert
        ipAddress.FirstOctet.Value.Should().Be(firstOctet);
        ipAddress.SecondOctet.Value.Should().Be(secondOctet);
        ipAddress.ThirdOctet.Value.Should().Be(thirdOctet);
        ipAddress.FourthOctet.Value.Should().Be(fourthOctet);
        ipAddress.Binary.Should().Be(binaryIpAddress);
        ipAddress.ToString().Should().Be(ipAddressString);
    }

    public static TheoryData<string, string, string, string, string> ipAddressWithString
    => new()
    {
            {
                "192", "168", "101", "10",
                "11000000.10101000.01100101.00001010"
            },
            {
                "10", "10", "10", "150",
                "00001010.00001010.00001010.10010110"
            },
            {
                "0", "0", "0", "0",
                "00000000.00000000.00000000.00000000"
            },
            {
                "255", "255", "255", "255",
                "11111111.11111111.11111111.11111111"
            }
    };

    [Theory]
    [MemberData(nameof(ipAddressWithString))]
    public void IpAddressV4_WithString_Should_Succeed(byte firstOctet,
       byte secondOctet,
       byte thirdOctet, byte fourthOctet,
   string binaryIpAddress)
    {
        // Setup
        var ipAddressString = $"{firstOctet}.{secondOctet}.{thirdOctet}.{fourthOctet}";

        // Act
        var ipAddress = new IpV4Address(firstOctet, secondOctet, thirdOctet, fourthOctet);

        // Assert
        ipAddress.FirstOctet.Value.Should().Be(firstOctet);
        ipAddress.SecondOctet.Value.Should().Be(secondOctet);
        ipAddress.ThirdOctet.Value.Should().Be(thirdOctet);
        ipAddress.FourthOctet.Value.Should().Be(fourthOctet);
        ipAddress.Binary.Should().Be(binaryIpAddress);
        ipAddress.ToString().Should().Be(ipAddressString);
    }


}