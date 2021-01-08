<?php

class Member_InfoController extends ApiController
{
    protected $verify = [
        'member_id' => 'int|required',
    ];
    
    public function indexAction() {
        $memberModel = new MemberModel($this->log);
        $result = $memberModel->findOne($this->params['member_id']);
        return $this->success($result);;
    }
}