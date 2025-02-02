using FluentAssertions;
using CidrConversion;
using CidrConverstionTest;

namespace CidrConversionTest;
public class OctetTest
{

  [Theory]
  [InlineData("0", "00000000")]
  [InlineData("1", "00000001")]
  [InlineData("3", "00000011")]
  [InlineData("255", "11111111")]
  [InlineData("100", "01100100")]
  public void Octet_WithString_Should_Succeed(string value, string expectedBinary)
  {
    var octet = new Octet(value);

    var expectedBinaryArray = expectedBinary.ToIntegerArray();
    var expectedValue = Convert.ToByte(value);
    //Assert
    octet.Value.Should().Be(expectedValue);
    octet.BinaryString.Should().Be(expectedBinary);
    octet.Binary.Should().BeEquivalentTo(expectedBinaryArray);
  }

  [Theory]
  [InlineData(0, "00000000")]
  [InlineData(1, "00000001")]
  [InlineData(2, "00000010")]
  [InlineData(3, "00000011")]
  [InlineData(4, "00000100")]
  [InlineData(5, "00000101")]
  [InlineData(255, "11111111")]
  public void Octet_WithByte_Should_Succeed(int value, string expected)
  {
    //Setup
    var binaryArray = expected.ToIntegerArray();

    //Act
    var octet = new Octet(value);
    byte valueInByte = octet;
    int valueInInt = octet;
    int[] binaryValue = octet;

    // Assert
    octet.Binary.Should().BeEquivalentTo(binaryArray);
    binaryValue.Should().BeEquivalentTo(binaryArray);
    octet.BinaryString.Should().Be(expected);
    octet.Value.Should().Be(Convert.ToByte(value));
    valueInByte.Should().Be(Convert.ToByte(valueInByte));
    valueInInt.Should().Be(value);
  }

  [Theory]
  [InlineData("00000000", 0)]
  [InlineData("00000001", 1)]
  [InlineData("00000010", 2)]
  [InlineData("00000011", 3)]
  [InlineData("00000100", 4)]
  [InlineData("00000101", 5)]
  [InlineData("11111111", 255)]
  [InlineData("10000000", 128)]
  [InlineData("11000000", 192)]
  public void Octet_WithBinaryArray_Should_Succeed(string value, byte expected)
  {
    //Setup
    var binaryArray = value.ToIntegerArray();

    //Act
    var octet = new Octet(binaryArray);
    byte valueInByte = octet;
    int valueInInt = octet;
    int[] binaryValue = octet;

    // Assert
    octet.Binary.Should().BeEquivalentTo(binaryArray);
    binaryValue.Should().BeEquivalentTo(binaryArray);
    octet.BinaryString.Should().Be(value);
    octet.Value.Should().Be(Convert.ToByte(expected));
    valueInByte.Should().Be(Convert.ToByte(expected));
    valueInInt.Should().Be(Convert.ToInt32(expected));
  }

  [Theory]
  [InlineData(-1)]
  [InlineData(256)]
  [InlineData(300)]
  public void Octet_InInt_WithByte_Should_ThrowArgumentOutOfRange(int value)
  {
    var act = () => new Octet(value);
    //Assert
    act.Should().Throw<ArgumentOutOfRangeException>();

  }

  [Theory]
  [InlineData("abcdefgh")]
  [InlineData("abcd1010")]
  [InlineData("1101a100")]
  [InlineData("12345678")]
  public void Octet_BinaryAsString_Invalid_FormatException(string value)
  {
    var act = () => new Octet(value);
    //Assert
    act.Should().Throw<FormatException>().WithMessage("Binary value in not in the correct format");
  }

  [Theory]
  [InlineData("300")]
  [InlineData("256")]
  [InlineData("-1")]
  [InlineData("5000")]
  public void Octet_DecimalAsString_Should_Throw_OverflowException(string value)
  {
    var act = () => new Octet(value);
    //Assert
    act.Should().Throw<OverflowException>();
  }
}

