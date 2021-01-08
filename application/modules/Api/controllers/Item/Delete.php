<?php
class Item_DeleteController extends ApiController
{
    protected $method = 'POST';
    protected $verify = [
        'item_id' => 'int|required',
    ];
    
    public function indexAction()
    {
        $itemModel = new ItemModel($this->log);
        //检查是否所属
        $info = $itemModel->findOne(['item_id' => $this->params['item_id']]);
        if ($info['uid'] != $this->params['uid']) {
            return $this->error(Helper_Conf::get("interfacecode.common.invalid_owner"));
        }
        
        $result = $itemModel->deleteOne($this->params['item_id']);
        return $this->success($result);
    }
}