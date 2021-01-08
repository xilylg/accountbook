<?php

class Item_UpdateController extends ApiController
{

    protected $method = 'POST';

    protected $verify = [
        'item_id' => 'int|required',
        'item_name' => 'string|required',
        'flow' => 'int|required'
    ];

    public function indexAction()
    {
        $itemModel = new ItemModel($this->log);
        $result = $itemModel->updateOne($this->params);
        return $this->success($result);
    }
}