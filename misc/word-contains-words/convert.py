import json
import string 

with open("words.txt") as f:
    data = f.read().split("\n")
data = [ data.split(" ")[0].lower() for data in data if data != ""]
data = [ data for data in data if "'" not in data and "-" not in data and len(data) <= 7]


denylist = [
    "",
    "b",
    "c",
    "d",
    "e",
    "f",
    "g",
    "h",
    "j",
    "k",
    "l",
    "m",
    "n",
    "o",
    "p",
    "q",
    "r",
    "s",
    "t",
    "u",
    "v",
    "w", 
    "x",
    "y", 
    "z",

    # from inspection on potential solutions
    "ba",
    "li",
    "ac",
    "ca",
    "de",
    "em",
    "ce",
    "et",
    "ag",
    "po",
    "er",
    "te",
    "es",
    "ti",
    "wi",
    "ni",
    "wa",
    "si",
    "ra",
    "xi",
    "ec",
    "re", 
    "da",
    "pto",
    "ga",
    "gi",
    "cu",
    "ar",
    "se", 
]

consonants = "bcdfghjklmnpqrstvwxyz"

rules = [
    # all words should have a vowel
    lambda w : not all([ letter in consonants for letter in w]),
    # words should only have alphabetical characters
    lambda w : all([letter in string.ascii_lowercase for letter in w]),
    # words should not be in the manual denylist
    lambda w : not w in denylist
]

words = [ word for word in data if all([rule(word) for rule in rules]) ]

print(json.dumps(words, indent=2))