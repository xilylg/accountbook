<?php
class Flow_DeleteController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'flow_id' => 'int|required',
    ];
    
    public function indexAction()
    {
        $flowModel = new FlowModel($this->log);
        //检查是否所属
        $info = $flowModel->findOne(['flow_id' => $this->params['flow_id']]);
        if ($info['uid'] != $this->params['uid']) {
            return $this->error(Helper_Conf::get("interfacecode.common.invalid_owner"));
        }
        
        $result = $flowModel->deleteOne($this->params['flow_id']);
        return $this->success($result);
    }
}