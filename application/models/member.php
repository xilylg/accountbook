<?php

class memberModel extends AbstractModel
{
	public function __construct($log = NULL) {
	    parent::__construct($log);
		$this->datamodel = new memberData($this->log);
	}
}