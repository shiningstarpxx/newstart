//
// Created by 裴星鑫 on 2022/1/4.
//
#include <vector>

using namespace std;

class Solution {
 public:
  int catMouseGame(vector<vector<int>>& graph) {
    graph_ = graph;
    n_ = graph_.size();
    dp_ = vector<vector<vector<int>>>(2 * n_ + 1, vector<vector<int>>(n_ + 1, vector<int>(n_ + 1, -1)));
    return dfs(0, 1, 2);
  }

  int dfs(int step, int mouse, int cat) {
    int ans = dp_[step][mouse][cat];
    if (mouse == 0) ans = 1;
    else if (cat == mouse) ans = 2;
    else if (step >= 2 * n_) ans = 0;
    if (ans == -1) {
      if (step % 2 == 0) {  // mouse
        bool win = false;
        bool draw = false;
        for (auto d : graph_[mouse]) {
          int t = dfs(step + 1, d, cat);
          if (t == 1) win = true;
          else if (t == 0) draw = true;
          if (win) break;
        }
        if (win) ans = 1;
        else if (draw) ans = 0;
        else ans = 2;
      } else {  // cat
        bool win = false;
        bool draw = false;
        for (auto d : graph_[cat]) {
          if (d == 0) continue;
          int t = dfs(step + 1, mouse, d);
          if (t == 2) win = true;
          else if (t == 0) draw = true;
          if (win) break;
        }
        if (win) ans = 2;
        else if (draw) ans = 0;
        else ans = 1;
      }
    }
    dp_[step][mouse][cat] = ans;
    return ans;
  }

 private:
  vector<vector<int>> graph_;
  vector<vector<vector<int>>> dp_;
  int n_;
};