class Robot:
    def __init__(self):
        self.phrase = "Thank you for teaching me "
        self.learned_words = ["thank", "you", "for", "teaching", "me", "i", "already", "know", "the", "word", "do", "not", "understand", "input"]
        self.known_word = "I already know the word "
        self.what = "I do not understand the input"
    
    def learn_word(self, wrd):
        if wrd.isalpha():
            if wrd.lower() not in self.learned_words:
                self.learned_words.append(wrd.lower())
                return self.phrase + wrd
            else:
                return self.known_word + wrd
        else:
            return self.what

vicky = Robot()

print(vicky.learn_word("hello"))
print(vicky.learn_word("world"))
print(vicky.learn_word("goodbye"))
print(vicky.learn_word("world"))
print(vicky.learn_word("World"))
