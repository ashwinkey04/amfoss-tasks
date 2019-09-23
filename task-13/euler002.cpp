#include <iostream>

using namespace std;

int main(){
    int t;
    cin >> t;
    long sol[t];
    long fib[100];
    fib[0]=1;
    fib[1]=2;
    for(int i=2; i<100;++i){
            fib[i]=fib[i-1]+fib[i-2];
    }
    for(int a0 = 0; a0 < t; a0++){
        long n;
        cin >> n;
        sol[a0]=2;    
        for(int i=2; i<100;++i){
            if(fib[i]%2==0&&fib[i]<n){
                sol[a0]+=fib[i];
            }
            else if(fib[i]>n){
                break;
            }        
        }        
    }
    for(int a0 = 0; a0 < t; a0++){
        cout<<sol[a0]<<endl;
    }      
    return 0;
}

