<?php
class AbstractModel
{
    protected $log;
    protected $fields = [];
    protected $datamodel;
    
    public function __construct($log = NULL){
        $this->log = is_object($log) ? $log : new Log_Log(LOG_PATH);
    }
    
    public function addOne($params)
    {
        $result = $this->datamodel->addOne($params);
        return $result;
    }
    
    public function updateOne($params)
    {
        return $this->datamodel->updateOne($params);
    }
    
    public function deleteOne($id)
    {
        return $this->datamodel->_deleteOne($this->datamodel->primaryKey, $id);
    }
    
}