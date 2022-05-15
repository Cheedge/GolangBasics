

def vanillaFib(i, num)-> int:
    if len(num)<=i & i>2 :
        n = vanillaFib(i-2, num)+vanillaFib(i-1, num)
        num.append(n)
        return n
    else:
        return num[i]

def dpFib(i) -> int:
    k = 2
    num = [0]*(i+1)
    print(k, len(num))
    num[0], num[1], num[2] = 0, 1, 1
    while(k<=i):
        num[k] = num[k-2] + num[k-1]
        k += 1
        print(num)
    return num[k-1]

if __name__=="__main__":
    num = [0, 1, 1]
    # res = vanillaFib(10, num)
    res1 = dpFib(100)
    print(res1, num)