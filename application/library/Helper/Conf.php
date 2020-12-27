<?php
/**
 * 配置文件存取器  (封装Yaf_Config_Ini, 便于使用)
 *
 * @package   Helper
 */
class Helper_Conf {
	
	private static $inst = array();
	
	private function __construct(){
		
	}
	
	public static function get($key){
		if (strpos($key, '.') !== false) {
			list($file, $path) = explode('.', $key, 2);
		}else{
			$file = $key;
		}
		
		if (!isset(self::$inst[$file])) {
			self::$inst[$file] = new Yaf_Config_Ini(APPPATH . '/config/' . $file . '.ini');
		}
		
		if (isset($path)) {
			$ret = self::$inst[$file]->get($path);
			if (is_a($ret,'Yaf_Config_Ini')){
				return $ret->toArray();
			}else{
				return $ret;
			}
		}else{
			return self::$inst[$file]->toArray();
		}
	}
	
}