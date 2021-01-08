<?php
class DataMemberModel extends AbstractModel
{
    protected $table = 'member';
    protected $fields = ["member_id",  "uid",  "nickname",  "create_time",  "relation",  "status"];
    protected $primaryKey = "member_id";
}