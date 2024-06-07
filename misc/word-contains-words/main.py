import json

with open("words.json") as f:
    data = f.read()

words = json.loads(data)

words = {word : 1 for word in words}
# The original car talk puzzler is:
#   What 7 letter word contains 8 words (including itself!)?

# The difficulty here is getting a reasonable english dictionary file.
# First considered https://github.com/dwyl/english-words has many "words" that don't seem legitimate.
# IE: 'zwitter': ['z', 'w', 'wi', 'wit', 'i', 'it', 't', 't', 'te', 'e']
# Much better was CustomDictionary: https://raw.githubusercontent.com/sujithps/Dictionary/master/Oxford%20English%20Dictionary.txt

seven_letter_words = [word for word in words if len(word) == 7]

print(seven_letter_words)

words_in_word = {}

for word in seven_letter_words:
    print(f"Checking {word}")
    words_in = []
    for start in range(0,6):
        for end in range(start + 1, 8):
            if word[start:end] in words:
                words_in.append(word[start:end])

    words_in_word[word] = words_in


word_count_in_word = {word : len(words) for (word, words) in words_in_word.items()}


#print(word_count_in_word)
solutions = {word: words for (word, words) in words_in_word.items() if len(words) == 8}
print(f"{len(solutions)} potential solutions")
print(solutions)
print(words_in_word["therein"])