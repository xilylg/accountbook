<?php

class Member_ListController extends ApiController
{
    protected $method = 'GET';
    
    public function indexAction() {
        $memberModel = new MemberModel($this->log);
        $result = $memberModel->getListByUid($this->params['uid']);
        return $this->success($result);;
    }
}