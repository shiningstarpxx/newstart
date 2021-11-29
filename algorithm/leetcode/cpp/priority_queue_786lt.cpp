//
// Created by 裴星鑫 on 2021/11/29.
//
#include <queue>
#include <vector>

using namespace std;

class Solution {
 public:
  // key point: a / b < c / d == a * d < b * c
  vector<int> kthSmallestPrimeFraction(vector<int>& arr, int k) {
    // return bruteForce(arr, k);
    return maxHeap(arr, k);
  }

  vector<int> maxHeap(vector<int>& arr, int k) {
    auto cmp = [&](const pair<int, int>& first, const pair<int, int>& second) {
      return arr[first.first] * arr[second.second] > arr[first.second] * arr[second.first];
    };
    priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> q(cmp);

    for (int j = 1; j < arr.size(); ++j) {
      q.emplace(0, j);
    }
    for (int _ = 1; _ < k; ++_) {
      auto [i, j] = q.top();
      q.pop();
      if (i + 1 < j) {
        q.emplace(i + 1, j);
      }
    }
    return {arr[q.top().first], arr[q.top().second]};
  }

  vector<int> bruteForce(vector<int>& arr, int k) {
    vector<pair<int, int>> data;
    for (int i = 0; i < arr.size(); i++) {
      for (int j = i + 1; j < arr.size(); j++)
        data.push_back({arr[i], arr[j]});
    }

    auto cmp = [](const pair<int, int>& first, const pair<int, int>& second) {
      return first.first * second.second < first.second * second.first;
    };

    sort(data.begin(), data.end(), cmp);

    return {data[k-1].first, data[k-1].second};
  }
};