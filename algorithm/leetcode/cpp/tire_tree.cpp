//
// Created by 裴星鑫 on 2022/3/18.
//
#include <string>
#include <vector>

using namespace std;

class Trie {
 public:
  /** Initialize your data structure here. */
  Trie() {
    children_ = vector<Trie*> (26, nullptr);
    is_end_ = false;
  }

  /** Inserts a word into the trie. */
  void insert(string word) {
    auto t = this;
    for (auto c : word) {
      if (t->children_[c - 'a'] == nullptr) t->children_[c - 'a'] = new Trie;
      t = t->children_[c - 'a'];
    }
    t->is_end_ = true;
  }

  /** Returns if the word is in the trie. */
  bool search(string word) {
    auto t = this;
    for (auto c : word) {
      if (t->children_[c - 'a'] == nullptr) return false;
      t = t->children_[c - 'a'];
    }
    return t->is_end_;
  }

  /** Returns if there is any word in the trie that starts with the given prefix. */
  bool startsWith(string prefix) {
    auto t = this;
    for (auto c : prefix) {
      if (t->children_[c - 'a'] == nullptr) return false;
      t = t->children_[c - 'a'];
    }
    return true;
  }

  vector<Trie*> children_;
  bool is_end_;
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */