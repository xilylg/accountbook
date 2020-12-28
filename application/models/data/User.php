<?php
class DataUserModel extends AbstractData
{
	protected $table = 'user';
	protected $db;
	protected $log;
	
	public function getUserByUid($uid)
	{
		return $this->db->table($this->table)->where('uid', '=', $uid)->fetchOne();
	}
}