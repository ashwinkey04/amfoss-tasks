#include <bits/stdc++.h>
#include<iostream>
#include<sstream>
#include<string>

using namespace std;

/*
 * Complete the timeConversion function below.
*/

int main()
{
    string s, r, t;
    cin>>s;
    if((s[8]=='A')){
        if(s[1]=='2'&&s[0]=='1'){
            r[0]='0';
            r[1]='0';
        }
        else{
            r[0]=s[0];
            r[1]=s[1];
        }
        
    }
    else {
        if(s[1]=='2'&&s[0]=='1'){
            r[0]=s[0];
            r[1]=s[1];
        }
        else {
            int num;
            string fst = s.substr(0,1);
            std::istringstream iss (fst);
            iss >> num;
            num+=12;            
            
            string outt;
            stringstream ss;
            ss << num;
            outt = ss.str();
            r.substr(0,1)=outt;
        }
    }
    
    for(int i=2; i<8;++i){
            r[i]=s[i];
        }
        
    cout<<r;

    return 0;
}
