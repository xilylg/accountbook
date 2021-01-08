<?php

class Book_UpdateController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'book_id' => 'int|required',
        'name' => 'string|required',
        'cover' => 'string',
        'cycle_type' => 'string|required',
        'cycle_start' => 'string|required',
        'cycle_end' => 'string|required',
    ];
    
    public function indexAction() {
        $bookModel = new BookModel($this->log);
        
        //检查是否所属
        $info = $bookModel->findOne(['book_id' => $this->params['book_id']]);
        if ($info['uid'] != $this->params['uid']) {
            return $this->error(Helper_Conf::get("interfacecode.common.invalid_owner"));
        }
        
        $result = $bookModel->updateOne($this->params);
        return $this->success($result);
    }
}