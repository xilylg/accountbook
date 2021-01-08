<?php

class Member_AddController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'nickname' => 'string|required',
        'relation' => 'string|required',
    ];
    
    public function indexAction() {
        $memberModel = new MemberModel($this->log);
        $result = $memberModel->addOne($this->params);
        return $this->success($result);
    }
}