<?php

class Item_InfoController extends ApiController
{
    protected $verify = [
        'item_id' => 'int|required',
    ];
    
    public function indexAction() {
        $itemModel = new ItemModel($this->log);
        $result = $itemModel->findOne($this->params['item_id']);
        return $this->success($result);;
    }
}