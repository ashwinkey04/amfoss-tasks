// Execute using http://cpp.sh
#include <iostream>
#include <cstring>

int main()
{
  char ft[100];
  std::cout << "Enter string ";
  std::cin>>ft;
  int n = strlen(ft);
  int cnt=1;
  int a=0;
  char cmp='0';
  for(int i=0;i<n;++i){
      if(ft[i]==cmp){
          cnt++;
      }
      else 
      cnt=1;
      if(a<cnt){
          a=cnt;
      }
      cmp=ft[i];
  }
  if(a>6){
      std::cout<<"Yes";
  }
  else
  std::cout<<"Nope";
}
