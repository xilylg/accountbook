<?php
class Context_Verify_Verify
{
	private $error = ['code' => 0, 'error_msg' => 'invalid rules'];
	
	/**
	 * 
	 * @param string $key
	 * @param string $value
	 * @param string $rules $value_type|another_rule|another_rule for example: string|required|min_length:1|max_length:2|min:10|max:100|
	 */
	public static function request($key, $value, $rules)
	{
		$rules = explode("|", $rules);
		if (empty($rules)) {
			return self::$error;
		}
		
		$class_name = $rules[0];
		
		if (!class_exists("Context_Verify_{$class_name}")) {
			return self::$error;
		}
		
		unshift($rules);

		$class = call_user_class("Context_Verify_{$class_name}", $key, $value);
		if (is_array($class)) {
			return $class;
		}
		foreach($rules as $rule) {
			list($method, $param) = explode(":", $rule);
			if (!method_exists($class, $method)) {
				self::$error['error_msg'] = "{$class_name}->{$rule}";
				return self::$error;
			}
			$result = $class->$method($param);
			if ($result['code'] == 0) {
				return $result;
			}
		}
		
		return self::$success;
	}
}