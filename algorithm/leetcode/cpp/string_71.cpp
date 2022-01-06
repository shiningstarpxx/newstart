//
// Created by 裴星鑫 on 2022/1/6.
//
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  string simplifyPath(string path) {
    auto split = [](string& path, char delim) -> vector<string> {
      vector<string> res;
      string cur;
      for (auto c : path) {
        if (c == delim) {
          res.push_back(cur);
          cur.clear();
        } else {
          cur += c;
        }
      }
      res.push_back(cur);
      return res;
    };

    vector<string> r = split(path, '/');
    vector<string> filter;
    for (auto s : r) {
      if (s == "..") {
        if (!filter.empty()) filter.pop_back();
      } else if (!s.empty() && s != ".") {
        filter.push_back(s);
      }
    }

    string ans;
    if (filter.empty()) ans = "/";
    else {
      for (auto s : filter) ans += "/" + s;
    }
    return ans;
  }
};