<?php
class Context_Verify_Int
{
	private $value;
	private $success = ['code' => 1];
	private $error = ['coce' => 0, 'error_msg' => ''];
	
	public function __construct($key, $value)
	{
		if(!is_number($value)) {
			$this->error['error_msg'] = "{$this->key} is not a number";
			return $this->error;
		}
		
		$this->value = $value;
		$this->key = $key;
	}
	
	public function required()
	{
		if ($this->value === NULL) {
			$this->error['error_msg'] = "{$this->key} is empty";
			return $this->error;
		}
		return $this->success;
	}
	
	public function min($min)
	{
		if ($this->value < $min) {
			$this->error['error_msg'] = "{$this->key} is less than {$min}";
			return $this->error;
		}
		return $this->success;
	}
	
	public function max($max)
	{
		if ($this->value > $max) {
			$this->error['error_msg'] = "{$this->key} is more than {$max}}";
			return $this->error;
		}
		return $this->success;
	}
}