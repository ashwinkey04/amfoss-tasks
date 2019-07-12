/*Run using https://www.onlinegdb.com/online_c++_debugger */

#include <iostream>
#include <math.h>

using namespace std;

int main()
{
    cout<<"Enter array size";
    int size;
    int mainsum=0, secsum=0;
    cin>>size;
    int a[size][size];
    for(int i=0; i<size ; ++i){
         for(int j=0; j<size ; ++j){
             cin>>a[i][j];
             cout<<" ";
         }
         cout<<endl;
    }
    for(int i=0; i<size ; ++i){
         for(int j=0; j<size ; ++j){
             if(i==j){
                 mainsum+=a[i][j];
             }
         }
    }
    for(int i=0; i<size ; ++i){
         for(int j=0; j<size ; ++j){
             if((i+j)==size-1){
                 secsum+=a[i][j];
             }
         }
    }
    int diff=secsum-mainsum;
    cout<<"The absolute difference between the diagonals is"<<abs(diff);

    return 0;
}

