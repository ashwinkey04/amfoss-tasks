#include <iostream>
#include <math.h>
using namespace std;

long maxPF(long n)  
{   long max=1; 
    while (n % 2 == 0)  
    {          
        n = n/2;
        max=2;  
    }
    for (long i = 3; i <= sqrt(n); i = i + 2)  
    {    
        while (n % i == 0)  
        {             
            max=i;
            n = n/i;  
        }  
    }  
    if (n > 2){          
        max=n;
    }
    return max;  
}
int main(){
    long t;
    cin >> t;
    long sol[t];
    for(long a0 = 0; a0 < t; a0++){
        long n;
        cin >> n;
        sol[a0]=maxPF(n);               
    }
    for(long a0 = 0; a0 < t; a0++){
       cout<<sol[a0]<<endl;
    }
    return 0;
}
