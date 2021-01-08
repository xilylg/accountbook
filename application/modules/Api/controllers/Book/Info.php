<?php

class Book_InfoController extends ApiController
{
    protected $verify = [
        'book_id' => 'int|required',
    ];
    
    public function indexAction() {
        $bookModel = new BookModel($this->log);
        $result = $bookModel->findOne($this->params['book_id']);
        return $this->success($result);;
    }
}