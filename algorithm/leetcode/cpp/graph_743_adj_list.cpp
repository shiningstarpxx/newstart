//
// Created by 裴星鑫 on 2022/1/27.
//
#include <stdio.h>
#include <functional>
#include <queue>
#include <vector>

#include "gtest/gtest.h"

using namespace std;

/*
 int[] he = new int[N], e = new int[M], ne = new int[M], w = new int[M];
int idx;

void add(int a, int b, int c) {
    e[idx] = b;            // 先把当前边的指向节点赋值
    ne[idx] = he[a];       // 保存当前的头节点
    he[a] = idx;           // 更改当前的头节点
    w[idx] = c;            // 保存当前边的权重
    idx++;                 // 边索引增加
}

首先 idx 是用来对边进行编号的，然后对存图用到的几个数组作简单解释：

he 数组：存储是某个节点所对应的边的集合（链表）的头结点；
e  数组：由于访问某一条边指向的节点；
ne 数组：由于是以链表的形式进行存边，该数组就是用于找到下一条边；
w  数组：用于记录某条边的权重为多少。
因此当我们想要遍历所有由 a 点发出的边时，可以使用如下方式：

for (int i = he[a]; i != -1; i = ne[i]) {
    int b = e[i], c = w[i]; // 存在由 a 指向 b 的边，权重为 c
}

*/

const static int kNodes = 110;
const static int kEdges = 6010;

class Solution {
 public:
  int networkDelayTime2(vector<vector<int>> &times, int n, int k) {
    n_ = n;
    k_ = k;
    printf("2******\n");

    generateAdj(times);

    printf("before******\n");
    // dijstra(k);

    int ans = 0;
    for (int i = 1; i <= n; i++) {
      ans = max(ans, dist_[i]);
    }
    return ans >= INT32_MAX / 2 ? -1 : ans;
  }

 private:
  void dijstra(int k) {
    auto c = [](pair<int, int>& a, pair<int, int> b) {return a.second - b.second;};
    priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(c)> q(c);
    dist_[k] = 0;
    q.push({k, 0});

    printf("*****\n");
    while (!q.empty()) {
      auto [node, dis] = q.top();
      q.pop();
      if (visit_[node]) continue;
      visit_[node] = true;
      for (int i = start_point_to_edge_[node]; i != -1; i = next_edge_[i]) {
        int end_point = edge_to_end_point_[i];
        if (dist_[end_point] > dist_[node] + weight_list_[i]) {
          dist_[end_point] = dist_[node] + weight_list_[i];
          q.push({end_point, dist_[end_point] });
        }
      }
    }
  }

  void generateAdj(vector<vector<int>>& times) {
    for (auto& v : times) {
      appendEdge(v);
    }
  }

  void appendEdge(vector<int>& edge) {
    weight_list_[edge_index_] = edge[2];
    edge_to_end_point_[edge_index_] = edge[1];
    next_edge_[edge_index_] = start_point_to_edge_[edge[0]];
    start_point_to_edge_[edge[0]] = edge_index_;
    edge_index_++;
  }

  int n_;
  int k_;
  int edge_index_ = 0;
  vector<int> weight_list_ = vector<int>(kEdges);
  vector<int> edge_to_end_point_ = vector<int>(kEdges);
  vector<int> next_edge_ = vector<int>(kEdges);
  vector<int> start_point_to_edge_ = vector<int>(kNodes, -1);

  vector<int> dist_ = vector<int>(kNodes, INT32_MAX / 2);
  vector<bool> visit_ = vector<bool>(kNodes, false);
};

TEST(Solution, adj) {
  Solution s;

  vector<vector<int>> times = {{2, 1, 1}, {2, 3, 1}, {3, 4, 1}};

  int r = s.networkDelayTime2(times, 4, 2);
  ASSERT_EQ(1, 2);

  ASSERT_EQ(2, r);
  ASSERT_EQ(3, r);
}