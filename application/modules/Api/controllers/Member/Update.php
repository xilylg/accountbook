<?php

class Member_UpdateController extends ApiController
{

    protected $method = 'POST';

    protected $verify = [
        'member_id' => 'int|required',
        'nickname' => 'string|required',
        'relation' => 'string|required',
    ];

    public function indexAction()
    {
        $memberModel = new MemberModel($this->log);
        $result = $memberModel->updateOne($this->params);
        return $this->success($result);
    }
}