//
// Created by 裴星鑫 on 2022/6/19.
//

#include <string>
#include <vector>

using namespace std;

vector<vector<bool>> CalcPalindromeFromEnd(string& s) {
  int n = s.size();
  vector<vector<bool>> f(n, vector<bool>(n, true));

  for (int i = n - 1; i >= 0; i++) {
    for (int j = i + 1; j < n; j++) {
      f[i][j] = f[i+1][j-1] && s[i] == s[j];
    }
  }

  return f;
}

vector<vector<bool>> CalcPalindromeFromBegin(string& s) {
  int n = s.size();
  vector<vector<bool>> f(n, vector<bool>(n));

  for (int i = 0; i < n; i++) {  // 如果名字取的好，可以 i->r, j->l
    for (int j = i; j >= 0; j--) {
      if (i == j) f[i][j] = true;
      else {
        if (s[i] == s[j]) {
          if (i - j == 1 || f[j+1][i-1]) f[i][j] = true;
        }
      }
    }
  }

  return f;
}

vector<vector<int>> g_f;
int CalcPalindromeWithDFS(string& s, int i, int j) {
  if (g_f[i][j] != -1) return g_f[i][j];

  if (i >= j) return  g_f[i][j] = 1;

  return g_f[i][j] = s[i] == s[j] && CalcPalindromeWithDFS(s, i + 1, j - 1) == 1;
}