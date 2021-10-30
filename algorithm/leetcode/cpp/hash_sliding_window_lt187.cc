
class Solution {
 public:
  vector<string> findRepeatedDnaSequences(string s) {
    vector<string> ans;
    const int L = 10;
    if (s.length() < L) return ans;
    unordered_map<string, int> cache;

    for (int i = 0; i + L - 1 < s.length(); i++) {
      auto t = s.substr(i, L);
      cache[t]++;
      if (cache[t] == 2) ans.push_back(t);
    }
    return ans;
  }
};