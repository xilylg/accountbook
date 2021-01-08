<?php
class Member_DeleteController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'member_id' => 'int|required',
    ];
    
    public function indexAction()
    {
        $memberModel = new MemberModel($this->log);
        //检查是否所属
        $info = $memberModel->findOne(['member_id' => $this->params['member_id']]);
        if ($info['uid'] != $this->params['uid']) {
            return $this->error(Helper_Conf::get("interfacecode.common.invalid_owner"));
        }
        
        $result = $memberModel->deleteOne($this->params['member_id']);
        return $this->success($result);
    }
}