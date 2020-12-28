<?php
class UserModel 
{
	private $userData;
	private $log;
	
	public function __construct($log = NULL) {
		$this->log = is_object($log) ? $log : new Log_Log(LOG_PATH);
		$this->userData = new UserData($this->log);
	}
	
	public function getUserinfoByUid($uid)
	{
		return $this->userData->getUserByUid($uid);
	}
}