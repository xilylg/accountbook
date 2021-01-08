<?php

class targetModel extends AbstractModel
{
	public function __construct($log = NULL) {
	    parent::__construct($log);
		$this->datamodel = new targetData($this->log);
	}
}