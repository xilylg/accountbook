<?php
class DataFlowModel extends AbstractModel
{
    protected $table = 'flow';
    protected $fields = ["flow_id",  "cost",  "uid",  "book_id",  "member_id",  "tip",  "review_star",  "review_time",  "create_time"];
    protected $primaryKey = "flow_id";
    
    public function setTable() {
        $this->datamodel->setTable($this->table . "_" . date('Y'));
    }
}