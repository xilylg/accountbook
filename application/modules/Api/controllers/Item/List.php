<?php

class Item_ListController extends ApiController
{
    protected $method = 'GET';
    
    public function indexAction() {
        $itemModel = new ItemModel($this->log);
        $result = $itemModel->getListByUid($this->params['uid']);
        return $this->success($result);;
    }
}