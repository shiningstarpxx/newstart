//
// Created by 裴星鑫 on 2022/3/19.
//

#include <iostream>
#include <string>
using namespace std;

int main(){
  string str1 = "Java2Blog";
  string str2= "Blog"; //substring to be checked

  int position = 0;
  int index_str;

  while(( index_str = str1.find(str2, position)) != string::npos){
    cout<< "Substring " << "'" << str2 << "'" << " Found at index: " << index_str << endl;
    position = index_str + 1;
  }
  return 0;
}