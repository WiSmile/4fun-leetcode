/*
 * @lc app=leetcode.cn id=452 lang=php
 *
 * [452] 用最少数量的箭引爆气球
 */
<?php
// @lc code=start
class Solution {

    /**
     * @param Integer[][] $points
     * @return Integer
     */
    function findMinArrowShots($points) {
        if (empty($points)) {
            return 0;
        }

        $arrStart = [];
        foreach ($points as $point) {
            $arrStart[] = $point[0];
        }

        array_multisort($arrStart, SORT_ASC, $points);

        $cnt = 1;
        $end = $points[0][1];
        for ($i = 1; $i < count($points); $i++) {
            if ($points[$i][0] > $end) {
                $end = $points[$i][1];
                $cnt++;
            } else {
                if ($points[$i][1] < $end) {
                    $end = $points[$i][1];
                }
            }
        }

        return $cnt;
    }
}
// @lc code=end

