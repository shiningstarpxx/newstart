//
// Created by 裴星鑫 on 2021/12/18.
//

#ifndef LEETCODE__DFS_FLOOD_FILL_419_H_
#define LEETCODE__DFS_FLOOD_FILL_419_H_

#include <vector>

using namespace std;

class dfs_flood_fill_419 {
   public:
    int countBattleships(vector<vector<char>>& board) {
      int res = 0;
      for (int i = 0; i < board.size(); i++) {
        for (int j = 0; j < board[i].size(); j++) {
          if (board[i][j] == 'X') {
            dfs(board, i, j);
            res++;
          }
        }
      }
      return res;
    }
    void dfs(vector<vector<char>>& board, int x, int y) {
      if (x < 0 || x >= board.size() || y < 0 || y >= board[x].size() || board[x][y] != 'X') return;
      board[x][y] = '.';
      dfs(board, x + 1, y);
      dfs(board, x - 1, y);
      dfs(board, x, y + 1);
      dfs(board, x, y - 1);
    }
};

#endif //LEETCODE__DFS_FLOOD_FILL_419_H_
