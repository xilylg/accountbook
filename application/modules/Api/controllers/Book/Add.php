<?php

class Book_AddController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'name' => 'string|required',
        'cover' => 'string',
        'cycle_type' => 'string|required',
        'cycle_start' => 'string|required',
        'cycle_end' => 'string|required',
    ];
    
    public function indexAction() {
        $bookModel = new BookModel($this->log);
        $result = $bookModel->addOne($this->params);
        return $this->success($result);
    }
}