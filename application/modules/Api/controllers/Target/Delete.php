<?php
class Target_DeleteController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'target_id' => 'int|required',
    ];
    
    public function indexAction()
    {
        $targetModel = new TargetModel($this->log);
        //检查是否所属
        $info = $targetModel->findOne(['target_id' => $this->params['target_id']]);
        if ($info['uid'] != $this->params['uid']) {
            return $this->error(Helper_Conf::get("interfacecode.common.invalid_owner"));
        }
        
        $result = $targetModel->deleteOne($this->params['target_id']);
        return $this->success($result);
    }
}