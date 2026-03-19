from math import sqrt
def check_root(strng):
    if strng.count(",") != 3:
        return 'incorrect input'
    
    nums = []

    for val in strng.split(","):
        if is_num(val) != 'NaN':
            if len(nums) != 0:
                if int(val) - nums[-1] == 1:
                    nums.append(int(val))
                else:
                    return 'not consecutive'
            else:
                nums.append(int(val))
        else:
            return 'incorrect input'

    acc = 1
    for num in nums:
        acc *= num
    return f"{acc + 1}, {int(sqrt(acc+1))}"

def is_num(strng):
    try:
        num = int(strng)
        return num
    except:
        return 'NaN'



print(check_root("-4,-3,-2,-1"))
print(check_root("-1,0,1,2"))
