def valid_braces(string):
    pairs = {"{": "}", "(": ")", "[": "]"}
    stack = []
    for itm in string:
        if itm in pairs.keys():
            stack.append(itm)
        else:
            if len(stack) == 0 or pairs[stack.pop()] != itm:
                return False
    return True if len(stack) == 0 else False

# print(valid_braces("()")) # True
# print(valid_braces("(}")) # false
# print(valid_braces("[]")) # true
# print(valid_braces("{}")) # true
# print(valid_braces("(((({{")) # false
# print(valid_braces("{}[]")) # true
# print(valid_braces("{[]}")) # true
# print(valid_braces("[(])")) # false

# print(valid_braces("{}()[]")) # true
# print(valid_braces("([{}])")) # true

# print(valid_braces("{}({})[]")) # true
print(valid_braces("{[]}[()]")) # true