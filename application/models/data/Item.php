<?php
class DataItemModel extends AbstractModel
{
    protected $fields = ["item_id",  "uid",  "item_name",  "create_time",  "flow"];
    protected $primaryKey = "item_id";
    
}