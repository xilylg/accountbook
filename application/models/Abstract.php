<?php
class AbstractModel
{
    private $log;
    private $fields = [];
    
    public function __construct($log = NULL){
        $this->log = is_object($log) ? $log : new Log_Log(LOG_PATH);
    }
    
    
}