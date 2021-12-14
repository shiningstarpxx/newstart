//
// Created by 裴星鑫 on 2021/12/14.
//

#include <queue>
#include <vector>

using namespace std;

class Solution {
 public:
  int scheduleCourse(vector<vector<int>>& courses) {
    sort(courses.begin(), courses.end(), [](vector<int>& c1, vector<int>& c2) -> bool {
      return c1[1] < c2[1];
    });

    priority_queue<int, vector<int>> q;
    int total = 0;
    for (auto& v : courses) {
      int duration = v[0];
      int deadline = v[1];
      total += duration;
      q.push(duration);
      if (total > deadline) {
        total -= q.top();
        q.pop();
      }
    }
    return q.size();
  }
};