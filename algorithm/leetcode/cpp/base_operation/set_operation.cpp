//
// Created by 裴星鑫 on 2022/3/30.
//
#include <iostream>
#include <set>
#include <vector>

using namespace std;

int main(int argc ,char** argv) {
  vector<int> d = {7, 9, 8, 6, 4, 3, 5, 2, 1};
  set<int> data(d.begin(), d.end());

  auto it = data.lower_bound(5);
  cout << *it << endl;
  return 0;
}
