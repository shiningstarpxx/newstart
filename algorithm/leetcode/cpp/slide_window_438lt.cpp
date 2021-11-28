//
// Created by 裴星鑫 on 2021/11/28.
//

#include <string>
#include <vector>

using namespace std;

// 这个问题的解法，很自然的是滑动窗口；
// 1。 一开始我使用了 hash表来存储对应的值，没有想好怎么比较两个hash表，所以用了一个total来表示是否是合理的比较，但是带来的问题是代码非常难写
//     可以看到438lt_wa的代码，在某些case下还没有调试通过
// 2. 这里的解法就是hash不好比较，就用数组比较，也就是内存直接比较。显然在只有小写的情况下，这样好写多了。
// 总结： 1的代码更通用，但是没有调试过；2的代码更特定，但是效率确实高
class Solution {
 public:
  vector<int> findAnagrams(string s, string p) {
    if (s.size() < p.size()) return {};

    vector<int> origin(26, 0);
    vector<int> target(26, 0);
    for (int i = 0; i < p.size(); i++) {
      origin[s[i] - 'a']++;
      target[p[i] - 'a']++;
    }

    vector<int> res;
    for (int i = 0; i < s.size() - p.size() + 1; i++) {
      if (origin == target) res.push_back(i);
      origin[s[i] - 'a']--;
      if (i + p.size() < s.size()) origin[s[i + p.size()] - 'a']++;
    }
    return res;
  }
};