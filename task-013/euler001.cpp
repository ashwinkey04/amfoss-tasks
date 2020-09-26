#include <map>
#include <set>
#include <list>
#include <cmath>
#include <ctime>
#include <deque>
#include <queue>
#include <stack>
#include <string>
#include <bitset>
#include <cstdio>
#include <limits>
#include <vector>
#include <climits>
#include <cstring>
#include <cstdlib>
#include <fstream>
#include <numeric>
#include <sstream>
#include <iostream>
#include <algorithm>
#include <unordered_map>

using namespace std;


int main(){
    int t;
    cin >> t;
    int sum[t];
    for(int a0 = 0; a0 < t; a0++){
        int n;
        cin >> n;
        sum[a0] = 0;
        for(int i =2; i<n; ++i){
            if (i%3==0 || i%5==0){
                sum[a0]+=i;
            }
        }
    }
    for (int j=0; j<t ; ++j){
        cout<<sum[j]<<endl;
    }
    return 0;
}

