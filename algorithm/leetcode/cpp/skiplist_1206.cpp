//
// Created by 裴星鑫 on 2022/8/6.
//
#include <vector>
#include <random>
using namespace std;

class Skiplist {
 public:
  constexpr static int MAX_LEVEl = 10;
  constexpr static double P_FACTOR = 0.25;

  struct SkipListNode {
    int val;
    vector<SkipListNode*> forward;
    SkipListNode(int val, int level = MAX_LEVEl) : val(val), forward(level, nullptr) {}
  };

  SkipListNode* head;
  int level;
  mt19937 gen{random_device{}()};
  uniform_real_distribution<double> dis;

  Skiplist() : head(new SkipListNode(-1)), level(0), dis(0, 1) {}

  bool search(int target) {
    auto current = head;
    for (int i = level - 1; i >= 0; i--) {
      while (current->forward[i] && current->forward[i]->val < target)
        current = current->forward[i];
    }
    current = current->forward[0];
    while (current && current->val == target) return true;
    return false;
  }

  void add(int num) {
    vector<SkipListNode*> update(MAX_LEVEl, head);
    auto cur = head;
    for (int i = level - 1; i >= 0; i--) {
      while (cur->forward[i] && cur->forward[i]->val < num)
        cur = cur->forward[i];
      update[i] = cur;
    }
    int lv = randomLevel();
    level = max(level, lv);
    auto n = new SkipListNode(num, lv);
    for (int i = 0; i < lv; i++) {
      n->forward[i] = update[i]->forward[i];
      update[i]->forward[i] = n;
    }
  }

  bool erase(int num) {
    vector<SkipListNode*> update(MAX_LEVEl, head);
    auto cur = head;
    for (int i = level - 1; i >= 0; i--) {
      while (cur->forward[i] && cur->forward[i]->val < num)
        cur = cur->forward[i];
      update[i] = cur;
    }
    cur = cur->forward[0];
    if (!cur || cur->val != num) return false;
    for (int i = 0; i < level; i++) {
      if (update[i]->forward[i] != cur) {
        break;
      }
      update[i]->forward[i] = cur->forward[i];
    }
    delete cur;
    while (level > 1 && head->forward[level - 1] == nullptr) level--;
    return true;
  }

  int randomLevel() {
    int lv = 1;
    /* 随机生成 lv */
    while (dis(gen) < P_FACTOR && lv < MAX_LEVEl) {
      lv++;
    }
    return lv;
  }
};

/**
 * Your Skiplist object will be instantiated and called as such:
 * Skiplist* obj = new Skiplist();
 * bool param_1 = obj->search(target);
 * obj->add(num);
 * bool param_3 = obj->erase(num);
 */