/*
 * @lc app=leetcode.cn id=455 lang=php
 *
 * [455] 分发饼干
 */
<?php
// @lc code=start
class Solution {

    /**
     * @param Integer[] $g
     * @param Integer[] $s
     * @return Integer
     */
    function findContentChildren($g, $s) {
        sort($s);
        sort($g);

        $child = 0;
        $cookie = 0;
        $total_g = count($g);
        $total_s = count($s);
        while ($child < $total_g && $cookie < $total_s) {
            if ($s[$cookie] >= $g[$child]) {
                $child++;
            }

            $cookie++;
        }

        return $child;
    }
}
// @lc code=end

