<?php
class Context_Request
{
	public function param($key) {
		return $_REQUEST[$key];
	}
	
	public function get($key) {
		return $_GET[$key];
	}
	
	public function post($key) {
		return $_POST[$key];
	}
}