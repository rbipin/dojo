def print_reverse(word):
	print ("Word Length:" ,len(word))
	reverse_word=""
	for letter in range(len(word)-1,-1,-1):
		print (word[letter],end="")

while True:
	word=input("\nEnter Word:")
	print_reverse(word)