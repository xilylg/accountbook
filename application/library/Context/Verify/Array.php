<?php
class Context_Verify_Array
{
	private $value;
	private $success = ['code' => 1];
	private $error = ['coce' => 0, 'error_msg' => ''];
	
	public function __construct($key, $value)
	{
		if(!is_array($value)) {
			$this->error['error_msg'] = "{$this->key} is not array";
			return $this->error;
		}
		
		$this->value = $value;
		$this->key = $key;
	}
	
	public function required()
	{
		if (empty($this->value)) {
			$this->error['error_msg'] = "{$this->key} is empty";
			return $this->error;
		}
		return $this->success;
	}
}