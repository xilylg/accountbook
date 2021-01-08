<?php

class Book_ListController extends ApiController
{
    protected $method = 'GET';
    
    public function indexAction() {
        $bookModel = new BookModel($this->log);
        $result = $bookModel->getListByUid($this->params['uid']);
        return $this->success($result);;
    }
}