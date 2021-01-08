<?php
class DataUserModel extends AbstractData
{
	protected $table = 'user';
	protected $fields = ["uid",  "nickname",  "password",  "gender",  "birthday",  "work",  "status",  "create_time",  "update_time"];
	protected $db;
	protected $log;
	
	public function getUserByUid($uid)
	{
		return $this->db->table($this->table)->where('uid', '=', $uid)->fetchOne();
	}
}