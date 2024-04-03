"""
This code demonstrates use of .append key and way to insert value into a list
"""
def censor(text,word):
	wordsinSen=text.split()
	newSentense=[]
	for w in wordsinSen:
		if w.lower()==word.lower():
			newSentense.append("*"*len(w))
		else:
			newSentense.append(w)
	return newSentense
	
print (" ".join(censor("this hack is wack hack", "hack")))