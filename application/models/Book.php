<?php
class BookModel extends AbstractModel
{
	public function __construct($log = NULL) {
	    parent::__construct($log);
		$this->datamodel = new BookData($this->log);
	}
}