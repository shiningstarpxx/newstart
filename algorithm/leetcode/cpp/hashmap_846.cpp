//
// Created by 裴星鑫 on 2021/12/30.
//
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
 public:
  bool isNStraightHand(vector<int>& hand, int groupSize) {
    if (hand.size() % groupSize != 0) return false;
    sort(hand.begin(), hand.end());
    unordered_map<int, int> cache;
    for (auto d : hand) cache[d]++;

    for (auto d : hand) {
      if (cache[d] == 0) continue;
      for (int i = 0; i < groupSize; i++) {
        if (cache[d + i] <= 0) return false;
        cache[d + i]--;
      }
    }

    return true;
  }
};

