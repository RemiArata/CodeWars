

def valid_braces(string):
    if string.count("(") != string.count("(") or string.count("{") != string.count("}") or string.count("[") != string.count("]") or len(string) % 2 != 0:
        return False

    mid = len(string) // 2
    for i in range(len(string)):
        if string[i] == "}":
            if string[i-1] == "{":
                continue
            elif i > mid - 1 and string[mid - ((i + 1) - mid)] == "{":
                    continue
            else:
                return False            
        elif string[i] == "]":
            if string[i-1] == "[":
                continue
            elif i > mid - 1 and string[mid - ((i + 1) - mid)] == "[":
                    continue
            else:
                return False
        elif string[i] == ")":
            if string[i-1] == "(":
                continue
            elif i > mid - 1 and string[mid - ((i + 1) - mid)] == "(":
                    continue
            else:
                return False
    return True

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
# print(valid_braces("{[]}[()]")) # true
