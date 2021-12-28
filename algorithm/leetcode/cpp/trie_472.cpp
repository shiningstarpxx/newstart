//
// Created by 裴星鑫 on 2021/12/28.
//
#include <string>
#include <vector>

using namespace std;

const int kASICCNUM = 26;

class Trie {
  bool is_end_;
  vector<Trie *> children_;
 public:
  Trie() {
    this->children_ = vector<Trie *>(kASICCNUM, nullptr);
    this->is_end_ = false;
  }

  void insert(const string & word) {
    Trie* node = this;
    for (int i = 0; i < word.size(); i++) {
      char ch = word[i];
      int index = ch - 'a';
      if (node->children_[index] == nullptr) {
        node->children_[index] = new Trie();
      }
      node = node->children_[index];
    }
    node->is_end_ = true;
  }

  bool dfs(const string& word, int start) {
    if (word.size() == start) return true;

    auto t = this;
    for (int i = start; i < word.size(); i++) {
      char c = word[i];
      int index = word[i] - 'a';
      t = t->children_[index];
      if (t == nullptr) return false;
      if (t->is_end_) {
        if (dfs(word, i + 1)) return true;
      }
    }
    return false;
  }
};

class Solution {
 public:
  vector<string> findAllConcatenatedWordsInADict(vector<string>& words) {
    Trie tree;
    sort(words.begin(), words.end(), [](string& a, string& b) ->bool {
      return a.size() < b.size();
    });

    vector<string> res;
    for (int i = 0; i < words.size(); i++) {
      string& s = words[i];
      if (s.size() == 0) continue;

      if (tree.dfs(s, 0)) {
        res.push_back(s);
      } else {
        tree.insert(s);
      }
    }
    return res;
  }
};
