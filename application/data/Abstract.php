<?php
class AbstractData
{
	protected $table;
	protected $db;
	protected $log;
	
	public function __construct($log) {
		$this->log = $log;
		$connectParams = Helper_Conf::get("application.PdoMysql");
		$this->db = Database_Factory::instanct( Helper_Conf::get("database.type"), $connectParams, $log,  Helper_Conf::get("database.table_pre"));
	}
}