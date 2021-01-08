<?php
class AbstractData
{
	protected $table;
	protected $db;
	protected $log;
	protected $primaryKey; //表主键名
	protected $fields = [];
	
	public function __construct($log) {
		$this->log = $log;
		$this->table = Helper_Conf::get("database.table_pre") . "_" . $this->table;
		$connectParams = Helper_Conf::get("application.PdoMysql");
		$this->db = Database_Factory::instanct( Helper_Conf::get("database.type"), $connectParams, $log,  Helper_Conf::get("database.table_pre"));
	}
	
	public function setTable($table) {
	    $this->table = Helper_Conf::get("database.table_pre") . "_" . $table;
	}
	
	public function setPrimaryKey($primaryKey) {
	    $this->primaryKey = $primaryKey;
	}
	
	/**
	 * 
	 * @param array $data 数据库$field_name => $field_value
	 * @return boolean
	 */
	private function _addOne(Array $data) {
	    return $this->db->table($this->table)->insert($data);
	}
	
	private function _updateOne(Array $data) {
	    if (empty($data[$this->primaryKey])) {
	        return ['code' => 0, 'err_msg' => "invalid {$this->primaryKey} value"];
	    }
	    return $this->db->table($this->table)->where($this->primaryKey, "=", $data[$this->primaryKey])->update($data);
	}
	
	private function _deleteOne(Array $data) {
	    if (empty($data[$this->primaryKey])) {
	        return ['code' => 0, 'err_msg' => "invalid {$this->primaryKey} value"];
	    }
	    return $this->db->table($this->table)->where($this->primaryKey, "=", $data[$this->primaryKey])->delete();
	}
	
	private function _findOne(Array $data) {
	    $conn = $this->db->table($this->table);
	    $rel = "";
	    foreach($data as $k => $v) {
	        $conn->where($k, '=', $v, $rel);
	        $rel = 'and';
	    }
	    return $conn->findOne();
	}
	
	private function _find(Array $data) {
	    $conn = $this->db->table($this->table);
	    $rel = "";
	    foreach($data as $k => $v) {
	        $conn->where($k, '=', $v, $rel);
	        $rel = 'and';
	    }
	    return $conn->query();
	}
	
	public function __call($method, $params) {
	    if (method_exists($this, "_".$method)) {
	        $data_keys = array_keys($params);
	        if (!Helper_Array::in_array($data_keys, $this->fields)) {
	            throw new Exception("invalid fields");
	        }
	        return call_user_func_array([$this, "_".$method], $params);
	    } else {
	        throw new Exception("invalid method");
	    }
	}
}