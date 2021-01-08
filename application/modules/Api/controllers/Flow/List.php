<?php

class Flow_ListController extends ApiController
{
    protected $method = 'GET';
    
    public function indexAction() {
        $flowModel = new FlowModel($this->log);
        $result = $flowModel->getListByUid($this->params['uid']);
        return $this->success($result);;
    }
}