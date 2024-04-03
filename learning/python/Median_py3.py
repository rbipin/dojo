from math import floor, ceil
def median(sequence):
	middleIndex=0
	firstItemIndex=0
	secondItemIndex=0
	medianValue=0
	sequence.sort()
	print (sequence)
	if (len(sequence)%2)!=0:
		#odd numbers
		middleIndex=floor(len(sequence)/float(2))
		medianValue=sequence[middleIndex]
	else:
		#even number
		firstItemIndex=floor((len(sequence)-1)/2)
		secondItemIndex=firstItemIndex+1
		medianValue=float(sequence[firstItemIndex]+sequence[secondItemIndex])/float(2)
	return medianValue