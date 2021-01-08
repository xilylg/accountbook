<?php

class Target_ListController extends ApiController
{
    protected $method = 'GET';
    
    public function indexAction() {
        $targetModel = new TargetModel($this->log);
        $result = $targetModel->getListByUid($this->params['uid']);
        return $this->success($result);;
    }
}