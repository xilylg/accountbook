<?php

class itemModel extends AbstractModel
{

    public function __construct($log = NULL)
    {
        parent::__construct($log);
        $this->datamodel = new itemData($this->log);
    }
}