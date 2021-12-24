//
// Created by 裴星鑫 on 2021/12/23.
//
#include <vector>
#include <string>
#include <unordered_set>

using namespace std;

class Solution {
 public:
  string longestDupSubstring(string s) {
    int n = s.length();
    calcRabinKarp(s);

    string ans = "";
    int l = 0, r = n;
    while (l < r) {
      int mid = l + (r - l + 1) / 2;
      string t = getSubStr(s, mid);
      if (t.empty()) r = mid - 1;
      else l = mid;
      if (t.size() > ans.size()) ans = t;
    }
    return ans;
  }

  void calcRabinKarp(string& s) {
    base_.resize(s.size() + 1);
    hash_data_.resize(s.size() + 1);
    const int P = 1313131;
    base_[0] = 1;
    for (int i = 0; i < s.size(); i++) {
        base_[i + 1] = base_[i] * P;
        hash_data_[i + 1] = hash_data_[i] * P + s[i] - 'a' + 1;
    }
  }

  string getSubStr(string& s, int len) {
    int n = s.size();
    unordered_set<int64_t> cache;
    for (int i = 1; i + len - 1 <= n; i++) {
      int tail = i + len - 1;
      int64_t cur = hash_data_[tail] - hash_data_[i - 1] * base_[len];
      if (cache.count(cur)) return s.substr(i - 1, len);
      cache.emplace(cur);
    }
    return "";
  }

 private:
  vector<int64_t> base_;
  vector<int64_t> hash_data_;
};