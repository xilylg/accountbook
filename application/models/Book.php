<?php
class BookModel extends AbstractModel
{
	private $bookData;
	private $log;
	
	public function __construct($log = NULL) {
	    parent::__construct($log);
		$this->bookData = new BookData($this->log);
	}
	
	public function addBook($params)
	{
	    $result = $this->bookData->addOne($params);
	    if 
	}
}