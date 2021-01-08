<?php
class DataItemModel extends AbstractModel
{
    protected $table = 'item';
    protected $fields = ["item_id",  "uid",  "item_name",  "create_time",  "flow"];
    protected $primaryKey = "item_id";
    
}