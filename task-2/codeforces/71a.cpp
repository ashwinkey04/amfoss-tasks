#include<iostream>
#include<string.h>

using namespace std;

int main()
{
    int n;
    cin>>n;
    string str[n];
    for(int i = 0; i<n; ++i){
        cin>>str[i];
    }
    
    for(int i = 0; i<n; ++i){
        if(str[i].length()>10){
            int x =str[i].length();
            cout<<str[i][0]<<x-2<<str[i][x-1]<<endl;
        }
        else
            cout<<"
    }
}
