/*
 * @lc app=leetcode.cn id=135 lang=php
 *
 * [135] 分发糖果
 */
<?php
// @lc code=start
class Solution {

    /**
     * @param Integer[] $ratings
     * @return Integer
     */
    function candy($ratings) {
        $length = count($ratings);
        $candyCnt = 0;
        $candyArr = [];

        foreach ($ratings as $rating) {
            $candyArr[] = 1;
        }

        for ($i = 1; $i < $length; $i++) {
            if ($ratings[$i] > $ratings[$i-1]) {
                $candyArr[$i] = $candyArr[$i-1] + 1;
            }
        }

        for ($j = $length - 2; $j >= 0; $j--) {
            if ($ratings[$j] > $ratings[$j+1] && $candyArr[$j] <= $candyArr[$j+1]) {
                $candyArr[$j] = $candyArr[$j+1] + 1;
            }
        }

        foreach ($candyArr as $value) {
            $candyCnt += $value;
        }

        return $candyCnt;
    }
}
// @lc code=end

