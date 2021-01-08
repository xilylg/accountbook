<?php

class Flow_InfoController extends ApiController
{
    protected $verify = [
        'flow_id' => 'int|required',
    ];
    
    public function indexAction() {
        $flowModel = new FlowModel($this->log);
        $result = $flowModel->findOne($this->params['flow_id']);
        return $this->success($result);;
    }
}