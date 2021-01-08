<?php
class Book_DeleteController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'book_id' => 'int|required',
    ];
    
    public function indexAction()
    {
        $bookModel = new BookModel($this->log);
        //检查是否所属
        $info = $bookModel->findOne(['book_id' => $this->params['book_id']]);
        if ($info['uid'] != $this->params['uid']) {
            return $this->error(Helper_Conf::get("interfacecode.common.invalid_owner"));
        }
        
        $result = $bookModel->deleteOne($this->params['id']);
        return $this->success($result);
    }
}