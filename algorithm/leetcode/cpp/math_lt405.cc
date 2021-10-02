

class Solution {
public:
    string toHex(int num) {
        // return calc(num);
        return bitCalc(num);
    }

   string calc(int num) {
       if (num == 0) {
           return "0";
       }
       long d = num;
       if (num < 0) {
           d += pow(2, 32);
       }

       string res;
       while (d) {
           int t = d % 16;
           char c;
           if (t < 10) c = ('0'+t);
           else c = ('a' + (t - 10));
           res += c;
           d /= 16;
       }
       reverse(res.begin(), res.end());
       return res;
   }

   string bitCalc(int num) {
       if (num == 0) {
           return "0";
       }

       string res;
       for (int i = 7; i >= 0; i--) {
           int b = (num >> (i * 4)) & 0xf;
           if (b > 0 || res.length() > 0) {
               if (b <= 9) res = res + char('0' + b);
               else res = res + char('a' + b - 10);
           }
       }
       return res;
   }
};