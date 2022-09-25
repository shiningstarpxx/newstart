//
// Created by 裴星鑫 on 2022/9/25.
//

#include <cstdio>
#include <iostream>
#include <stack>

using namespace std;

const int MAXN = 1000 + 10;

int main() {
  int n , target[MAXN];  // 看上去acm都喜欢放在外面

  while (true) {
    cin >> n;
    if (n == 0)  {
      break;
    }
    while (true) {
      int i = 1;
      cin >> target[i];
      if (target[i] == 0) {
        cout << endl;
        break;
      }
      for (int i = 2; i <= n; i++) cin >> target[i];
      int j = 1;
      i = 1;
      stack<int> s;
      bool right = true;
      while (j <= n) {
        if (i == target[j]) {i++, j++;}
        else if (!s.empty() && s.top() == target[j]) {
          s.pop();
          j++;
        } else if (i <= n) {
          s.push(i);
          i++;
        } else {
          right = false;
          break;
        }
      }
      printf("%s\n", right ? "Yes":"No");
    }
  }

  return 0;
}