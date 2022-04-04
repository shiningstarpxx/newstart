//
// Created by 裴星鑫 on 2022/3/30.
//
#include <vector>
#include <set>
#include <queue>

using namespace std;

class Solution {
 public:
  vector<int> busiestServers(int k, vector<int>& arrival, vector<int>& load) {
    set<int> available;  // for binary search
    for (int i = 0; i < k; i++) available.insert(i);

    priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> busy;  // 结束时间排序
    vector<int> request_count(k);
    for (int i = 0; i < arrival.size(); i++) {
      int current_time = arrival[i];

      // 先遍历繁忙队列，确认是否有服务需要更新
      while (!busy.empty()) {
        auto [finish, index] = busy.top();
        if (finish > current_time) break;
        available.insert(index);
        busy.pop();
      }

      // 如果没有可用的服务器放过
      if (available.empty()) continue;

      // 找到可以服务的server index
      auto it = available.lower_bound(i % k);
      if (it == available.end()) it = available.begin();

      busy.push({load[i] + current_time, *it});
      request_count[*it]++;
      available.erase(it);
    }

    vector<int> res;
    int max_count = *max_element(request_count.begin(), request_count.end());
    for (int i = 0; i < k; i++) {
      if (request_count[i] == max_count) res.push_back(i);
    }
    return res;
  }
};