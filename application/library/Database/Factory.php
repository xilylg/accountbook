<?php
class Database_Factory
{
	protected static $pool = [];
	protected static $connector;
	protected static $log;
	
	private function __construct(){
		
	}
	
	public static function instanct($database_type, Array $connect_params, $log = NULL) {
		$poolId = $database_type . md5(implode("", $connect_params['master']));
		if (isset(self::$pool[$poolId])) {
			return self::$pool[$poolId];
		}
		
		switch($database_type) {
			case 'mysql':
				self::$connector = new Database_Mysql($connect_params);
				break;
		}
		
		if (is_object($log)) {
			self::$log = $log;
		}
	}

	public function fetchOne() {
		try{
			return self::$connector->fetchOne();
		} catch(Exception $e) {
			$logMsg = ['msg' => $e->getMessage(), 'sql' => self::$connector->sql, 'bind_params' => $this->self::$connector->bindParams];
			if (empty(self::$log)) {
				throw new Exception($e);
			} else {
				$this->log->error('mysql', $logMsg);
				return ['code' => 0, 'error_msg' => $logMsg];
			}
		}
	}
	
	public function fetchAll() {
		try{
			return self::$connector->fetchOne();
		} catch(Exception $e) {
			$logMsg = ['msg' => $e->getMessage(), 'sql' => self::$connector->sql, 'bind_params' => $this->self::$connector->bindParams];
			if (empty(self::$log)) {
				throw new Exception($e);
			} else {
				$this->log->error('mysql', $logMsg);
			}
			return ['code' => 0, 'error_msg' => $logMsg];
		}
	}
	
	public function query($sql, $bind_params) {
		try{
			return self::$connector->query($sql, $bind_params);
		} catch(Exception $e) {
			$logMsg = ['msg' => $e->getMessage(), 'sql' => self::$connector->sql, 'bind_params' => $this->self::$connector->bindParams];
			if (empty(self::$log)) {
				throw new Exception($e);
			} else {
				$this->log->error('mysql', $logMsg);
			}
			return ['code' => 0, 'error_msg' => $logMsg];
		}
	}
	
	public function insert($data) {
		try{
			return self::$connector->insert($data);
		} catch(Exception $e) {
			$logMsg = ['msg' => $e->getMessage(), 'sql' => self::$connector->sql, 'bind_params' => $this->self::$connector->bindParams];
			if (empty(self::$log)) {
				throw new Exception($e);
			} else {
				$this->log->error('mysql', $logMsg);
			}
			return ['code' => 0, 'error_msg' => $logMsg];
		}
	}
	
	public function update($data) {
		try{
			return self::$connector->update($data);
		} catch(Exception $e) {
			$logMsg = ['msg' => $e->getMessage(), 'sql' => self::$connector->sql, 'bind_params' => $this->self::$connector->bindParams];
			if (empty(self::$log)) {
				throw new Exception($e);
			} else {
				$this->log->error('mysql', $logMsg);
				return ['code' => 0, 'error_msg' => $logMsg];
			}
		}
	}
	
	public function delete() {
		try{
			return self::$connector->delete();
		} catch(Exception $e) {
			$logMsg = ['msg' => $e->getMessage(), 'sql' => self::$connector->sql, 'bind_params' => $this->self::$connector->bindParams];
			if (empty(self::$log)) {
				throw new Exception($e);
			} else {
				$this->log->error('mysql', $logMsg);
				return ['code' => 0, 'error_msg' => $logMsg];
			}
		}
	}
	
	
	public function __call($method, $params) {
		try {
		    return user_method_call(self::$connector, $method, $params);
		} catch(Exception $e) {
			$log_msg = ['msg' => $e->getMessage(), 'sql' => self::$connector->sql, 'bind_params' => $this->self::$connector->bindParams];
			if (is_object(self::$log)) {
				self::$log->error('mysql', $log_msg);
			}
			return ['code' => 0, 'error_msg' => $log_msg];
		}
	}
}