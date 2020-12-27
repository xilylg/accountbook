<?php
class Context_Verify_String
{
	private $value;
	private $success = ['code' => 1];
	private $error = ['coce' => 0, 'error_msg' => ''];
	
	public function __construct($key, $value) 
	{
		if(!is_string($value)) {
			$this->error['error_msg'] = "{$this->key} is not a string";
			return $this->error;
		}
		
		$this->value = $value;
		$this->key = $key;
	}
	
	public function required()
	{
		if (mb_strlen($this->value) == 0) {
			$this->error['error_msg'] = "{$this->key} is empty";
			return $this->error;
		}
		return $this->success;
	}
	
	public function minLength($min)
	{
		if (mb_strlen($this->value) < $min) {
			$this->error['error_msg'] = "{$this->key}'s length is less than {$min}";
			return $this->error;
		}
		return $this->success;
	}
	
	public function maxLenght($max) 
	{
		if (mb_strlen($this->value) > $max) {
			$this->error['error_msg'] = "{$this->key}'s length is more than {$max}}";
			return $this->error;
		}
		return $this->success;
	}
}