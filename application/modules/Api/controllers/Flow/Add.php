<?php

class Flow_AddController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'cost' => 'int|required',
        'book_id' => 'int|required',
        'item_id' => 'int|required',
        'member_id' => 'int',
        'tip' => 'string',
    ];
    
    public function indexAction() {
        $flowModel = new FlowModel($this->log);
        $result = $flowModel->addOne($this->params);
        return $this->success($result);
    }
}