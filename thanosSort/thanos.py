def thanos_sort(arr):
    if arr == sorted(arr):
        return len(arr)
    else:
        if len(arr) % 2 == 0:
            return max(thanos_sort(arr[:len(arr)//2]), thanos_sort(arr[len(arr)//2:]))
        else:
            return max(thanos_sort(arr[:(len(arr)//2)+1]), thanos_sort(arr[(len(arr)//2)+1:]))
        
print(thanos_sort([1,2,3,7,5,6]))