<?php

class User_InfoController extends ApiController
{
    protected $verify = [
        'uid' => 'int|required',
    ];
    
    public function indexAction() {
        $userModel = new UserModel($this->log);
        $result = $userModel->findOne($this->params['uid']);
        return $this->success($result);;
    }
}