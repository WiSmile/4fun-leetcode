/*
 * @lc app=leetcode.cn id=763 lang=php
 *
 * [763] 划分字母区间
 */
<?php
// @lc code=start
class Solution {

    /**
     * @param String $S
     * @return Integer[]
     */
    function partitionLabels($S) {
        $location = [];
        for ($i = 0; $i < strlen($S); $i++) {
            $location[$S[$i]][] = $i;
        }

        $result = [0];
        $end = 0;
        foreach ($location as $v) {
            if ($end <= $v[0]) {
                $end = end($v);
                $result[] = $end - end($result) + 1;
            } else {
                if (end($v) > $end) $end = end($v);
            }
        }

        $cnt = count($result);
        if ($end > $result[$cnt-1] + $result[$cnt-2] - 1) {
            array_pop($result);
            $result[] = $end - end($result) + 1;
        }
        array_shift($result);

        return $result;
    }
}
// @lc code=end

