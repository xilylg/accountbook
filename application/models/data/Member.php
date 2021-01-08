<?php
class DataMemberModel extends AbstractModel
{
    protected $fields = ["member_id",  "uid",  "nickname",  "create_time",  "relation",  "status"];
    protected $primaryKey = "member_id";
}