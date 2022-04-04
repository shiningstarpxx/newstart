//
// Created by 裴星鑫 on 2022/3/30.
//
#include <queue>
#include <iostream>

using namespace std;

int main(int argc ,char** argv) {
  priority_queue<int, vector<int>, greater<int>> q;
  vector<int> d = {3, 2, 3, 9, 8, 7};
  for (auto v : d) q.push(v);
  while (!q.empty()) {
    cout << q.top() << endl;
    q.pop();
  }
  return 0;
}