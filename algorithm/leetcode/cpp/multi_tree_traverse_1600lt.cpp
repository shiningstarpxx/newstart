//
// Created by 裴星鑫 on 2021/11/21.
//

#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

class ThroneInheritance {
  unordered_map<string, vector<string>> edge_;
  unordered_map<string, int> dead_;
  string king_;
 public:
  ThroneInheritance(string kingName) {
    king_ = kingName;
  }

  void birth(string parentName, string childName) {
    edge_[parentName].push_back(childName);
  }

  void death(string name) {
    dead_[name] = 1;
  }

  vector<string> getInheritanceOrder() {
    vector<string> res;
    function<void(string & )> preOrder = [&](string &king) {
      if (!dead_.count(king)) res.push_back(king);
      for (auto it = edge_[king].begin(); it != edge_[king].end(); it++) {
        preOrder(*it);
      }
    };
    preOrder(king_);
    return res;
  }
};

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * ThroneInheritance* obj = new ThroneInheritance(kingName);
 * obj->birth(parentName,childName);
 * obj->death(name);
 * vector<string> param_3 = obj->getInheritanceOrder();
 */