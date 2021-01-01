<?php
class Log_Log
{
	const WARNING = "warning";
	const ERROR = "error";
	const DEBUG = "debug";
	const TRACE = "trace";
	protected $separator;
	protected $path = "";
	
	public function __construct($path, $separator = "|", $level = '') 
	{
		$this->separator = $separator;
		$this->path = $path;
	}
	
	public function waring($module, Array $msg) {
		$this->write("waring_{$module}", $msg);
	}
	
	public function error($module, Array $msg) {
		$this->write("error_{$module}", $msg);
	}
	
	public function trace($module, Array $msg) {
		$this->write("trace_{$module}", $msg);
	}
	
	public function write($module, Array $msg) {
		$str = date("Y-m-d H:i:s");
		$str .= implode($this->separator, $msg);
		$file = $this->path."/".$module ."_".date("Y-m-d");
		file_put_content($file, $str."\n", FILE_APPEND);
	}
}