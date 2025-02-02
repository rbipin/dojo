using System;
using System.Collections;

namespace cidr_conversion;

public class IpV4AddressRanges : IList<IpV4AddressRange>
{
  private readonly List<IpV4AddressRange> _ipRanges = new();
  public int Increment { get; set; }

  public IpV4AddressRange this[int index]
  {
    get => _ipRanges[index];
    set => _ipRanges[index] = value;
  }

  public int Count => _ipRanges.Count;

  public bool IsReadOnly => false;

  public void Add(IpV4AddressRange item) => _ipRanges.Add(item);

  public void Clear() => _ipRanges.Clear();

  public bool Contains(IpV4AddressRange item)
  {
    throw new NotImplementedException();
  }

  public void CopyTo(IpV4AddressRange[] array, int arrayIndex)
  {
    throw new NotImplementedException();
  }

  public IEnumerator<IpV4AddressRange> GetEnumerator()
  {
    throw new NotImplementedException();
  }

  public int IndexOf(IpV4AddressRange item)
  {
    throw new NotImplementedException();
  }

  public void Insert(int index, IpV4AddressRange item)
  {
    throw new NotImplementedException();
  }

  public bool Remove(IpV4AddressRange item)
  {
    throw new NotImplementedException();
  }

  public void RemoveAt(int index)
  {
    throw new NotImplementedException();
  }

  IEnumerator IEnumerable.GetEnumerator()
  {
    throw new NotImplementedException();
  }
}
