<?php
class DataBookModel extends AbstractData
{
    protected $fields = ['book_id', 'name', 'cover', 'uid', 'target_id', 'cycle_type', 'cycle_start', 'cycle_end', 'create_time', 'update_time', 'status'];
    public $primaryKey = 'book_id';
}