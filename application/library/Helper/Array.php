<?php
class Helper_Array 
{
    public static function in_array(array $needle, array $haystack) {
        foreach($needle as $item) {
            if (!in_array($item, $haystack)) {
                return false;
            }
        }
        return true;
    }
}