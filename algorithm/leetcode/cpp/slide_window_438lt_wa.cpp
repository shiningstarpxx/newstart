//
// Created by 裴星鑫 on 2021/11/28.
//

class Solution {
 public:
  vector<int> findAnagrams(string s, string p) {
    if (s.size() < p.size()) return {};

    int total = p.size();
    unordered_map<char, int> target_cache;
    for (auto c : p) target_cache[c]++;

    unordered_map<char, int> source_cache;
    vector<int> res;
    for (int l = 1 - p.size(), r = 0; r < s.size(); l++, r++) {
      if (target_cache.count(s[r])) {
        source_cache[s[r]] < target_cache[s[r]]) total--;
        source_cache[s[r]]++;
      }
      // print(source_cache);
      // cout << l << ", " << r << ", " << total << endl;
      if (total == 0) res.push_back(l);
      if (l >= 0 && source_cache[s[l]] > 0) {
        if (source_cache[s[r]] <= target_cache[s[r]]) total++;
        source_cache[s[l]]--;
      }
      // print(source_cache);
      // cout << "********" << endl;
    }
    return res;
  }

  void print(unordered_map<char, int>& v) {
    for (auto it = v.begin(); it != v.end(); it++) cout << it->first << ", " << it->second << endl;
  }
};

