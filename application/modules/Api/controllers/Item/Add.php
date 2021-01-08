<?php

class Item_AddController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'item_name' => 'string|required',
        'flow' => 'int|required',
    ];
    
    public function indexAction() {
        $itemModel = new ItemModel($this->log);
        $result = $itemModel->addOne($this->params);
        return $this->success($result);
    }
}