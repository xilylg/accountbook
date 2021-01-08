<?php

class Flow_UpdateController extends ApiController
{

    protected $method = 'POST';

    protected $verify = [
        'flow_id' => 'int|required',
        'cost' => 'int|required',
        'book_id' => 'int|required',
        'item_id' => 'int|required',
        'member_id' => 'int',
        'tip' => 'string',
        
        'review_star' => 'int',
    ];
    
    public function indexAction()
    {
        $flowModel = new FlowModel($this->log);
        $this->params['review_time'] = time();
        $result = $flowModel->updateOne($this->params);
        return $this->success($result);
    }
}