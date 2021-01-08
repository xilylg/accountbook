<?php
class BookModel extends AbstractModel
{
	public function __construct($log = NULL) {
	    parent::__construct($log);
		$this->datamodel = new BookData($this->log);
	}
	
	
	public function getBooksByUid($uid)
	{
	    return $this->datamodel->_findOne(['uid' => $uid]);
	}
}