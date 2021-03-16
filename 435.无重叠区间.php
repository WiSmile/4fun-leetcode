/*
 * @lc app=leetcode.cn id=435 lang=php
 *
 * [435] 无重叠区间
 */
<?php
// @lc code=start
class Solution {

    /**
     * @param Integer[][] $intervals
     * @return Integer
     */
    function eraseOverlapIntervals($intervals) {
        $endArr = [];
        foreach ($intervals as $section) {
            $endArr[] = $section[1];
        }
        array_multisort($endArr, SORT_ASC, $intervals);
        
        $end = 0;
        $cnt = 0;
        foreach ($intervals as $key => $section) {
            if ($key == 0) {
                $end = $section[1];
                continue;
            }

            if ($end > $section[0]) {
                $cnt++;
            } else {
                $end = $section[1];
            }
        }

        return $cnt;
    }
}
// @lc code=end

