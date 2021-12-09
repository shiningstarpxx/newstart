//
// Created by 裴星鑫 on 2021/12/9.
//

#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  bool validTicTacToe(vector<string> &board) {
    int x_count = 0;
    int o_count = 0;

    for (int i = 0; i < board.size(); i++) {
      for (auto c : board[i]) {
        if (c == 'X') x_count++;
        if (c == 'O') o_count++;
      }
    }

    if (x_count != o_count && x_count != o_count + 1) return false;
    if (winCheck(board, 'X') && x_count != o_count + 1) return false;
    if (winCheck(board, 'O') && x_count != o_count) return false;
    return true;
  }

  bool winCheck(vector<string> &board, char target) {
    for (int i = 0; i < 3; i++) {
      if (board[i][0] == target && board[i][1] == target && board[i][2] == target) return true;
      if (board[0][i] == target && board[1][i] == target && board[2][i] == target) return true;
    }
    if (board[0][0] == target && board[1][1] == target && board[2][2] == target) return true;
    if (board[2][0] == target && board[1][1] == target && board[0][2] == target) return true;
    return false;
  }
};