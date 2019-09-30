t=int(input())
a=[]
for a0 in range(t):
    n=int(input())
    pr=[]
    nw=[]
    for num in range(1, n + 1):   
       if num > 1:
           for i in range(2, num):
               if (num % i) == 0:
                   break
           else:               
               pr.append(num)
    for q1 in range(len(pr)):
        tmp=1
        while((pr[q1]**tmp)<=n):
            tmp=tmp+1
        nw.append(pr[q1]**(tmp-1))
    prod=1
    for q1 in nw:
        prod=prod*q1
    a.append(prod)
for k in a:
    print(k)
