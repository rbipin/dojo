"""
generate fibanocci series
Author: Bipin Radhakrishnan
"""


def fibano_Series(number):
	
	first=0
	
	second=1
	
	next=0
		
	print (first) 
	
	print(second)
	
	for count in range(1,number-2):
		
		next=first+second
		
		print (next)
		
		first=second
		
		second=next



num=int(input("Enter the Number of Fibanocci Seris to Print: "))

fibano_Series(num)	