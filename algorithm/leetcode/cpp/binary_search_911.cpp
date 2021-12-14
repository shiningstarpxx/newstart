//
// Created by 裴星鑫 on 2021/12/11.
//

#include <vector>
#include <unordered_map>

using namespace std;

class TopVotedCandidate {
 public:
  TopVotedCandidate(vector<int>& persons, vector<int>& times) {
    times_ = times;
    unordered_map<int, int> count;
    count[-1] = -1;
    int top = -1;
    for (auto c : persons) {
      count[c]++;
      if (count[c] >= count[top]) {
        top = c;
      }
      top_.push_back(top);
    }
  }

  int q(int t) {
    int p = upper_bound(times_.begin(), times_.end(), t) - times_.begin();
    return top_[p - 1];
  }
 private:
  vector<int> top_;
  vector<int> times_;
};

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * TopVotedCandidate* obj = new TopVotedCandidate(persons, times);
 * int param_1 = obj->q(t);
 */