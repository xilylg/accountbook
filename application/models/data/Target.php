<?php
class DataTargetModel extends AbstractModel
{
    protected $fields = ["target_id",  "uid",  "content",  "status",  "pic",  "expert_start_time",  "expert_finish_time",  "review_time",  "review_result",  "review_star",  "create_time",  "update_time"];
    protected $primaryKey = 'target_id';
}