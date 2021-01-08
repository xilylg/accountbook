<?php

class flowModel extends AbstractModel
{
	public function __construct($log = NULL) {
	    parent::__construct($log);
		$this->datamodel = new flowData($this->log);
	}
}