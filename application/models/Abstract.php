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
        $params['create_time'] = time();
        $result = $this->datamodel->addOne($params);
        return $result;
    }
    
    public function updateOne($params)
    {
        $params['update_time'] = time();
        return $this->datamodel->updateOne($params);
    }
    
    public function deleteOne($id)
    {
        return $this->datamodel->deleteOne($this->datamodel->primaryKey, $id);
    }
    
    public function findOne($id)
    {
        return $this->datamodel->findOne([
            $this->datamodel->primaryKey => $id
        ]);
    }
    
    public function getListByUid($uid)
    {
        return $this->datamodel->find([
            'uid' => $uid
        ]);
    }
    
}