//
// Created by 裴星鑫 on 2022/7/15.
//

//设计一个包含一些单词的特殊词典，并能够通过前缀和后缀来检索单词。
//
// 实现 WordFilter 类：
//
//
// WordFilter(string[] words) 使用词典中的单词 words 初始化对象。
// f(string pref, string suff) 返回词典中具有前缀 prefix 和后缀 suff 的单词的下标。如果存在不止一个满足要求的下标，
//返回其中 最大的下标 。如果不存在这样的单词，返回 -1 。
//
//
//
//
// 示例：
//
//
//输入
//["WordFilter", "f"]
//[[["apple"]], ["a", "e"]]
//输出
//[null, 0]
//解释
//WordFilter wordFilter = new WordFilter(["apple"]);
//wordFilter.f("a", "e"); // 返回 0 ，因为下标为 0 的单词：前缀 prefix = "a" 且 后缀 suff = "e" 。
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 10⁴
// 1 <= words[i].length <= 7
// 1 <= pref.length, suff.length <= 7
// words[i]、pref 和 suff 仅由小写英文字母组成
// 最多对函数 f 执行 10⁴ 次调用
//
// Related Topics 设计 字典树 字符串 👍 150 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class WordFilter {
 public:
  class Tire {
   public:
    Tire() {
    }

    void Insert(string& s, int ii, bool reverse) {
      auto t = this;
      t->indices.push_back(ii);
      for (int i = reverse ? s.size() - 1:0; reverse ? i >= 0 : i < s.size(); reverse ? i-- : i++) {
        int index = s[i] - 'a';
        if (t->data[index] == nullptr) t->data[index] = new Tire();
        t = t->data[index];
        t->indices.push_back(ii);
      }
    }

    vector<int> Prefix(string& s, bool reverse) {
      auto t = this;
      for (int i = reverse ? s.size() - 1 : 0; reverse ? i >= 0 : i < s.size(); reverse ? i-- : i++) {
        int index = s[i] - 'a';
        if (t->data[index] == nullptr) return {};
        t = t->data[index];
      }
      return t->indices;
    }
    Tire* data[26] = {};
    vector<int> indices;
  };

  WordFilter(vector<string>& words) {
    for (int i = 0; i < words.size(); i++) {
      tire_.Insert(words[i], i, false);
      tire_reverse_.Insert(words[i], i, true);
    }
  }

  int f(string pref, string suff) {
    auto p = tire_.Prefix(pref, false);
    if (p.empty()) return -1;
    auto s = tire_reverse_.Prefix(suff, true);
    if (s.empty()) return - 1;

    for (int i = p.size() - 1, j = s.size() - 1; i >= 0 && j >= 0; ) {
      if (p[i] > s[j]) i--;
      else if (p[i] < s[j]) j--;
      else return p[i];
    }
    return -1;
  }

  Tire tire_;
  Tire tire_reverse_;
};

/**
 * Your WordFilter object will be instantiated and called as such:
 * WordFilter* obj = new WordFilter(words);
 * int param_1 = obj->f(pref,suff);
 */
//leetcode submit region end(Prohibit modification and deletion)
