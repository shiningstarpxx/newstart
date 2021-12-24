//
// Created by 裴星鑫 on 2021/12/24.
//
#include <queue>
#include <vector>

using namespace std;

class Solution {
 public:
  int eatenApples(vector<int>& apples, vector<int>& days) {
    priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> q;
    int n = apples.size();
    int time = 0;
    int ans = 0;
    while (time < n || !q.empty()) {
      if (time < n && apples[time] > 0) q.push({time + days[time] - 1, apples[time]});
      while (!q.empty() && q.top().first < time) q.pop();
      if (!q.empty()) {
        auto p = q.top();
        q.pop();
        if (--p.second > 0) q.push(p);
        ans++;
      }
      time++;
    }

    return ans;
  }
};
