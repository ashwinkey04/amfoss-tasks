// Example program
#include <iostream>
#include <string>

using namespace std;

int main()
{
    int n,x=0;
    cin>>n;
    string str[n];
    for(int i=0; i<n;++i){
        cin>>str[i];
    }
    for(int i=0; i<n;++i){
        
        if(str[i][0]=='-'){
        --x;
        }
        else if (str[i][0]=='+'){
        ++x;
        }    
        else if(str[i][1]=='-'){
            --x;
        }
        else if(str[i][1]=='+'){
        ++x;
        
        }
    }
    cout<<x;
}
