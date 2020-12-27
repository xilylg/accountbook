<?php
class Database_Mysql
{
	protected $db;
	protected $connectParams;
	protected $master = false;
	protected $op;
	protected $sql;
	protected $columns = "*";
	protected $table = "";
	protected $where = "";
	protected $order = "";
	protected $group = "";
	protected $limit = "";
	protected $insertKeys = [];
	protected $updateKeys = [];
	protected $bindParams = [];
	
	public function __construct($connect_params) {
		$this->connectParams = $connect_params;
	}
	
	public function master($master = false) {
		$this->master = $master;
	}
	
	public function fields($columns = "*") {
		$this->fields = $columns;
	}
	
	public function table($table) {
		$this->table = $table;
	}
	
	public function where($key, $seq, $value, $pre_relation = NULL) {
		
		$this->where .= "{$pre_relation} {$key} {$seq} ";
		if (in_array($seq, ['in', 'not in'])) {
			$value = is_array($value) ? $value : [$value];
			foreach($value as $item) {
				$this->bindParams[] = $item;
				$values[] = "?";
			}
			$this->where .= "(" . implode(",", $values) . ")";
		} else {
			$this->where .= "?";
			$this->bindParams[] = $value;
		}
	}
	
	public function order($key, $desc) {
		$this->order[] = "{$key} {$desc}";
	}
	
	public function group($keys) {
		if (is_array($keys)) {
			$this->group[] = array_merge($this->group, $keys);
		} else {
			$this->group[] = $keys;
		}
	}
	
	public function fetchAll()
	{
		$this->op = "select";
		$stm = $this->exec();
		return $stm->fetchAll();
	}
	
	public function fetchOne()
	{
		$this->op = "select";
		$stm = $this->exec();
		return $stm->fetchOne();
	}
	
	public function query($sql, $bind_params)
	{
		$this->op = "select";
		$db = $this->connect();
		$stm = $db->prepare($sql);
		$stm->execute($bind_params);
		return $stm->fetchAll();
	}
	
	public function update(Array $update_data) {
		$this->op = "update";
		$this->master = true;
		$this->updateKey[] = array_keys($update_data);
		$this->bindParam[] = $update_data;
		return $this->exec();
	}
	
	public function insert(Array $insert_data) {
		$this->op = "insert";
		$this->master = true;
		$this->insertKey[] = array_keys($insert_data);
		$this->bindParams[] = $insert_data;
		return $this->exec();
	}
	
	public function delete() {
		$this->op = "delete";
		$this->master = true;
		return $this->exec();
	}
	
	public function exec() {
		$this->getSql();
		$db = $this->connect();
		$stm = $db->prepare($this->sql);
	    $stm->execute($this->bindParams);
	    return $stm;
	}
	
	public function getSql() {
		$this->sql = "{$this->op} {$this->columns} ";
		switch($this->op) {
			case 'select':
				$this->sql .= " from {$this->table} where {$this->where}";
				if (!empty($this->ordersF)) {
					$this->sql .= " order by " . implode(' ', $this->order);
				}
				if (!empty($this->group)) {
					$this->sql .= " group by " . implode(' ', $this->group);
				}
				if (!empty($this->limit)) {
					$this->sql .= " limit {$this->limit}";
				}
				break;
			case 'insert':
				$placeholder = array_fill("?", length($this->insertKeys));
				$this->sql .= " into {$this->table}(" . implode(",", $this->insertKeys).") values(".implode(",", $placeholder).")";
				break;
			case 'update':
				$updateData = [];
				foreach($this->updateKey as $item) {
					$updateData[] = "{$item}=?";
				}
				$this->sql .= " {$this->table} set " . implode(",", $updateData);
				break;
			case 'delete':
				$this->sql .= " from {$this->table} where {$this->where}";
				break;
		}
		return $this->sql;
	}
	
	public function connect() {
		$params = $this->master == true ? $this->connectParams['master'] : $this->connectParams['slave'];
		$poolId = "mysql" . md5(implode("", $params));
		if (isset(self::$pool[$poolId]) && is_object(self::$pool[$poolId])) {
			return self::$pool[$poolId];
		}
		self::$pool[$poolId] = new PDO("mysql:dbname={$params['database_name']};host={$params['database_host']};port={$params['database_port']};charset=utf8", $params['database_user'], $params['database_password']);
	}
}