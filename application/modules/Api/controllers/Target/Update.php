<?php

class Target_UpdateController extends ApiController
{

    protected $method = 'POST';

    protected $verify = [
        'target_id' => 'int|required',
        'content' => 'string|required',
        'pic' => 'string|required',
        'expert_start_time' => 'string|required',
        'expert_finish_time' => 'string|required',
        
        'review_result' => 'string',
        'review_star' => 'string',
    ];

    public function indexAction()
    {
        $targetModel = new TargetModel($this->log);
        $this->params['review_time'] = time();
        $result = $targetModel->updateOne($this->params);
        return $this->success($result);
    }
}