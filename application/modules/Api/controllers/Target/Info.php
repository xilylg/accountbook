<?php

class Target_InfoController extends ApiController
{
    protected $verify = [
        'target_id' => 'int|required',
    ];
    
    public function indexAction() {
        $targetModel = new TargetModel($this->log);
        $result = $targetModel->findOne($this->params['target_id']);
        return $this->success($result);;
    }
}