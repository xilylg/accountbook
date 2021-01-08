<?php

class Target_AddController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'content' => 'string|required',
        'pic' => 'string|required',
        'expert_start_time' => 'string|required',
        'expert_finish_time' => 'string|required',
    ];
    
    public function indexAction() {
        $targetModel = new TargetModel($this->log);
        $this->params['expert_start_time'] = strtotime($this->params['expert_start_time']);
        $this->params['expert_finish_time'] = strtotime($this->params['expert_finish_time']);
        $result = $targetModel->addOne($this->params);
        return $this->success($result);
    }
}