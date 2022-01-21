//
// Created by 裴星鑫 on 2021/12/22.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  int repeatedStringMatch(string a, string b) {
    int n = a.size();
    int m = b.size();
    int r = kmp(a, b);
    if (r == -1) return -1;

    if (n - r >= m) return 1;
    return (m + r - n - 1) / n + 2;
  }

  int kmp(string& haystack, string& needle) {
    int n = needle.size();
    int m = haystack.size();
    if (n == 0) return 0;

    vector<int> p(n);
    for (int i = 1, j = 0; i < n; i++) {
      while (j > 0 && needle[i] != needle[j]) {
        j = p[j - 1];
      }
      if (needle[i] == needle[j]) {
        j++;
      }
      p[i] = j;
    }

    for (int i = 0, j = 0; i - j < m; i++) {
      while (j > 0 && haystack[i % m] != needle[j]) {
        j = p[j - 1];
      }
      if (haystack[i % m] == needle[j]) j++;
      if (j == n) return i - n + 1;
    }

    return -1;
  }
};