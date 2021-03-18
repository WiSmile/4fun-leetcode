/*
 * @lc app=leetcode.cn id=605 lang=php
 *
 * [605] 种花问题
 */
<?php 
// @lc code=start
class Solution {

    /**
     * @param Integer[] $flowerbed
     * @param Integer $n
     * @return Boolean
     */
    function canPlaceFlowers($flowerbed, $n) {
        $len = count($flowerbed);
        $cnt = 0;

        if ($flowerbed[0] == 0 && $flowerbed[1] == 0) {
            $flowerbed[0] = 1;
            $cnt++;
        }

        for ($i = 1; $i < $len - 1; $i++) {
            if ($flowerbed[$i-1] == 0 && $flowerbed[$i] == 0 && $flowerbed[$i+1] == 0) {
                $flowerbed[$i] = 1;
                $cnt++;
            }
        }

        if ($flowerbed[$len-1] == 0 && $flowerbed[$len-2] == 0) {
            $flowerbed[$len-1] == 0;
            $cnt++;
        }

        return $cnt >= $n;
    }
}
// @lc code=end

